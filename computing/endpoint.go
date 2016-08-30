package computing

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/project"
)

type executorBody struct {
	Name  string `json:"name"`
	Pack  string `json:"pack"`
	Owner string `json:"owner"`
	Data  []byte `json:"data"`
}

func encodeExecutor2ExecutorBody(e *project.Executor) executorBody {
	return executorBody{
		Name:  e.Name,
		Pack:  e.Pack,
		Owner: e.Owner,
		Data:  e.Data,
	}
}

type jobBody struct {
	Name     string       `json:"name"`
	Status   string       `json:"status"`
	Executor executorBody `json:"executor"`
	Data     []byte       `json:"data"`
}

func encodeJob2JobBody(j *project.Job) jobBody {
	return jobBody{
		Name:     j.Name,
		Status:   j.Status.String(),
		Executor: encodeExecutor2ExecutorBody(j.Executor),
		Data:     j.Data,
	}
}

type projectBody struct {
	Name  string    `json:"name"`
	Owner string    `json:"owner"`
	Jobs  []jobBody `json:"jobs"`
}

func encodeProject2ProjectBody(p *project.Project) projectBody {
	var js []jobBody
	for _, j := range p.Jobs {
		js = append(js, encodeJob2JobBody(j))
	}
	return projectBody{
		Name:  p.Name,
		Owner: p.Owner,
		Jobs:  js,
	}
}

type computeCreateExecutorRequest struct {
	Name string
	Pack string
	Data []byte
}

type computeCreateExecutorResponse executorBody

func makeComputeCreateExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeCreateExecutorRequest)
		e, err := s.CreateExecutor(ctx, req.Name, req.Pack, req.Data)
		if err != nil {
			return err, nil
		}
		return encodeExecutor2ExecutorBody(e), nil
	}
}

type computeDeleteExecutorRequest struct {
	Name string `json:"-"`
}

type computeDeleteExecutorResponse struct{}

func makeComputeDeleteExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeDeleteExecutorRequest)
		err := s.DeleteExecutor(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return computeDeleteExecutorResponse{}, nil
	}
}

type computeGetExecutorRequest struct {
	Name string `json:"-"`
}

type computeGetExecutorResponse executorBody

func makeComputeGetExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeGetExecutorRequest)
		e, err := s.GetExecutor(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return encodeExecutor2ExecutorBody(e), nil
	}
}

type computeListExecutorsRequest struct{}

type computeListExecutorsResponse []executorBody

func makeComputeListExecutorsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		es, err := s.ListExecutors(ctx)
		if err != nil {
			return err, nil
		}

		var res []executorBody
		for _, e := range es {
			eb := encodeExecutor2ExecutorBody(e)
			res = append(res, eb)
		}

		return computeListExecutorsResponse(res), nil
	}
}

type computeCreateProjectRequest struct {
	Name string
}

type computeCreateProjectResponse projectBody

func makeComputeCreateProjectEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeCreateExecutorRequest)
		p, err := s.CreateProject(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return encodeProject2ProjectBody(p), nil
	}
}

type computeDeleteProjectRequest struct {
	Name string
}

type computeDeleteProjectResponse struct{}

func makeComputeDeleteProjectEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeDeleteExecutorRequest)
		err := s.DeleteProject(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return computeDeleteProjectResponse{}, nil
	}
}

type computeGetProjectRequest struct {
	Name string
}

type computeGetProjectResponse projectBody

func makeComputeGetProjectEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeGetProjectRequest)
		p, err := s.GetProject(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return encodeProject2ProjectBody(p), nil
	}
}

type computeListProjectsRequest struct{}

type computeListProjectsResponse []projectBody

func makeComputeListProjectsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ps, err := s.ListProjects(ctx)
		if err != nil {
			return err, nil
		}

		var res []projectBody
		for _, p := range ps {
			res = append(res, encodeProject2ProjectBody(p))
		}
		return res, nil
	}
}

type computeCreateJobRequest struct {
	Name     string
	Project  string
	Executor string
	Data     []byte
}

type computeCreateJobResponse jobBody

func makeComputeCreateJobEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeCreateJobRequest)
		j, err := s.CreateJob(ctx, req.Name, req.Project, req.Executor, req.Data)
		if err != nil {
			return err, nil
		}
		return encodeJob2JobBody(j), nil

	}
}

type computeUpdateJobRequest struct {
	Name    string
	Project string
	Status  string
}

type computeUpdateJobResponse jobBody

func makeComputeUpdateJobEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(computeUpdateJobRequest)
		j, err := s.UpdateJob(ctx, req.Name, req.Project, &project.Job{Status: project.LookupJobStatus(req.Status)})
		if err != nil {
			return err, nil
		}

		return encodeJob2JobBody(j), nil
	}
}
