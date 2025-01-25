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

func (h *EchoHandler) LogError(message string) {
	h.log.Error(message)
}

func (h *EchoHandler) LogFatal(message string) {
	h.log.Fatal(message)
}

func (h *EchoHandler) LogWarn(message string) {
	h.log.Warn(message)
}
