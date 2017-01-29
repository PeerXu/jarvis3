package computing

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/computing/encode_decode"
	. "github.com/PeerXu/jarvis3/computing/endpoint"
	. "github.com/PeerXu/jarvis3/computing/service"
	jkth "github.com/PeerXu/jarvis3/kit/transport/http"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
)

func MakeHandler(ctx context.Context, s Service, ss signing_service.Service) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(EncodeError),
	}

	withAuthorizationOptions := append(
		opts,
		kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
	)

	withRemoteValidateTokenEndpointFactory := endpoint.Chain(
		RemoteValidateTokenMiddleware(ss),
	)

	var computeCreateExecutorHandler *kithttp.Server
	{
		computeCreateExecutorEndpoint := MakeComputeCreateExecutorEndpoint(s)
		computeCreateExecutorEndpoint = withRemoteValidateTokenEndpointFactory(computeCreateExecutorEndpoint)
		computeCreateExecutorOptions := withAuthorizationOptions
		computeCreateExecutorHandler = kithttp.NewServer(
			ctx,
			computeCreateExecutorEndpoint,
			DecodeComputeCreateExecutorRequest,
			MakeResponseEncoder(http.StatusOK),
			computeCreateExecutorOptions...,
		)
	}

	var computeDeleteExecutorHandler *kithttp.Server
	{
		computeDeleteExecutorEndpoint := MakeComputeDeleteExecutorEndpoint(s)
		computeDeleteExecutorEndpoint = withRemoteValidateTokenEndpointFactory(computeDeleteExecutorEndpoint)
		computeDeleteExecutorOptions := withAuthorizationOptions
		computeDeleteExecutorHandler = kithttp.NewServer(
			ctx,
			computeDeleteExecutorEndpoint,
			DecodeComputeDeleteExecutorRequest,
			MakeResponseEncoder(http.StatusNoContent),
			computeDeleteExecutorOptions...,
		)
	}

	var computeGetExecutorHandler *kithttp.Server
	{
		computeGetExecutorEndpoint := MakeComputeGetExecutorEndpoint(s)
		computeGetExecutorEndpoint = withRemoteValidateTokenEndpointFactory(computeGetExecutorEndpoint)
		computeGetExecutorOptions := withAuthorizationOptions
		computeGetExecutorHandler = kithttp.NewServer(
			ctx,
			computeGetExecutorEndpoint,
			DecodeComputeGetExecutorByIDRequest,
			MakeResponseEncoder(http.StatusOK),
			computeGetExecutorOptions...,
		)
	}

	var computeListExecutorsHandler *kithttp.Server
	{
		computeListExecutorsEndpoint := MakeComputeListExecutorsEndpoint(s)
		computeListExecutorsEndpoint = withRemoteValidateTokenEndpointFactory(computeListExecutorsEndpoint)
		computeListExecutorsOptions := withAuthorizationOptions
		computeListExecutorsHandler = kithttp.NewServer(
			ctx,
			computeListExecutorsEndpoint,
			DecodeComputeListExecutorsRequest,
			MakeResponseEncoder(http.StatusOK),
			computeListExecutorsOptions...,
		)
	}

	var computeCreateProjectHandler *kithttp.Server
	{
		computeCreateProjectEndpoint := MakeComputeCreateProjectEndpoint(s)
		computeCreateProjectEndpoint = withRemoteValidateTokenEndpointFactory(computeCreateProjectEndpoint)
		computeCreateProjectOptions := withAuthorizationOptions
		computeCreateProjectHandler = kithttp.NewServer(
			ctx,
			computeCreateProjectEndpoint,
			DecodeComputeCreateProjectRequest,
			MakeResponseEncoder(http.StatusOK),
			computeCreateProjectOptions...,
		)
	}

	var computeDeleteProjectByIDHandler *kithttp.Server
	{
		computeDeleteProjectByIDEndpoint := MakeComputeDeleteProjectByIDEndpoint(s)
		computeDeleteProjectByIDEndpoint = withRemoteValidateTokenEndpointFactory(computeDeleteProjectByIDEndpoint)
		computeDeleteProjectByIDOptions := withAuthorizationOptions
		computeDeleteProjectByIDHandler = kithttp.NewServer(
			ctx,
			computeDeleteProjectByIDEndpoint,
			DecodeComputeDeleteProjectByIDRequest,
			MakeResponseEncoder(http.StatusNoContent),
			computeDeleteProjectByIDOptions...,
		)
	}

	var computeGetProjectHandler *kithttp.Server
	{
		computeGetProjectEndpoint := MakeComputeGetProjectEndpoint(s)
		computeGetProjectEndpoint = withRemoteValidateTokenEndpointFactory(computeGetProjectEndpoint)
		computeGetProjectOptions := withAuthorizationOptions
		computeGetProjectHandler = kithttp.NewServer(
			ctx,
			computeGetProjectEndpoint,
			DecodeComputeGetProjectRequest,
			MakeResponseEncoder(http.StatusOK),
			computeGetProjectOptions...,
		)
	}

	var computeListProjectsHandler *kithttp.Server
	{
		computeListProjectsEndpoint := MakeComputeListProjectsEndpoint(s)
		computeListProjectsEndpoint = withRemoteValidateTokenEndpointFactory(computeListProjectsEndpoint)
		computeListProjectsOptions := withAuthorizationOptions
		computeListProjectsHandler = kithttp.NewServer(
			ctx,
			computeListProjectsEndpoint,
			DecodeComputeListProjectsRequest,
			MakeResponseEncoder(http.StatusOK),
			computeListProjectsOptions...,
		)
	}

	var computeCreateTaskHandler *kithttp.Server
	{
		computeCreateTaskEndpoint := MakeComputeCreateTaskEndpoint(s)
		computeCreateTaskEndpoint = withRemoteValidateTokenEndpointFactory(computeCreateTaskEndpoint)
		computeCreateTaskOptions := withAuthorizationOptions
		computeCreateTaskHandler = kithttp.NewServer(
			ctx,
			computeCreateTaskEndpoint,
			DecodeComputeCreateTaskRequest,
			MakeResponseEncoder(http.StatusOK),
			computeCreateTaskOptions...,
		)
	}

	var computeUpdateTaskHandler *kithttp.Server
	{
		computeUpdateTaskEndpoint := MakeComputeUpdateTaskEndpoint(s)
		computeUpdateTaskEndpoint = withRemoteValidateTokenEndpointFactory(computeUpdateTaskEndpoint)
		computeUpdateTaskOptions := withAuthorizationOptions
		computeUpdateTaskHandler = kithttp.NewServer(
			ctx,
			computeUpdateTaskEndpoint,
			DecodeComputeUpdateTaskRequest,
			MakeResponseEncoder(http.StatusOK),
			computeUpdateTaskOptions...,
		)
	}

	var computeGetTaskHandler *kithttp.Server
	{
		computeGetTaskEndpoint := MakeComputeGetTaskEndpoint(s)
		computeGetTaskEndpoint = withRemoteValidateTokenEndpointFactory(computeGetTaskEndpoint)
		computeGetTaskOptions := withAuthorizationOptions
		computeGetTaskHandler = kithttp.NewServer(
			ctx,
			computeGetTaskEndpoint,
			DecodeComputeGetTaskRequest,
			MakeResponseEncoder(http.StatusOK),
			computeGetTaskOptions...,
		)
	}

	var computeDeleteTaskHandler *kithttp.Server
	{
		computeDeleteTaskEndpoint := MakeComputeDeleteTaskEndpoint(s)
		computeDeleteTaskEndpoint = withRemoteValidateTokenEndpointFactory(computeDeleteTaskEndpoint)
		computeDeleteTaskOptions := withAuthorizationOptions
		computeDeleteTaskHandler = kithttp.NewServer(
			ctx,
			computeDeleteTaskEndpoint,
			DecodeComputeDeleteTaskRequest,
			MakeResponseEncoder(http.StatusNoContent),
			computeDeleteTaskOptions...,
		)
	}

	var computePopReadyTaskHandler *kithttp.Server
	{
		computePopReadyTaskEndpoint := MakeComputePopReadyTaskEndpoint(s)
		computePopReadyTaskEndpoint = withRemoteValidateTokenEndpointFactory(computePopReadyTaskEndpoint)
		computePopReadyTaskOptions := withAuthorizationOptions
		computePopReadyTaskHandler = kithttp.NewServer(
			ctx,
			computePopReadyTaskEndpoint,
			DecodeComputePopReadyTaskRequest,
			MakeResponseEncoder(http.StatusOK),
			computePopReadyTaskOptions...,
		)
	}

	r.Handle("/computing/v1/executors", computeCreateExecutorHandler).Methods("POST")
	r.Handle("/computing/v1/executors/{executor_id}", computeDeleteExecutorHandler).Methods("DELETE")
	r.Handle("/computing/v1/executors/{executor_id}", computeGetExecutorHandler).Methods("GET")
	r.Handle("/computing/v1/executors", computeListExecutorsHandler).Methods("GET")
	r.Handle("/computing/v1/projects", computeCreateProjectHandler).Methods("POST")
	r.Handle("/computing/v1/projects/{project_id}", computeDeleteProjectByIDHandler).Methods("DELETE")
	r.Handle("/computing/v1/projects/{project_id}", computeGetProjectHandler).Methods("GET")
	r.Handle("/computing/v1/projects", computeListProjectsHandler).Methods("GET")
	r.Handle("/computing/v1/tasks", computeCreateTaskHandler).Methods("POST")
	r.Handle("/computing/v1/tasks/{task_id}", computeUpdateTaskHandler).Methods("PATCH")
	r.Handle("/computing/v1/tasks/{task_id}", computeGetTaskHandler).Methods("GET")
	r.Handle("/computing/v1/tasks/{task_id}", computeDeleteTaskHandler).Methods("DELETE")
	r.Handle("/computing/v1/readyTasks", computePopReadyTaskHandler).Methods("DELETE")

	return r
}
