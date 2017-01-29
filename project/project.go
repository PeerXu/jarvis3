package project

import (
	"github.com/PeerXu/jarvis3/user"
	"github.com/PeerXu/jarvis3/utils"
)

type ProjectID string

func (id ProjectID) String() string {
	return string(id)
}

type Project struct {
	ID      ProjectID
	OwnerID user.UserID
	Name    string
	Tasks   []*Task
}

func NewProject(ownerID user.UserID, name string) *Project {
	return &Project{
		ID:      ProjectID(utils.NewRandomUUIDString()),
		OwnerID: ownerID,
		Name:    name,
		Tasks:   []*Task{},
	}
}

type TaskID string

func (id TaskID) String() string {
	return string(id)
}

type Task struct {
	ID         TaskID
	ProjectID  ProjectID
	ExecutorID ExecutorID
	Name       string
	Status     TaskStatus
	Data       []byte
}

func NewTask(projID ProjectID, execID ExecutorID, name string, data []byte) *Task {
	return &Task{
		ID:         TaskID(utils.NewRandomUUIDString()),
		ProjectID:  projID,
		ExecutorID: execID,
		Name:       name,
		Status:     TaskStatus_Prepare,
		Data:       data,
	}
}

type TaskStatus int

const (
	TaskStatus_Unknown TaskStatus = 0
	TaskStatus_Error   TaskStatus = 1
	TaskStatus_Prepare TaskStatus = 2
	TaskStatus_Ready   TaskStatus = 3
	TaskStatus_Running TaskStatus = 4
	TaskStatus_Stop    TaskStatus = 5
)

var TaskStatus_name = map[int]string{
	0: "unknown",
	1: "error",
	2: "prepare",
	3: "ready",
	4: "running",
	5: "stop",
}

var TaskStatus_value = map[string]int{
	"unknown": 0,
	"error":   1,
	"prepare": 2,
	"ready":   3,
	"running": 4,
	"stop":    5,
}

func (x TaskStatus) String() string {
	return TaskStatus_name[int(x)]
}

func LookupTaskStatus(s string) TaskStatus {
	return TaskStatus(TaskStatus_value[s])
}

type ExecutorID string

func (id ExecutorID) String() string {
	return string(id)
}

type Executor struct {
	ID      ExecutorID
	OwnerID user.UserID
	Name    string
	Pack    string
	Data    []byte
}

func NewExecutor(ownerID user.UserID, name string, pack string, data []byte) *Executor {
	return &Executor{
		ID:      ExecutorID(utils.NewRandomUUIDString()),
		OwnerID: ownerID,
		Name:    name,
		Pack:    pack,
		Data:    data,
	}
}
