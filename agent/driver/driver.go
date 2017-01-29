package driver

import (
	"github.com/PeerXu/jarvis3/container"
	"github.com/PeerXu/jarvis3/project"
)

type AgentDriver interface {
	FetchTask() (*project.Task, error)
	CompleteTask(*project.Task) error

	CreateContainer(*project.Task) (container.Container, error)
	CloseContainer(container.Container) error
}
