package project

type Repository interface {
	CreateExecutor(*Executor) (*Executor, error)
	DeleteExecutor(owner string, name string) error
	GetExecutor(owner string, name string) (*Executor, error)
	ListExecutors(owner string) ([]*Executor, error)
	CreateProject(*Project) (*Project, error)
	DeleteProject(owner string, name string) error
	GetProject(owner string, name string) (*Project, error)
	ListProjects(owner string) ([]*Project, error)
	CreateJob(owner string, project string, job *Job) (*Job, error)
	GetJob(owner string, project string, name string) (*Job, error)
	ListJobs(owner string, project string) ([]*Job, error)
	UpdateJob(owner string, name string, proj string, job *Job) (*Job, error)
}
