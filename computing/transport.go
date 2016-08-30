package computing

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	jkth "github.com/PeerXu/jarvis3/kit/transport/http"
	"github.com/PeerXu/jarvis3/project"
	"github.com/PeerXu/jarvis3/signing"
)

func MakeHandler(ctx context.Context, s Service, ss signing.Service) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	var computeCreateExecutorHandler *kithttp.Server
	{
		computeCreateExecutorEndpoint := makeComputeCreateExecutorEndpoint(s)
		computeCreateExecutorEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeCreateExecutorEndpoint)
		computeCreateExecutorOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeCreateExecutorHandler = kithttp.NewServer(
			ctx,
			computeCreateExecutorEndpoint,
			decodeComputeCreateExecutorRequest,
			makeResponseEncoder(http.StatusOK),
			computeCreateExecutorOptions...,
		)
	}

	var computeDeleteExecutorHandler *kithttp.Server
	{
		computeDeleteExecutorEndpoint := makeComputeDeleteExecutorEndpoint(s)
		computeDeleteExecutorEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeDeleteExecutorEndpoint)
		computeDeleteExecutorOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeDeleteExecutorHandler = kithttp.NewServer(
			ctx,
			computeDeleteExecutorEndpoint,
			decodeComputeDeleteExecutorRequest,
			makeResponseEncoder(http.StatusNoContent),
			computeDeleteExecutorOptions...,
		)
	}

	var computeGetExecutorHandler *kithttp.Server
	{
		computeGetExecutorEndpoint := makeComputeGetExecutorEndpoint(s)
		computeGetExecutorEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeGetExecutorEndpoint)
		computeGetExecutorOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeGetExecutorHandler = kithttp.NewServer(
			ctx,
			computeGetExecutorEndpoint,
			decodeComputeGetExecutorRequest,
			makeResponseEncoder(http.StatusOK),
			computeGetExecutorOptions...,
		)
	}

	var computeListExecutorsHandler *kithttp.Server
	{
		computeListExecutorsEndpoint := makeComputeListExecutorsEndpoint(s)
		computeListExecutorsEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeListExecutorsEndpoint)
		computeListExecutorsOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeListExecutorsHandler = kithttp.NewServer(
			ctx,
			computeListExecutorsEndpoint,
			decodeComputeListExecutorsRequest,
			makeResponseEncoder(http.StatusOK),
			computeListExecutorsOptions...,
		)
	}

	var computeCreateProjectHandler *kithttp.Server
	{
		computeCreateProjectEndpoint := makeComputeCreateProjectEndpoint(s)
		computeCreateProjectEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeCreateProjectEndpoint)
		computeCreateProjectOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeCreateProjectHandler = kithttp.NewServer(
			ctx,
			computeCreateProjectEndpoint,
			decodeComputeCreateProjectRequest,
			makeResponseEncoder(http.StatusOK),
			computeCreateProjectOptions...,
		)
	}

	var computeDeleteProjectHandler *kithttp.Server
	{
		computeDeleteProjectEndpoint := makeComputeDeleteProjectEndpoint(s)
		computeDeleteProjectEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeDeleteProjectEndpoint)
		computeDeleteProjectOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeDeleteProjectHandler = kithttp.NewServer(
			ctx,
			computeDeleteProjectEndpoint,
			decodeComputeDeleteProjectRequest,
			makeResponseEncoder(http.StatusNoContent),
			computeDeleteProjectOptions...,
		)
	}

	var computeGetProjectHandler *kithttp.Server
	{
		computeGetProjectEndpoint := makeComputeGetProjectEndpoint(s)
		computeGetProjectEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeGetProjectEndpoint)
		computeGetProjectOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeGetProjectHandler = kithttp.NewServer(
			ctx,
			computeGetProjectEndpoint,
			decodeComputeGetProjectRequest,
			makeResponseEncoder(http.StatusOK),
			computeGetProjectOptions...,
		)
	}

	var computeListProjectsHandler *kithttp.Server
	{
		computeListProjectsEndpoint := makeComputeListProjectsEndpoint(s)
		computeListProjectsEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeListProjectsEndpoint)
		computeListProjectsOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeListProjectsHandler = kithttp.NewServer(
			ctx,
			computeListProjectsEndpoint,
			decodeComputeListProjectsRequest,
			makeResponseEncoder(http.StatusOK),
			computeListProjectsOptions...,
		)
	}

	var computeCreateJobHandler *kithttp.Server
	{
		computeCreateJobEndpoint := makeComputeCreateJobEndpoint(s)
		computeCreateJobEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeCreateJobEndpoint)
		computeCreateJobOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeCreateJobHandler = kithttp.NewServer(
			ctx,
			computeCreateJobEndpoint,
			decodeComputeCreateJobRequest,
			makeResponseEncoder(http.StatusOK),
			computeCreateJobOptions...,
		)
	}

	var computeUpdateJobHandler *kithttp.Server
	{
		computeUpdateJobEndpoint := makeComputeUpdateJobEndpoint(s)
		computeUpdateJobEndpoint = endpoint.Chain(
			RemoteValidateAccessTokenMiddleware(ss),
		)(computeUpdateJobEndpoint)
		computeUpdateJobOptions := append(
			opts,
			kithttp.ServerBefore(jkth.ExtractAuthorizationFromRequestHeader),
		)
		computeUpdateJobHandler = kithttp.NewServer(
			ctx,
			computeUpdateJobEndpoint,
			decodeComputeUpdateJobRequest,
			makeResponseEncoder(http.StatusOK),
			computeUpdateJobOptions...,
		)
	}

	r.Handle("/computing/v1/executors", computeCreateExecutorHandler).Methods("POST")
	r.Handle("/computing/v1/executors/{name}", computeDeleteExecutorHandler).Methods("DELETE")
	r.Handle("/computing/v1/executors/{name}", computeGetExecutorHandler).Methods("GET")
	r.Handle("/computing/v1/executors", computeListExecutorsHandler).Methods("GET")
	r.Handle("/computing/v1/projects", computeCreateProjectHandler).Methods("POST")
	r.Handle("/computing/v1/projects/{name}", computeDeleteProjectHandler).Methods("DELETE")
	r.Handle("/computing/v1/projects/{name}", computeGetProjectHandler).Methods("GET")
	r.Handle("/computing/v1/projects", computeListProjectsHandler).Methods("GET")
	r.Handle("/computing/v1/projects/{project}/jobs", computeCreateJobHandler).Methods("POST")
	r.Handle("/computing/v1/projects/{project}/jobs/{name}", computeUpdateJobHandler).Methods("PATCH")

	return r
}

func decodeComputeCreateExecutorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
		Pack string `json:"pack"`
		Data []byte `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return computeCreateExecutorRequest{
		Name: body.Name,
		Pack: body.Pack,
		Data: body.Data,
	}, nil
}

func decodeComputeDeleteExecutorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return computeDeleteExecutorRequest{Name: name}, nil
}

func decodeComputeGetExecutorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return computeGetExecutorRequest{Name: name}, nil
}

func decodeComputeListExecutorsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return computeListExecutorsRequest{}, nil
}

func decodeComputeCreateProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}

	return computeCreateProjectRequest{Name: body.Name}, nil
}

func decodeComputeDeleteProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return computeDeleteProjectRequest{Name: name}, nil
}

func decodeComputeGetProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	return computeGetProjectRequest{Name: name}, nil
}

func decodeComputeListProjectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return computeListProjectsRequest{}, nil
}

func decodeComputeCreateJobRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name     string `json:"name"`
		Project  string `json:"-"`
		Executor string `json:"executor"`
		Data     []byte `json:"data"`
	}

	vars := mux.Vars(r)
	project := vars["project"]

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}
	body.Project = project

	return computeCreateJobRequest{
		Name:     body.Name,
		Project:  body.Project,
		Executor: body.Executor,
		Data:     body.Data,
	}, nil
}

func decodeComputeUpdateJobRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name    string `json:"-"`
		Project string `json:"-"`
		Status  string `json:"status"`
	}

	vars := mux.Vars(r)
	proj := vars["project"]
	name := vars["name"]

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}

	if body.Status == "" {
		body.Status = project.JobStatus_Unknown.String()
	}
	body.Name = name
	body.Project = proj

	return computeUpdateJobRequest{
		Name:    body.Name,
		Project: body.Project,
		Status:  body.Status,
	}, nil
}

func makeResponseEncoder(code int) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if err, ok := response.(error); ok {
			encodeError(ctx, err, w)
			return nil
		}
		return encodeResponse(ctx, w, code, response)
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, code int, response interface{}) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	if code != http.StatusNoContent {
		return json.NewEncoder(w).Encode(response)
	}

	return nil
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	cerr, ok := err.(*computeError)
	if !ok {
		cerr = newComputeError(errorServerError, "unknown error", err)
	}

	var code int
	switch cerr.Type {
	case errorInvalidRequest:
		code = http.StatusBadRequest
	case errorNotFound:
		code = http.StatusNotFound
	case errorServerError:
	default:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(cerr)
}
