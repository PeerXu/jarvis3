package executor

type NewContextFn func() Context
type NewRequestFn func(ctx Context) Request
type NewResponseFn func(req Request) Response
type NewExecutorFn func() Executor

type Constructor struct {
	ContextMaker  NewContextFn
	RequestMaker  NewRequestFn
	ExecutorMaker NewExecutorFn
}

func NewConstructor(fn NewExecutorFn) *Constructor {
	return &Constructor{
		ContextMaker:  NewContext,
		RequestMaker:  NewRequest,
		ExecutorMaker: fn,
	}
}
