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
		fmt.Println("No command provided")
		return nil, fmt.Errorf("no command provided")
	}

	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		fmt.Println("Error: ", err)
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
		fmt.Println("Error getting open connections: ", err)
		return nil
	}

	mapper := parser.ParseNetStatOutput(string(result))

	for key, value := range mapper {
		if value["state"] == "ESTABLISHED" {
			fmt.Printf("Established connection %s with protocol %s\n", key, value["protocol"])
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
	fmt.Println("Stopping terminal")
	panic("Stop")
}

var Module = fx.Options(
	fx.Provide(NewTerminal, logger.NewEchoHandler, zap.NewExample),
	fx.Invoke(func(t *Terminal) {
		go t.Start()
	}),
)
