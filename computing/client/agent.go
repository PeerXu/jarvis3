package client

import (
	"fmt"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	computing_endpoint "github.com/PeerXu/jarvis3/computing/endpoint"
	computing_service "github.com/PeerXu/jarvis3/computing/service"
	jmiddlewares "github.com/PeerXu/jarvis3/kit/middlewares/http"
	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	"github.com/PeerXu/jarvis3/utils"
)

type agentClient struct {
	computing_endpoint.Endpoints
	utils.Environment
}

func NewForAgent(instance string, environment utils.Environment, logger log.Logger) (computing_service.Service, error) {
	cli := agentClient{}

	fmt.Println("Agent Token: " + environment.Get("AgentToken"))

	opts := []httptransport.ClientOption{}
	encodeRequestFactory := jhttptransport.EncodeRequestFuncChain(
		jmiddlewares.LoadEnvironmentIntoContextMiddleware(environment),
		jmiddlewares.AssignAgentTokenToContextForClientMiddleware)
	decodeResponseFactory := jhttptransport.DecodeResponseFuncChain(
		jmiddlewares.DecodeErrorResponseMiddleware)

	eps, err := computing_endpoint.MakeClientEndpoints(instance, encodeRequestFactory, decodeResponseFactory, opts)
	cli.Endpoints = eps
	cli.Environment = environment

	return cli, err
}
