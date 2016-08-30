package project

type Project struct {
	Name  string
	Owner string

	Jobs []*Job
}

func NewProject(name, owner string) *Project {
	return &Project{
		Name:  name,
		Owner: owner,
		Jobs:  []*Job{},
	}
}

type Job struct {
	Name     string
	Status   JobStatus
	Executor *Executor
	Data     []byte
}

const RANDOM_NAME = ""

func NewJob(name string, executor *Executor, data []byte) *Job {
	return &Job{
		Name:     name,
		Status:   JobStatus_Prepare,
		Executor: executor,
		Data:     data,
	}
}

type JobStatus int

const (
	JobStatus_Unknown JobStatus = 0
	JobStatus_Error   JobStatus = 1
	JobStatus_Prepare JobStatus = 2
	JobStatus_Ready   JobStatus = 3
	JobStatus_Running JobStatus = 4
	JobStatus_Stop    JobStatus = 5
)

var JobStatus_name = map[int]string{
	0: "unknown",
	1: "error",
	2: "prepare",
	3: "ready",
	4: "running",
	5: "stop",
}

var JobStatus_value = map[string]int{
	"unknown": 0,
	"error":   1,
	"prepare": 2,
	"ready":   3,
	"running": 4,
	"stop":    5,
}

func (x JobStatus) String() string {
	return JobStatus_name[int(x)]
}

func LookupJobStatus(s string) JobStatus {
	return JobStatus(JobStatus_value[s])
}

type Executor struct {
	Name  string
	Pack  string
	Owner string
	Data  []byte
}

func NewExecutor(name string, pack string, owner string, data []byte) *Executor {
	return &Executor{
		Name:  name,
		Pack:  pack,
		Owner: owner,
		Data:  data,
	}
}
