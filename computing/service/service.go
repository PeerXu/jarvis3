package service

import (
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/project"
)

type Service interface {
	CreateExecutor(ctx context.Context, name string, pack string, data []byte) (*project.Executor, error)
	DeleteExecutorByID(ctx context.Context, execID project.ExecutorID) error
	GetExecutorByID(ctx context.Context, execID project.ExecutorID) (*project.Executor, error)
	ListExecutors(ctx context.Context) ([]*project.Executor, error)

	CreateProject(ctx context.Context, name string) (*project.Project, error)
	DeleteProjectByID(ctx context.Context, id project.ProjectID) error
	GetProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error)
	ListProjects(ctx context.Context) ([]*project.Project, error)

	CreateTask(ctx context.Context, projID project.ProjectID, execID project.ExecutorID, name string, data []byte) (*project.Task, error)
	DeleteTaskByID(ctx context.Context, id project.TaskID) error
	GetTaskByID(ctx context.Context, id project.TaskID) (*project.Task, error)
	UpdateTaskByID(ctx context.Context, id project.TaskID, task *project.Task) (*project.Task, error)
	PopReadyTask(ctx context.Context) (*project.Task, error)
}
