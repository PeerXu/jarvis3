package executor

type Context interface {
	Metadata() Metadata
	Data() Data
}

type context struct {
	metadata Metadata
	data     Data
}

func NewContext() Context {
	return &context{
		metadata: NewMetadata(),
		data:     NewData(),
	}
}

func (ctx *context) Metadata() Metadata {
	return ctx.metadata
}

func (ctx *context) Data() Data {
	return ctx.data
}
