package commandhandler

import (
	"alejandroblanco2001/scanneros/internal/platform/logger"

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
	h.logger.Log("Handling command")
}

var Module = fx.Options(
	fx.Provide(NewCommandHandler, logger.NewEchoHandler, zap.NewExample),
	fx.Invoke(func(c *CommandHandler) {
		go c.HandleCommand()
	}),
)
