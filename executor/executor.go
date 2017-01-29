package executor

type Executor interface {
	Metadata() Metadata
	Data() Data
	Execute(Request) (Response, error)
}
