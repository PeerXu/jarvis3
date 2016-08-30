package repository

import (
	"sync"

	"github.com/PeerXu/jarvis3/project"
)

type projectRepository struct {
	mtx       sync.RWMutex
	projects  map[string]map[string]*project.Project
	executors map[string]map[string]*project.Executor
}

func NewProjectRepository() *projectRepository {
	return &projectRepository{
		mtx:       sync.RWMutex{},
		projects:  map[string]map[string]*project.Project{},
		executors: map[string]map[string]*project.Executor{},
	}
}

func (r *projectRepository) CreateExecutor(e *project.Executor) (*project.Executor, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var userExecutors map[string]*project.Executor
	var ok bool

	if userExecutors, ok = r.executors[e.Owner]; !ok {
		userExecutors = map[string]*project.Executor{}
		r.executors[e.Owner] = userExecutors
	}

	userExecutors[e.Name] = e
	return e, nil
}

func (r *projectRepository) DeleteExecutor(owner string, name string) (err error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var userExecutors map[string]*project.Executor
	var ok bool

	if userExecutors, ok = r.executors[owner]; !ok {
		err = project.ErrOwnerNotFound
	}

	if _, ok = userExecutors[name]; !ok {
		err = project.ErrExecutorNotFound

	}

	delete(userExecutors, name)

	return
}

func (r *projectRepository) GetExecutor(owner string, name string) (*project.Executor, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if userExecutors, ok := r.executors[owner]; ok {
		if executor, ok := userExecutors[name]; ok {
			return executor, nil
		}
		return nil, project.ErrExecutorNotFound
	}

	return nil, project.ErrOwnerNotFound
}

func (r *projectRepository) ListExecutors(owner string) ([]*project.Executor, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if userExecutors, ok := r.executors[owner]; ok {
		var executors []*project.Executor
		for _, e := range userExecutors {
			executors = append(executors, e)
		}
		return executors, nil
	}

	return nil, project.ErrOwnerNotFound
}

func (r *projectRepository) CreateProject(p *project.Project) (*project.Project, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var userProjects map[string]*project.Project
	var ok bool

	if userProjects, ok = r.projects[p.Owner]; !ok {
		userProjects = map[string]*project.Project{}
		r.projects[p.Owner] = userProjects
	}

	userProjects[p.Name] = p
	return p, nil
}

func (r *projectRepository) DeleteProject(owner string, name string) (err error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	var userProjects map[string]*project.Project
	var ok bool

	if userProjects, ok = r.projects[owner]; !ok {
		return project.ErrOwnerNotFound
	}

	if _, ok = userProjects[name]; !ok {
		return project.ErrProjectNotFound
	}

	delete(userProjects, name)
	return nil
}

func (r *projectRepository) GetProject(owner string, name string) (*project.Project, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	var userProjects map[string]*project.Project
	var proj *project.Project
	var ok bool

	if userProjects, ok = r.projects[owner]; !ok {
		return nil, project.ErrOwnerNotFound
	}

	if proj, ok = userProjects[name]; !ok {
		return nil, project.ErrProjectNotFound
	}

	return proj, nil
}

func (r *projectRepository) ListProjects(owner string) ([]*project.Project, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	var userProjects map[string]*project.Project
	var projects []*project.Project
	var ok bool

	if userProjects, ok = r.projects[owner]; !ok {
		return nil, project.ErrOwnerNotFound
	}

	for _, p := range userProjects {
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *projectRepository) CreateJob(owner string, proj string, job *project.Job) (*project.Job, error) {
	p, err := r.GetProject(owner, proj)
	if err != nil {
		return nil, err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	p.Jobs = append(p.Jobs, job)
	return job, nil
}

func (r *projectRepository) GetJob(owner string, proj string, name string) (*project.Job, error) {
	p, err := r.GetProject(owner, proj)
	if err != nil {
		return nil, err
	}

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	var job *project.Job
	for _, j := range p.Jobs {
		if j.Name == name {
			job = j
			break
		}
	}

	if job == nil {
		return nil, project.ErrJobNotFound
	}

	return job, nil
}

func (r *projectRepository) ListJobs(owner string, proj string) ([]*project.Job, error) {
	p, err := r.GetProject(owner, proj)
	if err != nil {
		return nil, err
	}

	r.mtx.RLock()
	defer r.mtx.RUnlock()

	return p.Jobs, nil
}

func (r *projectRepository) UpdateJob(owner string, name string, proj string, job *project.Job) (*project.Job, error) {
	j, err := r.GetJob(owner, proj, name)
	if err != nil {
		return nil, err
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	if job.Status != project.JobStatus_Unknown {
		j.Status = job.Status
	}

	return j, nil
}
