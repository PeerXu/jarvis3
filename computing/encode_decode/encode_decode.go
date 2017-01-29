package encode_decode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	. "github.com/PeerXu/jarvis3/computing/error"
	jerrors "github.com/PeerXu/jarvis3/errors"
	"github.com/PeerXu/jarvis3/project"
	"github.com/PeerXu/jarvis3/user"
)

func EncodeComputeCreateExecutorRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/executors").Methods("POST")
	r.Method, r.URL.Path = "POST", "/computing/v1/executors"
	return EncodeRequest(ctx, r, request)
}

func EncodeComputeListExecutorsRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/executors").Methods("GET")
	r.Method, r.URL.Path = "GET", "/computing/v1/executors"
	return nil
}

func EncodeComputeDeleteExecutorByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/executors/{executor_id}").Methods("DELETE")
	req := request.(ComputeDeleteExecutorByIDRequest)
	r.Method, r.URL.Path = "DELETE", "/computing/v1/executors/"+req.ID
	return nil
}

func EncodeComputeGetExecutorByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/executors/{executor_id}").Methods("GET")
	req := request.(ComputeGetExecutorByIDRequest)
	r.Method, r.URL.Path = "GET", "/computing/v1/executors/"+req.ID
	return nil
}

func EncodeComputeCreateProjectRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/projects").Methods("POST")
	r.Method, r.URL.Path = "POST", "/computing/v1/projects"
	return EncodeRequest(ctx, r, request)
}

func EncodeComputeListProjectsRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/projects").Methods("GET")
	r.Method, r.URL.Path = "GET", "/computing/v1/projects"
	return EncodeRequest(ctx, r, request)
}

func EncodeComputeDeleteProjectByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/projects/{project_id}").Methods("DELETE")
	req := request.(ComputeDeleteProjectByIDRequest)
	r.Method, r.URL.Path = "DELETE", "/computing/v1/projects/"+req.ID
	return nil
}

func EncodeComputeCreateTaskRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/tasks").Methods("POST")
	r.Method, r.URL.Path = "POST", "/computing/v1/tasks"
	return EncodeRequest(ctx, r, request)
}

func EncodeComputePopReadyTaskRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/readyTasks").Methods("DELETE")
	r.Method, r.URL.Path = "DELETE", "/computing/v1/readyTasks"
	return nil
}

func EncodeComputeGetProjectByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/projects/{project_id}").Methods("GET")
	req := request.(ComputeGetProjectByIDRequest)
	r.Method, r.URL.Path = "GET", "/computing/v1/projects/"+req.ID
	return nil
}

func EncodeComputeUpdateTaskByIDRequest(ctx context.Context, r *http.Request, request interface{}) error {
	// r.Path("/computing/v1/tasks/{task_id}").Methods("PATCH")
	req := request.(ComputeUpdateTaskByIDRequest)
	r.Method, r.URL.Path = "PATCH", "/computing/v1/tasks/"+req.ID
	return EncodeRequest(ctx, r, request)
}

func DecodeComputeCreateExecutorResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeCreateExecutorResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeListExecutorsResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeListExecutorsResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeDeleteExecutorByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	return ComputeDeleteExecutorByIDResponse{}, nil
}

func DecodeComputeGetExecutorByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeGetExecutorByIDResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeCreateProjectResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeCreateProjectResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeListProjectsResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeListProjectsResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeDeleteProjectByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeDeleteProjectByIDResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeCreateTaskResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeCreateTaskResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputePopReadyTaskResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputePopReadyTaskResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeGetProjectByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeGetProjectByIDResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeUpdateTaskByIDResponse(ctx context.Context, res *http.Response) (interface{}, error) {
	var response ComputeUpdateTaskByIDResponse
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func DecodeComputeCreateExecutorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string `json:"name"`
		Pack string `json:"pack"`
		Data []byte `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return ComputeCreateExecutorRequest{
		Name: body.Name,
		Pack: body.Pack,
		Data: body.Data,
	}, nil
}

func DecodeComputeDeleteExecutorRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["executor_id"]
	return ComputeDeleteExecutorByIDRequest{ID: id}, nil
}

func DecodeComputeGetExecutorByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["executor_id"]
	return ComputeGetExecutorByIDRequest{ID: id}, nil
}

func DecodeComputeListExecutorsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return ComputeListExecutorsRequest{}, nil
}

func DecodeComputeCreateProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name string
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}

	return ComputeCreateProjectRequest{Name: body.Name}, nil
}

func DecodeComputeDeleteProjectByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["project_id"]
	return ComputeDeleteProjectByIDRequest{ID: id}, nil
}

func DecodeComputeGetProjectRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id := vars["project_id"]
	return ComputeGetProjectByIDRequest{ID: id}, nil
}

func DecodeComputeListProjectsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return ComputeListProjectsRequest{}, nil
}

func DecodeComputeCreateTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name       string `json:"name"`
		ProjectID  string `json:"project_id"`
		ExecutorID string `json:"executor_id"`
		Data       []byte `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}

	return ComputeCreateTaskRequest{
		Name:       body.Name,
		ProjectID:  body.ProjectID,
		ExecutorID: body.ExecutorID,
		Data:       body.Data,
	}, nil
}

func DecodeComputeUpdateTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		ID     string `json:"-"`
		Status string `json:"status"`
	}

	vars := mux.Vars(r)
	taskID := vars["task_id"]

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return err, nil
	}

	if body.Status == "" {
		body.Status = project.TaskStatus_Unknown.String()
	}

	body.ID = taskID

	return ComputeUpdateTaskByIDRequest{
		ID:     body.ID,
		Status: body.Status,
	}, nil
}

func DecodeComputeGetTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	taskID := vars["task_id"]

	return ComputeGetTaskByIDRequest{ID: taskID}, nil
}

func DecodeComputeDeleteTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	taskID := vars["task_id"]

	return ComputeDeleteTaskByIDRequest{ID: taskID}, nil
}

func DecodeComputePopReadyTaskRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return ComputePopReadyTaskRequest{}, nil
}

func MakeResponseEncoder(code int) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
		if err, ok := response.(error); ok {
			EncodeError(ctx, err, w)
			return nil
		}
		return EncodeResponse(ctx, w, code, response)
	}
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, code int, response interface{}) error {
	if code == 0 {
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	if code != http.StatusNoContent {
		return json.NewEncoder(w).Encode(response)
	}

	return nil
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	jerr, ok := err.(jerrors.JarvisError)
	if !ok {
		jerr = NewComputeError(jerrors.ErrorServerError, "unknown error", err)
	}

	code := http.StatusInternalServerError
	switch jerr.Type {
	case jerrors.ErrorAccessDenied:
		code = http.StatusUnauthorized
	case jerrors.ErrorInvalidRequest:
		code = http.StatusBadRequest
	case jerrors.ErrorNotFound:
		code = http.StatusNotFound
	case jerrors.ErrorServerError:
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(jerr)
}

func EncodeRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type ExecutorBody struct {
	ID      string `json:"id"`
	OwnerID string `json:"owner_id"`
	Name    string `json:"name"`
	Pack    string `json:"pack"`
	Data    []byte `json:"data"`
}

type ExecutorsBody []ExecutorBody

func EncodeExecutor2ExecutorBody(e *project.Executor) ExecutorBody {
	return ExecutorBody{
		ID:      e.ID.String(),
		OwnerID: e.OwnerID.String(),
		Name:    e.Name,
		Pack:    e.Pack,
		Data:    e.Data,
	}
}

func DecodeExecutorBody2Executor(b ExecutorBody) *project.Executor {
	return &project.Executor{
		ID:      project.ExecutorID(b.ID),
		OwnerID: user.UserID(b.OwnerID),
		Name:    b.Name,
		Pack:    b.Pack,
		Data:    b.Data,
	}
}

func EncodeExecutors2ExecutorsBody(es []*project.Executor) ExecutorsBody {
	bs := ExecutorsBody{}
	for _, e := range es {
		bs = append(bs, EncodeExecutor2ExecutorBody(e))
	}
	return bs
}

func DecodeExecutorsBody2Executors(bs ExecutorsBody) []*project.Executor {
	es := []*project.Executor{}
	for _, b := range []ExecutorBody(bs) {
		es = append(es, DecodeExecutorBody2Executor(b))
	}
	return es
}

type TaskBody struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	ExecutorID string `json:"executor_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Data       []byte `json:"data"`
}

func EncodeTask2TaskBody(j *project.Task) TaskBody {
	return TaskBody{
		ID:         j.ID.String(),
		ProjectID:  j.ProjectID.String(),
		ExecutorID: j.ExecutorID.String(),
		Name:       j.Name,
		Status:     j.Status.String(),
		Data:       j.Data,
	}
}

func DecodeTaskBody2Task(b TaskBody) *project.Task {
	return &project.Task{
		ID:         project.TaskID(b.ID),
		ProjectID:  project.ProjectID(b.ProjectID),
		ExecutorID: project.ExecutorID(b.ExecutorID),
		Name:       b.Name,
		Data:       []byte(b.Data),
	}
}

type ProjectBody struct {
	ID      string     `json:"id"`
	OwnerID string     `json:"owner_id"`
	Name    string     `json:"name"`
	Tasks   []TaskBody `json:"tasks"`
}

type ProjectsBody []ProjectBody

func EncodeProject2ProjectBody(p *project.Project) ProjectBody {
	var js []TaskBody
	for _, j := range p.Tasks {
		js = append(js, EncodeTask2TaskBody(j))
	}
	return ProjectBody{
		ID:      p.ID.String(),
		Name:    p.Name,
		OwnerID: p.OwnerID.String(),
		Tasks:   js,
	}
}

func EncodeProjects2ProjectsBody(ps []*project.Project) ProjectsBody {
	bs := ProjectsBody{}

	for _, p := range ps {
		bs = append(bs, EncodeProject2ProjectBody(p))
	}

	return bs
}

func DecodeProjectBody2Project(b ProjectBody) *project.Project {
	return &project.Project{
		ID:   project.ProjectID(b.ID),
		Name: b.Name,
	}
}

func DecodeProjectsBody2Projects(bs ProjectsBody) []*project.Project {
	ps := []*project.Project{}
	for _, b := range []ProjectBody(bs) {
		ps = append(ps, DecodeProjectBody2Project(b))
	}
	return ps
}

type ComputeCreateExecutorRequest struct {
	Name string `json:"name"`
	Pack string `json:"pack"`
	Data []byte `json:"data"`
}

type ComputeCreateExecutorResponse ExecutorBody

type ComputeDeleteExecutorByIDRequest struct {
	ID string `json:"-"`
}

type ComputeDeleteExecutorByIDResponse struct{}

type ComputeGetExecutorByIDRequest struct {
	ID string `json:"-"`
}

type ComputeGetExecutorByIDResponse ExecutorBody

type ComputeListExecutorsRequest struct{}

type ComputeListExecutorsResponse ExecutorsBody

type ComputeCreateProjectRequest struct {
	Name string
}

type ComputeCreateProjectResponse ProjectBody

type ComputeDeleteProjectByIDRequest struct {
	ID string
}

type ComputeDeleteProjectByIDResponse struct{}

type ComputeGetProjectByIDRequest struct {
	ID string
}

type ComputeGetProjectByIDResponse ProjectBody

type ComputeListProjectsRequest struct{}

type ComputeListProjectsResponse []ProjectBody

type ComputeCreateTaskRequest struct {
	Name       string `json:"name"`
	ProjectID  string `json:"project_id"`
	ExecutorID string `json:"executor_id"`
	Data       []byte `json:"data"`
}

type ComputeCreateTaskResponse TaskBody

type ComputeUpdateTaskByIDRequest struct {
	ID     string
	Status string
}

type ComputeUpdateTaskByIDResponse TaskBody

type ComputeGetTaskByIDRequest struct {
	ID string
}

type ComputeGetTaskByIDResponse TaskBody

type ComputeDeleteTaskByIDRequest struct {
	ID string
}

type ComputeDeleteTaskByIDResponse struct{}

type ComputePopReadyTaskRequest struct{}

type ComputePopReadyTaskResponse TaskBody
