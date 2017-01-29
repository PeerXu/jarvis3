package executor

type Response interface {
	Data() Data
	Request() Request
}

type response struct {
	data    Data
	request Request
}

func NewResponse(req Request) Response {
	res := &response{request: req, data: NewData()}
	return res
}

func (res *response) Data() Data {
	return res.data
}

func (res *response) Request() Request {
	return res.request
}
