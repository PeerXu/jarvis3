package project

import "github.com/PeerXu/jarvis3/user"

type Repository interface {
	CreateExecutor(*Executor) (*Executor, error)
	DeleteExecutorByID(id ExecutorID) error
	GetExecutorByID(id ExecutorID) (*Executor, error)
	ListExecutors(id user.UserID) ([]*Executor, error)
	CreateProject(*Project) (*Project, error)
	DeleteProjectByID(id ProjectID) error
	GetProjectByID(id ProjectID) (*Project, error)
	ListProjects(id user.UserID) ([]*Project, error)
	CreateTask(task *Task) (*Task, error)
	DeleteTaskByID(id TaskID) error
	GetTaskByID(id TaskID) (*Task, error)
	ListTasksByProjectID(projID ProjectID) ([]*Task, error)
	UpdateTaskByID(id TaskID, task *Task) (*Task, error)
	PopReadyTask() (*Task, error)
}
