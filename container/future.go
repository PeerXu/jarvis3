package container

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

type Future interface {
	Wait() error
}

type future struct {
	context Context
}

func NewFuture(ctx Context) Future {
	return &future{ctx}
}

func (f *future) Wait() error {
	document := js.Global.Get("document")
	doneBuf := document.Call("getElementById", "__JVS_DONE_BUF_"+f.context.Container.Code)

	for {
		if doneBuf.Get("innerHTML").String() != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
