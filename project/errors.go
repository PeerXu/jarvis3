package project

import "errors"

var (
	ErrUnknown          = errors.New("project unknown error")
	ErrOwnerNotFound    = errors.New("owner not found")
	ErrExecutorNotFound = errors.New("executor not found")
	ErrProjectNotFound  = errors.New("project not found")
	ErrJobNotFound      = errors.New("job not found")
)
