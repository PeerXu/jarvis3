package v1

import (
	"math/rand"

	"golang.org/x/net/context"

	"github.com/PeerXu/jarvis3/agent/driver"
	"github.com/PeerXu/jarvis3/agent/sdk"
	computing_service "github.com/PeerXu/jarvis3/computing/service"
	"github.com/PeerXu/jarvis3/container"
	"github.com/PeerXu/jarvis3/project"
	"github.com/PeerXu/jarvis3/utils"
	jvs_ctx "github.com/PeerXu/jarvis3/utils/context"
	jvs_log "github.com/PeerXu/jarvis3/utils/log"
)

type clients struct {
	Computing computing_service.Service
}

type agentDriver struct {
	logger  jvs_log.Logger
	clients clients
}

func NewAgentDriver(ctx jvs_ctx.Context, logger jvs_log.Logger) driver.AgentDriver {
	env := utils.NewEnvironment()

	env.Set("AgentToken", ctx.Metadata().Get("AgentToken"))
	host := ctx.Metadata().Get("ComputingService-Host")

	computingClient, err := sdk.NewComputingClient(host, env)
	if err != nil {
		logger.Log("message", "new computing client failed", "error", err)
		panic(err)
	}

	clis := clients{
		Computing: computingClient,
	}

	return &agentDriver{
		logger:  logger,
		clients: clis,
	}
}

func (drv *agentDriver) FetchTask() (*project.Task, error) {
	ctx := context.Background()
	task, err := drv.clients.Computing.PopReadyTask(ctx)
	if err != nil {
		drv.logger.Log("message", "fetch task failed", "error", err)
		return nil, err
	}

	return task, nil
}

func (drv *agentDriver) CompleteTask(task *project.Task) error {
	ctx := context.Background()
	task, err := drv.clients.Computing.UpdateTaskByID(ctx, task.ID, task)
	if err != nil {
		drv.logger.Log("message", "complete task failed", "error", err)
		return err
	}
	return nil
}

var randomLetters = "012345679abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomCodeGenerator(n int) string {
	cs := make([]byte, n)
	for i := 0; i < n; i++ {
		cs[i] = randomLetters[rand.Intn(len(randomLetters))]
	}
	return string(cs)
}

func (drv *agentDriver) CreateContainer(task *project.Task) (container.Container, error) {
	srv_ctx := context.Background()
	cntr_ctx := container.NewContext()

	exec, err := drv.clients.Computing.GetExecutorByID(srv_ctx, task.ExecutorID)
	if err != nil {
		drv.logger.Log("message", "get executor failed", "error", err)
		return nil, err
	}

	proj, err := drv.clients.Computing.GetProjectByID(srv_ctx, task.ProjectID)
	if err != nil {
		drv.logger.Log("message", "get project failed", "error", err)
		return nil, err
	}

	cntr_ctx.Executor.Path = exec.Pack
	cntr_ctx.Executor.Data = string(exec.Data)
	cntr_ctx.Executor.Metadata["ID"] = exec.ID.String()
	cntr_ctx.Executor.Metadata["Name"] = exec.Name
	cntr_ctx.Executor.Metadata["Owner-ID"] = exec.OwnerID.String()

	cntr_ctx.Container.Code = randomCodeGenerator(16)

	// cntr_ctx.Request.Data = string(task.Data)
	cntr_ctx.Request.Metadata["Task-ID"] = task.ID.String()
	cntr_ctx.Request.Metadata["Task-Name"] = task.Name
	cntr_ctx.Request.Metadata["Project-ID"] = proj.ID.String()
	cntr_ctx.Request.Metadata["Project-Name"] = proj.Name

	cntr := container.NewContainer(cntr_ctx)

	return cntr, nil
}

func (drv *agentDriver) CloseContainer(cntr container.Container) error {
	cntr.Close()
	return nil
}
