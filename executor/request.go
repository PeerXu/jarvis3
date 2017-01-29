package executor

type Request interface {
	Context() Context
	Data() Data
}

type request struct {
	context Context
	data    Data
}

func NewRequest(ctx Context) Request {
	req := &request{
		context: ctx,
		data:    NewData(),
	}
	return req
}

func (req *request) Context() Context {
	return req.context
}

func (req *request) Data() Data {
	return req.data
}
