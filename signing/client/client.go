package client

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	jmiddlewares "github.com/PeerXu/jarvis3/kit/middlewares/http"
	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	signing_endpoint "github.com/PeerXu/jarvis3/signing/endpoint"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
	"github.com/PeerXu/jarvis3/utils"
)

type clientClient struct {
	signing_endpoint.Endpoints
	utils.Environment
}

func NewForClient(instance string, environment utils.Environment, logger log.Logger) (signing_service.Service, error) {
	cli := clientClient{Environment: environment}

	opts := []httptransport.ClientOption{}
	encodeRequestFactory := jhttptransport.EncodeRequestFuncChain(
		jmiddlewares.LoadEnvironmentIntoContextMiddleware(environment),
		jmiddlewares.AssignAccessTokenToContextForClientMiddleware,
	)
	decodeResponseFactory := jhttptransport.DecodeResponseFuncChain(
		jmiddlewares.DecodeErrorResponseMiddleware,
	)

	eps, err := signing_endpoint.MakeClientEndpoints(instance, encodeRequestFactory, decodeResponseFactory, opts)
	cli.Endpoints = eps

	return cli, err
}
