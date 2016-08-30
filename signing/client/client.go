package client

import (
	"github.com/go-kit/kit/log"

	"github.com/PeerXu/jarvis3/signing"
	"github.com/PeerXu/jarvis3/utils"
)

type client struct {
	signing.Endpoints
	utils.Environment
}

func New(instance string, environment utils.Environment, logger log.Logger) (signing.Service, error) {
	cli := client{Environment: environment}
	eps, err := signing.MakeClientEndpoints(instance, cli.Environment)
	cli.Endpoints = eps
	return cli, err
}
