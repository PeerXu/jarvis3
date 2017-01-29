package computing

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/computing/service"
	"github.com/PeerXu/jarvis3/project"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateExecutor(ctx context.Context, name, pack string, data []byte) (e *project.Executor, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateExecutor",
			"name", name,
			"pack", pack,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateExecutor(ctx, name, pack, data)
}

func (s *loggingService) DeleteExecutorByID(ctx context.Context, execID project.ExecutorID) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "DeleteExecutorByID",
			"executor_id", execID,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteExecutorByID(ctx, execID)
}

func (s *loggingService) GetExecutorByID(ctx context.Context, execID project.ExecutorID) (exec *project.Executor, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetExecutorByID",
			"executor_id", execID,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetExecutorByID(ctx, execID)
}

func (s *loggingService) ListExecutors(ctx context.Context) (es []*project.Executor, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ListExecutors",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.ListExecutors(ctx)
}

func (s *loggingService) CreateProject(ctx context.Context, name string) (p *project.Project, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateProject",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateProject(ctx, name)
}

func (s *loggingService) DeleteProjectByID(ctx context.Context, projID project.ProjectID) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "DeleteProjectByID",
			"project_id", projID.String(),
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteProjectByID(ctx, projID)
}

func (s *loggingService) GetProjectByID(ctx context.Context, projID project.ProjectID) (p *project.Project, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetProjectByID",
			"project_id", projID,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetProjectByID(ctx, projID)
}

func (s *loggingService) ListProjects(ctx context.Context) (ps []*project.Project, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "ListProjects",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.ListProjects(ctx)
}

func (s *loggingService) CreateTask(ctx context.Context, projID project.ProjectID, execID project.ExecutorID, name string, data []byte) (t *project.Task, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateTask",
			"project_id", projID.String(),
			"executor_id", execID.String(),
			"name", name,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateTask(ctx, projID, execID, name, data)
}

func (s *loggingService) UpdateTaskByID(ctx context.Context, taskID project.TaskID, task *project.Task) (t *project.Task, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "UpdateTaskByID",
			"task_id", taskID,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.UpdateTaskByID(ctx, taskID, task)
}

func (s *loggingService) PopReadyTask(ctx context.Context) (t *project.Task, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "PopReadyTask",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.PopReadyTask(ctx)
}
