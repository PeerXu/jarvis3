package computing

import (
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"

	jcontext "github.com/PeerXu/jarvis3/context"
	"github.com/PeerXu/jarvis3/project"
)

type Service interface {
	CreateExecutor(ctx context.Context, name string, pack string, data []byte) (*project.Executor, error)
	DeleteExecutor(ctx context.Context, name string) error
	GetExecutor(ctx context.Context, name string) (*project.Executor, error)
	ListExecutors(ctx context.Context) ([]*project.Executor, error)
	CreateProject(ctx context.Context, name string) (*project.Project, error)
	DeleteProject(ctx context.Context, name string) error
	GetProject(ctx context.Context, name string) (*project.Project, error)
	ListProjects(ctx context.Context) ([]*project.Project, error)
	CreateJob(ctx context.Context, name string, proj string, executor string, data []byte) (*project.Job, error)
	UpdateJob(ctx context.Context, name string, proj string, job *project.Job) (*project.Job, error)
}

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
	e := project.NewExecutor(name, pack, u.Username, data)
	_, err := s.projectRepository.CreateExecutor(e)
	if err != nil {
		return nil, newComputeError(errorServerError, "failed to create executor", err)
	}
	return e, nil
}

func (s *service) DeleteExecutor(ctx context.Context, name string) error {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()
	err := s.projectRepository.DeleteExecutor(u.Username, name)
	if err != nil {
		return newComputeError(errorServerError, "failed to delete executor", err)
	}
	return nil
}

func (s *service) GetExecutor(ctx context.Context, name string) (*project.Executor, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()
	e, err := s.projectRepository.GetExecutor(u.Username, name)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
		case project.ErrExecutorNotFound:
			return nil, newComputeError(errorNotFound, "executor not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to get executor", err)
		}
	}
	return e, nil
}

func (s *service) ListExecutors(ctx context.Context) ([]*project.Executor, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	executors, err := s.projectRepository.ListExecutors(u.Username)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			return nil, newComputeError(errorNotFound, "owner not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to list executors", err)
		}
	}
	return executors, nil
}

func (s *service) CreateProject(ctx context.Context, name string) (*project.Project, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	proj := project.NewProject(name, u.Username)
	p, err := s.projectRepository.CreateProject(proj)

	if err != nil {
		return nil, newComputeError(errorServerError, "failed to create project", err)
	}

	return p, nil
}

func (s *service) DeleteProject(ctx context.Context, name string) error {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	err := s.projectRepository.DeleteProject(u.Username, name)
	if err != nil {
		return newComputeError(errorServerError, "failed to delete project", err)
	}

	return nil
}

func (s *service) GetProject(ctx context.Context, name string) (*project.Project, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	proj, err := s.projectRepository.GetProject(u.Username, name)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			return nil, newComputeError(errorNotFound, "owner not found", err)
		case project.ErrProjectNotFound:
			return nil, newComputeError(errorNotFound, "project not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to get project", err)
		}
	}
	return proj, nil
}

func (s *service) ListProjects(ctx context.Context) ([]*project.Project, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	projs, err := s.projectRepository.ListProjects(u.Username)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			return nil, newComputeError(errorNotFound, "owner not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to list projects", err)
		}
	}
	return projs, nil
}

func (s *service) CreateJob(ctx context.Context, name string, proj string, executor string, data []byte) (*project.Job, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	e, err := s.projectRepository.GetExecutor(u.Username, executor)
	if err != nil {
		switch err {
		case project.ErrOwnerNotFound:
			return nil, newComputeError(errorNotFound, "owner not found", err)
		case project.ErrExecutorNotFound:
			return nil, newComputeError(errorNotFound, "executor not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to get executor", err)
		}
	}

	j := project.NewJob(name, e, data)

	j, err = s.projectRepository.CreateJob(u.Username, proj, j)
	if err != nil {
		switch err {
		case project.ErrProjectNotFound:
			return nil, newComputeError(errorNotFound, "project not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to create job", err)
		}
	}

	return j, nil
}

func (s *service) UpdateJob(ctx context.Context, name string, proj string, job *project.Job) (*project.Job, error) {
	jctx := ctx.Value("JarvisContext").(jcontext.Context)
	u := jctx.User()

	job, err := s.projectRepository.UpdateJob(u.Username, name, proj, job)
	if err != nil {
		switch err {
		case project.ErrJobNotFound:
			return nil, newComputeError(errorNotFound, "job not found", err)
		default:
			return nil, newComputeError(errorServerError, "failed to update job", err)
		}
	}

	return job, nil
}
