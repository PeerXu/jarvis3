package repository

import (
	"sync"

	"github.com/PeerXu/jarvis3/project"
	"github.com/PeerXu/jarvis3/user"
)

type projectRepository struct {
	mtx          sync.RWMutex
	schMtx       sync.Mutex
	executors    map[project.ExecutorID]*project.Executor
	projects     map[project.ProjectID]*project.Project
	tasks        map[project.TaskID]*project.Task
	readyTasks   map[project.TaskID]*project.Task
	runningTasks map[project.TaskID]*project.Task
}

func NewProjectRepository() *projectRepository {
	return &projectRepository{
		mtx:          sync.RWMutex{},
		executors:    map[project.ExecutorID]*project.Executor{},
		projects:     map[project.ProjectID]*project.Project{},
		tasks:        map[project.TaskID]*project.Task{},
		readyTasks:   map[project.TaskID]*project.Task{},
		runningTasks: map[project.TaskID]*project.Task{},
	}
}

func (r *projectRepository) CreateExecutor(e *project.Executor) (*project.Executor, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.executors[e.ID] = e

	return e, nil
}

func (r *projectRepository) DeleteExecutorByID(id project.ExecutorID) error {
	_, err := r.GetExecutorByID(id)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.executors, id)

	return nil
}

func (r *projectRepository) GetExecutorByID(id project.ExecutorID) (*project.Executor, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if e, ok := r.executors[id]; ok {
		return e, nil
	}

	return nil, project.ErrExecutorNotFound
}

func (r *projectRepository) ListExecutors(userID user.UserID) ([]*project.Executor, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	executors := []*project.Executor{}

	for _, exec := range r.executors {
		if exec.OwnerID == userID {
			executors = append(executors, exec)
		}
	}

	return executors, nil
}

func (r *projectRepository) CreateProject(p *project.Project) (*project.Project, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	r.projects[p.ID] = p
	return p, nil
}

func (r *projectRepository) DeleteProjectByID(id project.ProjectID) error {
	_, err := r.GetProjectByID(id)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	delete(r.projects, id)

	return nil
}

func (r *projectRepository) GetProjectByID(id project.ProjectID) (*project.Project, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if proj, ok := r.projects[id]; ok {
		return proj, nil
	}

	return nil, project.ErrProjectNotFound
}

func (r *projectRepository) ListProjects(userID user.UserID) ([]*project.Project, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	projects := []*project.Project{}

	for _, proj := range r.projects {
		if proj.OwnerID == userID {
			projects = append(projects, proj)
		}
	}

	return projects, nil
}

func (r *projectRepository) CreateTask(task *project.Task) (*project.Task, error) {
	p, err := r.GetProjectByID(task.ProjectID)
	if err != nil {
		return nil, err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	p.Tasks = append(p.Tasks, task)
	r.tasks[task.ID] = task
	return task, nil
}

func (r *projectRepository) GetTaskByID(id project.TaskID) (*project.Task, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if j, ok := r.tasks[id]; ok {
		return j, nil
	}
	return nil, project.ErrTaskNotFound
}

func (r *projectRepository) ListTasksByProjectID(projID project.ProjectID) ([]*project.Task, error) {
	p, err := r.GetProjectByID(projID)
	if err != nil {
		return nil, err
	}

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	return p.Tasks, nil
}

func (r *projectRepository) UpdateTaskByID(id project.TaskID, task *project.Task) (*project.Task, error) {
	t, err := r.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	if task.Status == project.TaskStatus_Ready {
		delete(r.runningTasks, t.ID)
		r.readyTasks[t.ID] = t
	}

	if task.Status == project.TaskStatus_Running {
		delete(r.readyTasks, t.ID)
		r.runningTasks[t.ID] = t
	}

	if task.Status == project.TaskStatus_Stop || task.Status == project.TaskStatus_Error {
		if t.Status != project.TaskStatus_Running {
			return nil, project.ErrTaskNotRunning
		}
		delete(r.runningTasks, t.ID)
	}

	if task.Status != project.TaskStatus_Unknown {
		t.Status = task.Status
	}

	return t, nil
}

func (r *projectRepository) DeleteTaskByID(id project.TaskID) error {
	task, err := r.GetTaskByID(id)
	if err != nil {
		return err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	p, err := r.GetProjectByID(task.ProjectID)
	if err != nil {
		return err
	}

	var ts []*project.Task
	for _, t := range p.Tasks {
		if t.ID != task.ID {
			ts = append(ts, t)
		}
	}
	p.Tasks = ts

	delete(r.tasks, id)

	return nil
}

func (r *projectRepository) PopReadyTask() (*project.Task, error) {
	r.schMtx.Lock()
	defer r.schMtx.Unlock()

	if len(r.readyTasks) == 0 {
		return nil, project.ErrNotReadyTask
	}

	var task *project.Task
	var err error
	// FIXME(Peer): nofair scheduler
	for id := range r.readyTasks {
		task = r.readyTasks[id]
		task.Status = project.TaskStatus_Running
		task, err = r.UpdateTaskByID(task.ID, task)
		break
	}

	if err != nil {
		return nil, err
	}
	return task, nil
}
