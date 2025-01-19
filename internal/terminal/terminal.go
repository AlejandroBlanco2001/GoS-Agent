package terminal

import (
	parser "alejandroblanco2001/scanneros/internal/terminal/parser"
	"fmt"
	"os/exec"
	"time"
)

type Terminal struct {
	OS string
}

func NewTerminal(os string) *Terminal {
	return &Terminal{
		OS: os,
	}
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

	time.Sleep(5 * time.Second)

	return mapper
}

func (t *Terminal) Start() {
	for {
		result := t.GetOpenConnections()
		fmt.Println(result)
	}
}

func (t *Terminal) Stop() {
	fmt.Println("Stopping terminal")
	panic("Stop")
}
