package container

type Context struct {
	Executor struct {
		Metadata map[string]string
		Data     string
		Path     string
	}
	Container struct {
		Metadata map[string]string
		Data     string
		Code     string
	}
	Request struct {
		Metadata map[string]string
		Data     string
	}
}

func NewContext() Context {
	ctx := Context{}
	ctx.Executor.Metadata = make(map[string]string)
	ctx.Container.Metadata = make(map[string]string)
	ctx.Request.Metadata = make(map[string]string)
	return ctx
}
