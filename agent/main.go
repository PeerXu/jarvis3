package main

import (
	"os"

	"github.com/go-kit/kit/log"

	agent_driver "github.com/PeerXu/jarvis3/agent/driver/v1"
	jarvis_context "github.com/PeerXu/jarvis3/utils/context"
)

func main() {
	drv_ctx := jarvis_context.NewContext()
	drv_ctx.Metadata().Set("AgentToken", "THIVo7U56x5YScYHfbtXIvVxeiiaJm+XZHmBmY6+qJwLSzc5cBFegu1vQSXI+nMR5Nfe+pItqud4Zmf36TbNTw==")
	drv_ctx.Metadata().Set("ComputingService-Host", "localhost:27182")

	drv_logger := log.NewContext(log.NewLogfmtLogger(os.Stdout)).With("layer", "driver", "name", "v1")
	drv := agent_driver.NewAgentDriver(drv_ctx, drv_logger)

	eng_ctx := jarvis_context.NewContext()
	eng_logger := log.NewContext(log.NewLogfmtLogger(os.Stdout)).With("layer", "engine")
	eng := NewAgentEngine(drv, eng_ctx, eng_logger)

	eng.Launch()
}
