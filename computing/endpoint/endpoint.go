package endpoint

import (
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/computing/encode_decode"
	. "github.com/PeerXu/jarvis3/computing/service"
	jerrors "github.com/PeerXu/jarvis3/errors"
	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	"github.com/PeerXu/jarvis3/project"
)

type Endpoints struct {
	ComputeCreateExecutorEndpoint     endpoint.Endpoint
	ComputeDeleteExecutorByIDEndpoint endpoint.Endpoint
	ComputeGetExecutorByIDEndpoint    endpoint.Endpoint
	ComputeListExecutorsEndpoint      endpoint.Endpoint
	ComputeCreateProjectEndpoint      endpoint.Endpoint
	ComputeDeleteProjectByIDEndpoint  endpoint.Endpoint
	ComputeGetProjectByIDEndpoint     endpoint.Endpoint
	ComputeListProjectsEndpoint       endpoint.Endpoint
	ComputeCreateTaskEndpoint         endpoint.Endpoint
	ComputeDeleteTaskByIDEndpoint     endpoint.Endpoint
	ComputeGetTaskByIDEndpoint        endpoint.Endpoint
	ComputeUpdateTaskByIDEndpoint     endpoint.Endpoint
	ComputePopReadyTaskEndpoint       endpoint.Endpoint
}

func MakeClientEndpoints(
	instance string,
	encodeRequestFactory jhttptransport.EncodeRequestFuncMiddleware,
	decodeResponseFactory jhttptransport.DecodeResponseFuncMiddleware,
	opts []httptransport.ClientOption) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}

	tgt.Path = ""

	var computeCreateExecutorClient endpoint.Endpoint
	{
		computeCreateExecutorEncodeRequest := encodeRequestFactory(EncodeComputeCreateExecutorRequest)
		computeCreateExecutorDecodeResponse := decodeResponseFactory(DecodeComputeCreateExecutorResponse)
		computeCreateExecutorOptions := opts
		computeCreateExecutorClient = httptransport.NewClient("POST", tgt, computeCreateExecutorEncodeRequest, computeCreateExecutorDecodeResponse, computeCreateExecutorOptions...).Endpoint()
	}

	var computeListExecutorsClient endpoint.Endpoint
	{
		computeListExecutorsEncodeRequest := encodeRequestFactory(EncodeComputeListExecutorsRequest)
		computeListExecutorsDecodeResponse := decodeResponseFactory(DecodeComputeListExecutorsResponse)
		computeListExecutorsOptions := opts
		computeListExecutorsClient = httptransport.NewClient("GET", tgt, computeListExecutorsEncodeRequest, computeListExecutorsDecodeResponse, computeListExecutorsOptions...).Endpoint()
	}

	var computeDeleteExecutorByIDClient endpoint.Endpoint
	{
		computeDeleteExecutorByIDEncodeRequest := encodeRequestFactory(EncodeComputeDeleteExecutorByIDRequest)
		computeDeleteExecutorByIDDecodeResponse := decodeResponseFactory(DecodeComputeDeleteExecutorByIDResponse)
		computeDeleteExecutorByIDOptions := opts
		computeDeleteExecutorByIDClient = httptransport.NewClient("DELETE", tgt, computeDeleteExecutorByIDEncodeRequest, computeDeleteExecutorByIDDecodeResponse, computeDeleteExecutorByIDOptions...).Endpoint()
	}

	var computeGetExecutorByIDClient endpoint.Endpoint
	{
		computeGetExecutorByIDEncodeRequest := encodeRequestFactory(EncodeComputeGetExecutorByIDRequest)
		computeGetExecutorByIDDecodeResponse := decodeResponseFactory(DecodeComputeGetExecutorByIDResponse)
		computeGetExecutorByIDOptions := opts
		computeGetExecutorByIDClient = httptransport.NewClient("GET", tgt, computeGetExecutorByIDEncodeRequest, computeGetExecutorByIDDecodeResponse, computeGetExecutorByIDOptions...).Endpoint()
	}

	var computeCreateProjectClient endpoint.Endpoint
	{
		computeCreateProjectEncodeRequest := encodeRequestFactory(EncodeComputeCreateProjectRequest)
		computeCreateProjectDecodeResponse := decodeResponseFactory(DecodeComputeCreateProjectResponse)
		computeCreateProjectOptions := opts
		computeCreateProjectClient = httptransport.NewClient("POST", tgt, computeCreateProjectEncodeRequest, computeCreateProjectDecodeResponse, computeCreateProjectOptions...).Endpoint()
	}

	var computeListProjectsClient endpoint.Endpoint
	{
		computeListProjectsEncodeRequest := encodeRequestFactory(EncodeComputeListProjectsRequest)
		computeListProjectsDecodeResponse := decodeResponseFactory(DecodeComputeListProjectsResponse)
		computeListProjectsOptions := opts
		computeListProjectsClient = httptransport.NewClient("GET", tgt, computeListProjectsEncodeRequest, computeListProjectsDecodeResponse, computeListProjectsOptions...).Endpoint()
	}

	var computeDeleteProjectByIDClient endpoint.Endpoint
	{
		computeDeleteProjectByIDEncodeRequest := encodeRequestFactory(EncodeComputeDeleteProjectByIDRequest)
		computeDeleteProjectByIDDecodeResponse := decodeResponseFactory(DecodeComputeDeleteProjectByIDResponse)
		computeDeleteProjectByIDOptions := opts
		computeDeleteProjectByIDClient = httptransport.NewClient("DELETE", tgt, computeDeleteProjectByIDEncodeRequest, computeDeleteProjectByIDDecodeResponse, computeDeleteProjectByIDOptions...).Endpoint()
	}

	var computeCreateTaskClient endpoint.Endpoint
	{
		computeCreateTaskEncodeRequest := encodeRequestFactory(EncodeComputeCreateTaskRequest)
		computeCreateTaskDecodeResponse := decodeResponseFactory(DecodeComputeCreateTaskResponse)
		computeCreateTaskOptions := opts
		computeCreateTaskClient = httptransport.NewClient("POST", tgt, computeCreateTaskEncodeRequest, computeCreateTaskDecodeResponse, computeCreateTaskOptions...).Endpoint()
	}

	var computePopReadyTaskClient endpoint.Endpoint
	{
		computePopReadyTaskEncodeRequest := encodeRequestFactory(EncodeComputePopReadyTaskRequest)
		computePopReadyTaskDecodeResponse := decodeResponseFactory(DecodeComputePopReadyTaskResponse)
		computePopReadyTaskOptions := opts
		computePopReadyTaskClient = httptransport.NewClient("DELETE", tgt, computePopReadyTaskEncodeRequest, computePopReadyTaskDecodeResponse, computePopReadyTaskOptions...).Endpoint()
	}

	var computeGetProjectByIDClient endpoint.Endpoint
	{
		computeGetProjectByIDEncodeRequest := encodeRequestFactory(EncodeComputeGetProjectByIDRequest)
		computeGetProjectByIDDecodeResponse := decodeResponseFactory(DecodeComputeGetProjectByIDResponse)
		computeGetProjectByIDOptions := opts
		computeGetProjectByIDClient = httptransport.NewClient("GET", tgt, computeGetProjectByIDEncodeRequest, computeGetProjectByIDDecodeResponse, computeGetProjectByIDOptions...).Endpoint()
	}

	var computeUpdateTaskByIDClient endpoint.Endpoint
	{
		computeUpdateTaskByIDEncodeRequest := encodeRequestFactory(EncodeComputeUpdateTaskByIDRequest)
		computeUpdateTaskByIDDecodeResponse := decodeResponseFactory(DecodeComputeUpdateTaskByIDResponse)
		computeUpdateTaskByIDOptions := opts
		computeUpdateTaskByIDClient = httptransport.NewClient("PATCH", tgt, computeUpdateTaskByIDEncodeRequest, computeUpdateTaskByIDDecodeResponse, computeUpdateTaskByIDOptions...).Endpoint()
	}

	return Endpoints{
		ComputeCreateExecutorEndpoint:     computeCreateExecutorClient,
		ComputeListExecutorsEndpoint:      computeListExecutorsClient,
		ComputeDeleteExecutorByIDEndpoint: computeDeleteExecutorByIDClient,
		ComputeGetExecutorByIDEndpoint:    computeGetExecutorByIDClient,
		ComputeCreateProjectEndpoint:      computeCreateProjectClient,
		ComputeListProjectsEndpoint:       computeListProjectsClient,
		ComputeDeleteProjectByIDEndpoint:  computeDeleteProjectByIDClient,
		ComputeCreateTaskEndpoint:         computeCreateTaskClient,
		ComputePopReadyTaskEndpoint:       computePopReadyTaskClient,
		ComputeGetProjectByIDEndpoint:     computeGetProjectByIDClient,
		ComputeUpdateTaskByIDEndpoint:     computeUpdateTaskByIDClient,
	}, nil
}

func (e Endpoints) handleError(response interface{}, err error) (error, bool) {
	if err != nil {
		return err, true
	}

	if err, ok := response.(jerrors.JarvisError); ok {
		return err, true
	}

	return nil, false
}

func (e Endpoints) CreateExecutor(ctx context.Context, name string, pack string, data []byte) (*project.Executor, error) {
	request := ComputeCreateExecutorRequest{
		Name: name,
		Pack: pack,
		Data: data,
	}
	response, err := e.ComputeCreateExecutorEndpoint(ctx, request)
	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeCreateExecutorResponse)
	return DecodeExecutorBody2Executor(ExecutorBody(res)), nil
}

func (e Endpoints) GetExecutorByID(ctx context.Context, execID project.ExecutorID) (*project.Executor, error) {
	request := ComputeGetExecutorByIDRequest{execID.String()}
	response, err := e.ComputeGetExecutorByIDEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeGetExecutorByIDResponse)
	return DecodeExecutorBody2Executor(ExecutorBody(res)), nil
}

func (e Endpoints) DeleteExecutorByID(ctx context.Context, execID project.ExecutorID) error {
	request := ComputeDeleteExecutorByIDRequest{execID.String()}
	response, err := e.ComputeDeleteExecutorByIDEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return err
	}

	return nil
}

func (e Endpoints) ListExecutors(ctx context.Context) ([]*project.Executor, error) {
	request := ComputeListExecutorsRequest{}
	response, err := e.ComputeListExecutorsEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeListExecutorsResponse)
	return DecodeExecutorsBody2Executors(ExecutorsBody(res)), nil
}

func (e Endpoints) CreateProject(ctx context.Context, name string) (*project.Project, error) {
	request := ComputeCreateProjectRequest{name}
	response, err := e.ComputeCreateProjectEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeCreateProjectResponse)
	return DecodeProjectBody2Project(ProjectBody(res)), nil
}

func (e Endpoints) DeleteProjectByID(ctx context.Context, id project.ProjectID) error {
	request := ComputeDeleteProjectByIDRequest{id.String()}
	response, err := e.ComputeDeleteProjectByIDEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return err
	}

	return nil
}

func (e Endpoints) GetProjectByID(ctx context.Context, id project.ProjectID) (*project.Project, error) {
	request := ComputeGetProjectByIDRequest{id.String()}
	response, err := e.ComputeGetProjectByIDEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeGetProjectByIDResponse)
	return DecodeProjectBody2Project(ProjectBody(res)), nil
}

func (e Endpoints) ListProjects(ctx context.Context) ([]*project.Project, error) {
	request := ComputeListProjectsRequest{}
	response, err := e.ComputeListProjectsEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeListProjectsResponse)
	return DecodeProjectsBody2Projects(ProjectsBody(res)), nil
}

func (e Endpoints) CreateTask(ctx context.Context, projID project.ProjectID, execID project.ExecutorID, name string, data []byte) (*project.Task, error) {
	request := ComputeCreateTaskRequest{
		ProjectID:  projID.String(),
		ExecutorID: execID.String(),
		Name:       name,
		Data:       data,
	}

	response, err := e.ComputeCreateTaskEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeCreateTaskResponse)
	return DecodeTaskBody2Task(TaskBody(res)), nil
}

func (e Endpoints) DeleteTaskByID(ctx context.Context, id project.TaskID) error {
	return jerrors.NotImplementNowError
}

func (e Endpoints) GetTaskByID(ctx context.Context, id project.TaskID) (*project.Task, error) {
	return nil, jerrors.NotImplementNowError
}

func (e Endpoints) UpdateTaskByID(ctx context.Context, id project.TaskID, task *project.Task) (*project.Task, error) {
	request := ComputeUpdateTaskByIDRequest{
		ID:     id.String(),
		Status: task.Status.String(),
	}
	response, err := e.ComputeUpdateTaskByIDEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputeUpdateTaskByIDResponse)
	return DecodeTaskBody2Task(TaskBody(res)), nil
}

func (e Endpoints) PopReadyTask(ctx context.Context) (*project.Task, error) {
	request := ComputePopReadyTaskRequest{}
	response, err := e.ComputePopReadyTaskEndpoint(ctx, request)

	if err, ok := e.handleError(response, err); ok {
		return nil, err
	}

	res := response.(ComputePopReadyTaskResponse)
	return DecodeTaskBody2Task(TaskBody(res)), nil
}

func MakeComputeCreateExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeCreateExecutorRequest)
		e, err := s.CreateExecutor(ctx, req.Name, req.Pack, req.Data)
		if err != nil {
			return err, nil
		}
		return EncodeExecutor2ExecutorBody(e), nil
	}
}

func MakeComputeDeleteExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeDeleteExecutorByIDRequest)
		err := s.DeleteExecutorByID(ctx, project.ExecutorID(req.ID))
		if err != nil {
			return err, nil
		}
		return ComputeDeleteExecutorByIDResponse{}, nil
	}
}

func MakeComputeGetExecutorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeGetExecutorByIDRequest)
		e, err := s.GetExecutorByID(ctx, project.ExecutorID(req.ID))
		if err != nil {
			return err, nil
		}
		return EncodeExecutor2ExecutorBody(e), nil
	}
}

func MakeComputeListExecutorsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		es, err := s.ListExecutors(ctx)
		if err != nil {
			return err, nil
		}

		return EncodeExecutors2ExecutorsBody(es), nil
	}
}

func MakeComputeCreateProjectEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeCreateProjectRequest)
		p, err := s.CreateProject(ctx, req.Name)
		if err != nil {
			return err, nil
		}
		return EncodeProject2ProjectBody(p), nil
	}
}

func MakeComputeDeleteProjectByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeDeleteProjectByIDRequest)
		err := s.DeleteProjectByID(ctx, project.ProjectID(req.ID))
		if err != nil {
			return err, nil
		}
		return ComputeDeleteProjectByIDResponse{}, nil
	}
}

func MakeComputeGetProjectEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeGetProjectByIDRequest)
		p, err := s.GetProjectByID(ctx, project.ProjectID(req.ID))
		if err != nil {
			return err, nil
		}
		return EncodeProject2ProjectBody(p), nil
	}
}

func MakeComputeListProjectsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ps, err := s.ListProjects(ctx)
		if err != nil {
			return err, nil
		}

		return EncodeProjects2ProjectsBody(ps), nil
	}
}

func MakeComputeCreateTaskEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeCreateTaskRequest)
		t, err := s.CreateTask(ctx, project.ProjectID(req.ProjectID), project.ExecutorID(req.ExecutorID), req.Name, req.Data)
		if err != nil {
			return err, nil
		}
		return EncodeTask2TaskBody(t), nil
	}
}

func MakeComputeUpdateTaskEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeUpdateTaskByIDRequest)
		t := &project.Task{Status: project.LookupTaskStatus(req.Status)}
		t, err := s.UpdateTaskByID(ctx, project.TaskID(req.ID), t)
		if err != nil {
			return err, nil
		}

		return EncodeTask2TaskBody(t), nil
	}
}

func MakeComputeGetTaskEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeGetTaskByIDRequest)
		t, err := s.GetTaskByID(ctx, project.TaskID(req.ID))
		if err != nil {
			return err, nil
		}

		return EncodeTask2TaskBody(t), nil
	}
}

func MakeComputeDeleteTaskEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ComputeDeleteTaskByIDRequest)
		err := s.DeleteTaskByID(ctx, project.TaskID(req.ID))
		if err != nil {
			return err, nil
		}

		return ComputeDeleteTaskByIDResponse{}, nil
	}
}

func MakeComputePopReadyTaskEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, err := s.PopReadyTask(ctx)
		if err != nil {
			return err, nil
		}

		return EncodeTask2TaskBody(t), nil
	}
}
