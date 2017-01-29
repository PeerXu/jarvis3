package client

import (
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	jmiddlewares "github.com/PeerXu/jarvis3/kit/middlewares/http"
	jhttptransport "github.com/PeerXu/jarvis3/kit/transport/http"
	signing_endpoint "github.com/PeerXu/jarvis3/signing/endpoint"
	signing_service "github.com/PeerXu/jarvis3/signing/service"
)

type serviceClient struct {
	signing_endpoint.Endpoints
}

func NewForService(instance string, logger log.Logger) (signing_service.Service, error) {
	cli := serviceClient{}

	encodeRequestFactory := jhttptransport.EncodeRequestFuncChain(
		jmiddlewares.LoadRequestIntoContextMiddleware(),
		jmiddlewares.AssignAccessTokenToContextForServiceMiddleware,
	)
	decodeResponseFactory := jhttptransport.DecodeResponseFuncChain(
		jmiddlewares.DecodeErrorResponseMiddleware,
	)
	opts := []httptransport.ClientOption{}

	eps, err := signing_endpoint.MakeClientEndpoints(instance, encodeRequestFactory, decodeResponseFactory, opts)
	cli.Endpoints = eps

	return cli, err
}
