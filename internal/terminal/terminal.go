package terminal

import (
	logger "alejandroblanco2001/scanneros/internal/platform/logger"
	parser "alejandroblanco2001/scanneros/internal/terminal/parser"
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Terminal struct {
	OS     string
	logger *logger.EchoHandler
}

func NewTerminal(lc fx.Lifecycle, logger *logger.EchoHandler) *Terminal {
	terminal := &Terminal{
		OS:     runtime.GOOS,
		logger: logger,
	}

	logger.Log("Starting terminal, OS: " + terminal.OS)

	lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go terminal.Start()
			return nil
		},
		OnStop: func(context context.Context) error {
			terminal.Stop()
			return nil
		},
	})

	return terminal
}

func (t *Terminal) run(includeOutput bool, command []string) ([]byte, error) {
	if len(command) == 0 {
		t.logger.LogError("No command provided")
		return nil, fmt.Errorf("no command provided")
	}

	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to execute command: %v, output: %s", err, out))
		return nil, fmt.Errorf("failed to execute command: %v, output: %s", err, out)
	}

	if includeOutput {
		return out, nil
	}

	return nil, nil
}

func (t *Terminal) GetOpenConnections() map[string]map[string]string {
	result, err := t.run(true, OpenConnections)

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get open connections: %v", err))
		return nil
	}

	mapper := parser.ParseNetStatOutput(string(result))

	for key, value := range mapper {
		if value["state"] == "ESTABLISHED" {
			t.logger.Log(fmt.Sprintf("Established connection %s with protocol %s", key, value["protocol"]))
		}
	}

	return mapper
}

func (t *Terminal) Start() {
	for {
		_ = t.GetOpenConnections()
		time.Sleep(15 * time.Second)
	}
}

func (t *Terminal) Stop() {
	t.logger.Log("Stopping terminal")
	panic("Stop")
}

var Module = fx.Options(
	fx.Provide(NewTerminal, logger.NewEchoHandler, zap.NewExample),
	fx.Invoke(func(t *Terminal) {
		go t.Start()
	}),
)
