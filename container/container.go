package container

type Container interface {
	Run() error
	RunAsync() (Future, error)
	Close()
}
