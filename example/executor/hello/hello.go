package hello

import (
	"fmt"

	"github.com/PeerXu/jarvis3/executor"
)

type HelloExecutor struct {
	metadata executor.Metadata
	data     executor.Data
}

func (e *HelloExecutor) Execute(req executor.Request) (executor.Response, error) {
	res := executor.NewResponse(req)
	fmt.Println("hello, world!!!!")
	res.Data().Set("hello, world!")
	return res, nil
}

func (e *HelloExecutor) Metadata() executor.Metadata {
	return e.metadata
}

func (e *HelloExecutor) Data() executor.Data {
	return e.data
}

func NewHelloExecutor() executor.Executor {
	return &HelloExecutor{
		metadata: executor.NewMetadata(),
		data:     executor.NewData(),
	}
}

var Constructor *executor.Constructor

func init() {
	Constructor = executor.NewConstructor(NewHelloExecutor)
}
