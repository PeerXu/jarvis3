package computing

import (
	"time"

	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

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

func (s *loggingService) DeleteExecutor(ctx context.Context, name string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "DeleteEexecutor",
			"name", name,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteExecutor(ctx, name)
}

func (s *loggingService) GetExecutor(ctx context.Context, name string) (e *project.Executor, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetExecutor",
			"name", name,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetExecutor(ctx, name)
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

func (s *loggingService) DeleteProject(ctx context.Context, name string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "DeleteProject",
			"name", name,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.DeleteProject(ctx, name)
}

func (s *loggingService) GetProject(ctx context.Context, name string) (p *project.Project, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "GetProject",
			"name", name,
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.GetProject(ctx, name)
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

func (s *loggingService) CreateJob(ctx context.Context, name string, proj string, executor string, data []byte) (j *project.Job, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "CreateJob",
			"took", time.Since(begin),
			"error", err,
		)
	}(time.Now())
	return s.Service.CreateJob(ctx, name, proj, executor, data)
}

func (s *loggingService) UpdateJob(ctx context.Context, name string, proj string, job *project.Job) (j *project.Job, err error) {
	defer func(begin time.Time) {
	}(time.Now())
	return s.Service.UpdateJob(ctx, name, proj, job)
}
