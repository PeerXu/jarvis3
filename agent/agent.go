package main

import (
	"time"

	"github.com/PeerXu/jarvis3/agent/driver"
	"github.com/PeerXu/jarvis3/project"
	jvs_ctx "github.com/PeerXu/jarvis3/utils/context"
	jvs_log "github.com/PeerXu/jarvis3/utils/log"
)

type AgentEngine interface {
	Launch() error
	Terminate() error
}

type agentEngine struct {
	logger jvs_log.Logger

	Driver  driver.AgentDriver
	Context jvs_ctx.Context

	terminated bool
}

func NewAgentEngine(drv driver.AgentDriver, ctx jvs_ctx.Context, logger jvs_log.Logger) AgentEngine {
	return &agentEngine{
		logger: logger,

		Driver:  drv,
		Context: ctx,

		terminated: false,
	}
}

func (agt *agentEngine) Launch() error {
	for {
		time.Sleep(1 * time.Second)
		if agt.terminated {
			agt.logger.Log("message", "Engine terminated")
			break
		}

		task, err := agt.Driver.FetchTask()
		if err != nil {
			agt.logger.Log("message", "fetch task failed", "error", err)
			continue
		}

		cntr, err := agt.Driver.CreateContainer(task)
		if err != nil {
			agt.logger.Log("message", "create container failed", "error", err)
			continue
		}

		err = cntr.Run()
		if err != nil {
			task.Status = project.TaskStatus_Error
		} else {
			task.Status = project.TaskStatus_Stop
		}
		agt.logger.Log("message", "task executed", "status", task.Status.String(), "error", err)

		err = agt.Driver.CloseContainer(cntr)
		if err != nil {
			agt.logger.Log("message", "close container failed", "error", err)
		}

		err = agt.Driver.CompleteTask(task)
		if err != nil {
			agt.logger.Log("message", "complete task failed", "error", err)
		}
	}
	return nil
}

func (agt *agentEngine) Terminate() error {
	agt.terminated = true
	return nil
}
