package sdk

import (
	"os"

	"github.com/go-kit/kit/log"

	computing_client "github.com/PeerXu/jarvis3/computing/client"
	computing_service "github.com/PeerXu/jarvis3/computing/service"
	"github.com/PeerXu/jarvis3/utils"
)

func NewComputingClient(instance string, env utils.Environment) (computing_service.Service, error) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)

	return computing_client.NewForAgent(instance, env, logger)
}
