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
	OS                   string
	logger               *logger.EchoHandler
	EthernetAdapterNames []string
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

func (t *Terminal) run(includeOutput bool, command []string) (string, error) {
	if len(command) == 0 {
		t.logger.LogError("No command provided")
		return "", fmt.Errorf("no command provided")
	}

	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to execute command: %v, output: %s", err, out))
		return "", fmt.Errorf("failed to execute command: %v, output: %s", err, out)
	}

	if includeOutput {
		return RemoveOutputCommandPrefix(out), nil
	}

	return "", nil
}

func (t *Terminal) GetOpenConnections() {
	result, err := t.run(true, OpenConnections)

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get open connections: %v", err))
	}

	mapper := parser.ParseNetStatOutput(string(result))

	for key, value := range mapper {
		if value["state"] == "ESTABLISHED" {
			t.logger.Log(fmt.Sprintf("Established connection %s with protocol %s", key, value["protocol"]))
		}
	}
}

func (t *Terminal) GetOpenConnectionStatistics() {
	result, err := t.run(true, OpenConnectionStatisticsPowerShell)

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get open connection statistics: %v", err))
	}

	mapper, err := parser.ParseNetAdapterStatistics(string(result))

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to parse open connection statistics: %v", err))
	}

	t.logger.Log(fmt.Sprintf("Open connection statistics: %v", mapper))

	for key, value := range mapper {
		t.logger.Log(fmt.Sprintf("Adapter %s Statistics: ", key))
		t.logger.Log(fmt.Sprintf("Recieved bytes: %f mb", BytesToMB(value["ReceivedBytes"])))
		t.logger.Log(fmt.Sprintf("Recieved unicast packets: %f mb", BytesToMB(value["ReceivedUnicastPackets"])))
		t.logger.Log(fmt.Sprintf("Sent bytes: %f mb", BytesToMB(value["SentBytes"])))
		t.logger.Log(fmt.Sprintf("Sent unicast packets: %f mb", BytesToMB(value["SentUnicastPackets"])))
	}
}

func (t *Terminal) GetInterfaceNames() ([]string, error) {
	result, err := t.run(true, GetInterfaceNames)

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get interface names: %v", err))
		return nil, nil
	}

	interfaceNames, err := parser.ParseInterfaceNames(result)

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to parse interface names: %v", err))
		return nil, nil
	}

	// Just to avoid calling the command again for every use of the interface names
	t.EthernetAdapterNames = interfaceNames

	return interfaceNames, nil
}

func (t *Terminal) Start() {
	for {
		t.GetOpenConnectionStatistics()
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
