package logger

import "go.uber.org/zap"

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{
		log: log,
	}
}

func (h *EchoHandler) Log(message string) {
	h.log.Info(message)
}
