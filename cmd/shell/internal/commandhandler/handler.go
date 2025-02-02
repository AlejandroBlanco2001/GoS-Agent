package commandhandler

import (
	zero "alejandroblanco2001/scanneros/cmd/shell/internal/zeromq"
	"alejandroblanco2001/scanneros/internal/platform/logger"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type CommandHandler struct {
	logger *logger.EchoHandler
}

func NewCommandHandler(logger *logger.EchoHandler) *CommandHandler {
	return &CommandHandler{
		logger: logger,
	}
}

func (h *CommandHandler) HandleCommand() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			h.logger.Log("Handling command...")
			zero.NewClient("Hello")
		}
	}()

	select {}
}

var Module = fx.Options(
	fx.Provide(NewCommandHandler, logger.NewEchoHandler, zap.NewExample),
	fx.Invoke(func(c *CommandHandler) {
		go c.HandleCommand()
	}),
)
