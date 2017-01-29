package project

import "errors"

var (
	ErrUnknown          = errors.New("project unknown error")
	ErrOwnerNotFound    = errors.New("owner not found")
	ErrExecutorNotFound = errors.New("executor not found")
	ErrProjectNotFound  = errors.New("project not found")
	ErrTaskNotFound     = errors.New("task not found")

	ErrNotReadyTask   = errors.New("ready task not found")
	ErrNotRunningTask = errors.New("running task not found")

	ErrTaskNotRunning = errors.New("can't change to stop or error for not running task")
)
