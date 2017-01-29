package computing

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/computing/error"
	. "github.com/PeerXu/jarvis3/computing/service"
	jcontext "github.com/PeerXu/jarvis3/context"
	jerrors "github.com/PeerXu/jarvis3/errors"
	"github.com/PeerXu/jarvis3/project"
)

type service struct {
	logger            log.Logger
	projectRepository project.Repository
}

func NewService(logger log.Logger, projectRepository project.Repository) Service {
	return &service{
		logger:            logger,
		projectRepository: projectRepository,
	}
}

func (s *service) CreateExecutor(ctx context.Context, name string, pack string, data []byte) (*project.Executor, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()
	e := project.NewExecutor(u.ID, name, pack, data)
	_, err := s.projectRepository.CreateExecutor(e)
	if err != nil {
		err = NewComputeError(jerrors.ErrorServerError, "failed to create executor", err)
		return nil, err
	}
	return e, nil
}

func (s *service) DeleteExecutorByID(ctx context.Context, execID project.ExecutorID) error {
	err := s.projectRepository.DeleteExecutorByID(execID)
	if err != nil {
		switch err {
		case project.ErrExecutorNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "executor not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to delete executor", err)
		}
		return err
	}
	return nil
}

func (s *service) GetExecutorByID(ctx context.Context, execID project.ExecutorID) (*project.Executor, error) {
	exec, err := s.projectRepository.GetExecutorByID(execID)
	if err != nil {
		switch err {
		case project.ErrExecutorNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "executor not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to get executor", err)
		}
		return nil, err
	}
	return exec, nil
}

func (s *service) ListExecutors(ctx context.Context) ([]*project.Executor, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	executors, err := s.projectRepository.ListExecutors(u.ID)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "owner not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to list executors", err)
		}
		return nil, err
	}
	return executors, nil
}

func (s *service) CreateProject(ctx context.Context, name string) (*project.Project, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	proj := project.NewProject(u.ID, name)
	p, err := s.projectRepository.CreateProject(proj)

	if err != nil {
		return nil, NewComputeError(jerrors.ErrorServerError, "failed to create project", err)
	}

	return p, nil
}

func (s *service) DeleteProjectByID(ctx context.Context, id project.ProjectID) error {
	err := s.projectRepository.DeleteProjectByID(id)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to delete project", err)
		}
		return err
	}
	return nil
}

func (s *service) GetProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error) {
	proj, err := s.projectRepository.GetProjectByID(id)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to get project", err)
		}
		return nil, err
	}
	return proj, nil
}

func (s *service) ListProjects(ctx context.Context) ([]*project.Project, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	projs, err := s.projectRepository.ListProjects(u.ID)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "owner not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to list projects", err)
		}
		return nil, err
	}
	return projs, nil
}

func (s *service) CreateTask(ctx context.Context, projID project.ProjectID, execID project.ExecutorID, name string, data []byte) (*project.Task, error) {
	t := project.NewTask(projID, execID, name, data)

	t, err := s.projectRepository.CreateTask(t)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to create task", err)
		}
		return nil, err
	}

	go func(t *project.Task) {
		t.Status = project.TaskStatus_Ready
		t, err := s.projectRepository.UpdateTaskByID(t.ID, t)
		if err != nil {
			s.logger.Log("post create task#%v failed, %v", t.ID, err)
		}
	}(t)

	return t, nil
}

func (s *service) UpdateTaskByID(ctx context.Context, id project.TaskID, task *project.Task) (*project.Task, error) {
	t, err := s.projectRepository.UpdateTaskByID(id, task)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		case project.ErrTaskNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "task not found", err)
		case project.ErrTaskNotRunning:
			err = NewComputeError(jerrors.ErrorInvalidRequest, "status can't change stop or error when not running", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to get task", err)
		}
		return nil, err
	}
	return t, nil
}

func (s *service) GetTaskByID(ctx context.Context, id project.TaskID) (*project.Task, error) {
	t, err := s.projectRepository.GetTaskByID(id)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		case project.ErrTaskNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "task not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to get task", err)
		}
		return nil, err
	}
	return t, nil
}

func (s *service) DeleteTaskByID(ctx context.Context, id project.TaskID) error {
	err := s.projectRepository.DeleteTaskByID(id)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "project not found", err)
		case project.ErrTaskNotFound:
			err = NewComputeError(jerrors.ErrorNotFound, "task not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to delete task", err)
		}
		return err
	}
	return nil
}

func (s *service) PopReadyTask(ctx context.Context) (*project.Task, error) {
	task, err := s.projectRepository.PopReadyTask()
	if err != nil {
		switch err {
		case project.ErrNotReadyTask:
			err = NewComputeError(jerrors.ErrorNotFound, "ready task not found", err)
		default:
			err = NewComputeError(jerrors.ErrorServerError, "failed to pop ready task", err)
		}
		return nil, err
	}

	go func(task *project.Task) {
		time.Sleep(10 * time.Second)

		t, err := s.projectRepository.GetTaskByID(task.ID)
		if t.Status == project.TaskStatus_Running {
			s.logger.Log("post pop ready task#%v: timeout, reset to ready task", t.ID)
			t.Status = project.TaskStatus_Ready
			_, err = s.projectRepository.UpdateTaskByID(t.ID, task)
			if err != nil {
				s.logger.Log("post pop ready task#%v: %v", t.ID, err)
			}
		}

	}(task)

	return task, nil
}
