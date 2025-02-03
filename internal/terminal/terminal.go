// Description: This file contains the terminal package that will be used to interact with the terminal and execute commands
package terminal

import (
	logger "alejandroblanco2001/scanneros/internal/platform/logger"
	parser "alejandroblanco2001/scanneros/internal/terminal/parser"
	"fmt"
	"os/exec"
	"runtime"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Terminal struct {
	OS                   string
	logger               *logger.EchoHandler
	EthernetAdapterNames []string
	CommandDictionary    map[string][]string
}

func returnDictionaryOfCommands() map[string][]string {
	if runtime.GOOS == "windows" {
		return map[string][]string{
			"OpenConnections":          OpenConnections,
			"OpenConnectionStatistics": OpenConnectionStatisticsPowerShell,
			"GetInterfaceNames":        GetInterfaceNames,
		}
	}

	return map[string][]string{
		"OpenConnections":          OpenConnectionsLinux,
		"OpenConnectionStatistics": OpenConnectionStatisticsLinux,
		"GetInterfaceNames":        GetInterfaceNamesLinux,
	}
}

func NewTerminal(lc fx.Lifecycle, logger *logger.EchoHandler) *Terminal {

	terminal := &Terminal{
		OS:                runtime.GOOS,
		logger:            logger,
		CommandDictionary: returnDictionaryOfCommands(),
	}

	logger.Log("Starting terminal, OS: " + terminal.OS)

	return terminal
}

func (t *Terminal) run(includeOutput bool, command string) (string, error) {
	if len(command) == 0 {
		t.logger.LogError("No command provided")
		return "", fmt.Errorf("no command provided")
	}

	if _, ok := t.CommandDictionary[command]; !ok {
		t.logger.LogError(fmt.Sprintf("Command %s not found in dictionary", command))
		return "", fmt.Errorf("command %s not found in dictionary", command)
	}

	commandToRun := t.CommandDictionary[command]

	out, err := exec.Command(commandToRun[0], commandToRun[1:]...).Output()

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
	result, err := t.run(true, "OpenConnections")

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get open connections: %v", err))
	}

	var mapper map[string]map[string]string

	if t.OS == "windows" {
		mapper = parser.ParseNetStatOutput(string(result))
	} else {
		mapper = parser.ParseNetStatOutputLinux(string(result))
	}

	for key, value := range mapper {
		t.logger.Log(fmt.Sprintf("Connection %s: ", key))

		for k, v := range value {
			t.logger.Log(fmt.Sprintf("Key: %s, Value: %s", k, v))
		}
	}
}

func (t *Terminal) GetOpenConnectionStatistics() map[string]map[string]int64 {
	result, err := t.run(true, "OpenConnectionStatistics")
	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get open connection statistics: %v", err))
		return nil
	}

	var mapper map[string]map[string]int64

	if t.OS == "windows" {
		mapper, err = parser.ParseNetAdapterStatistics(string(result), t.EthernetAdapterNames)
	} else {
		mapper, err = parser.ParseNetAdapterStatisticsLinux(string(result), t.EthernetAdapterNames)
	}

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to parse open connection statistics: %v", err))
		return nil
	}

	t.logStatistics(mapper)

	return mapper
}

func (t *Terminal) logStatistics(mapper map[string]map[string]int64) {
	for key, value := range mapper {
		t.logger.Log(fmt.Sprintf("Adapter %s Statistics:", key))
		t.logger.Log(fmt.Sprintf("Received bytes: %f MB", BytesToMB(value["ReceivedBytes"])))
		t.logger.Log(fmt.Sprintf("Sent bytes: %f MB", BytesToMB(value["SentBytes"])))
	}
}

func (t *Terminal) GetInterfaceNames() ([]string, error) {
	result, err := t.run(true, "GetInterfaceNames")

	if err != nil {
		t.logger.LogError(fmt.Sprintf("Failed to get interface names: %v", err))
		return nil, nil
	}

	if t.OS == "windows" {
		if _, err := parser.ParseInterfaceNames(string(result)); err != nil {
			t.logger.LogError(fmt.Sprintf("Failed to parse interface names: %v", err))
			return nil, nil
		}

		interfaceNames, _ := parser.ParseInterfaceNames(string(result))

		t.EthernetAdapterNames = interfaceNames
	} else {
		if _, err := parser.ParseInterfaceNamesLinux(string(result)); err != nil {
			t.logger.LogError(fmt.Sprintf("Failed to parse interface names: %v", err))
			return nil, nil
		}

		interfaceNames, _ := parser.ParseInterfaceNamesLinux(string(result))

		t.EthernetAdapterNames = interfaceNames
	}

	return t.EthernetAdapterNames, nil
}

func (t *Terminal) Stop() {
	t.logger.Log("Stopping terminal")
	panic("Stop")
}

var Module = fx.Options(
	fx.Provide(NewTerminal, logger.NewEchoHandler, zap.NewExample),
)
