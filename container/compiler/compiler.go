package compiler

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"honnef.co/go/js/xhr"

	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/js"
)

type Context struct {
	GoCode string
	JsCode string
	Err    error

	done     bool
	waitChan chan bool

	packages      map[string]*compiler.Archive
	pkgsToLoad    map[string]struct{}
	importContext *compiler.ImportContext
	fileSet       *token.FileSet
}

func NewContext(code string) *Context {
	ctx := &Context{
		GoCode: code,

		done:     false,
		waitChan: make(chan bool),

		packages:   make(map[string]*compiler.Archive),
		pkgsToLoad: make(map[string]struct{}),
		fileSet:    token.NewFileSet(),
	}

	ctx.importContext = &compiler.ImportContext{
		Packages: make(map[string]*types.Package),
		Import: func(path string) (*compiler.Archive, error) {
			if pkg, found := ctx.packages[path]; found {
				return pkg, nil
			}

			ctx.pkgsToLoad[path] = struct{}{}
			return &compiler.Archive{}, nil
		},
	}

	return ctx
}

func (ctx *Context) Wait() {
	if !ctx.done {
		<-ctx.waitChan
	}
	return
}

func (ctx *Context) Done() {
	ctx.done = true
	close(ctx.waitChan)
}

func (ctx *Context) SetError(err error) {
	ctx.Err = err
	ctx.Done()
}

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (cmp *Compiler) Compile(code string) (string, error) {
	ctx := NewContext(code)
	cmp.LowCompile(ctx)
	ctx.Wait()
	return ctx.JsCode, ctx.Err
}

func (cmp *Compiler) LowCompile(ctx *Context) {
	file, err := parser.ParseFile(ctx.fileSet, "prog.go", []byte(ctx.GoCode), parser.ParseComments)
	if err != nil {
		ctx.SetError(err)
		return
	}

	buf := bytes.NewBuffer(nil)
	var compile func()
	compile = func() {
		ctx.pkgsToLoad = make(map[string]struct{})
		mainPkg, err := compiler.Compile("main", []*ast.File{file}, ctx.fileSet, ctx.importContext, false)
		ctx.packages["main"] = mainPkg
		if err != nil && len(ctx.pkgsToLoad) == 0 {
			ctx.SetError(err)
			return
		}

		var allPkgs []*compiler.Archive
		if len(ctx.pkgsToLoad) == 0 {
			allPkgs, _ = compiler.ImportDependencies(mainPkg, ctx.importContext.Import)
		}

		if len(ctx.pkgsToLoad) != 0 {
			pkgsReceived := 0
			for path := range ctx.pkgsToLoad {
				req := xhr.NewRequest("GET", "/agent/v1/pkg/"+path+".a.js")
				req.ResponseType = xhr.ArrayBuffer
				go func(path string) {
					err := req.Send(nil)
					if err != nil || req.Status != 200 {
						ctx.SetError(err)
						return
					}
					data := js.Global.Get("Uint8Array").New(req.Response).Interface().([]byte)
					ctx.packages[path], err = compiler.ReadArchive(path+".a", path, bytes.NewReader(data), ctx.importContext.Packages)
					if err != nil {
						ctx.SetError(err)
						return
					}

					pkgsReceived++
					if pkgsReceived == len(ctx.pkgsToLoad) {
						compile()
					}
				}(path)
			}
			return
		}

		compiler.WriteProgramCode(allPkgs, &compiler.SourceMapFilter{Writer: buf})
		ctx.JsCode = buf.String()
		ctx.Done()
	}
	go func() { compile() }()

	return
}

func Compile(code string) (string, error) {
	cmp := NewCompiler()
	return cmp.Compile(code)
}
