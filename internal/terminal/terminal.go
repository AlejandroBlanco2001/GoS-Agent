package terminal

import (
	parser "alejandroblanco2001/scanneros/internal/terminal/parser"
	"fmt"
	"os/exec"
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
		return nil, nil
	}

	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	if includeOutput {
		return out, nil
	}

	return nil, nil
}

func (t *Terminal) GetOpenConnections() []map[string]string {
	result, err := t.run(true, OpenConnections)

	if err != nil {
		fmt.Println("Error getting open connections: ", err)
		return nil
	}

	mapper := parser.ParseNetStatOutput(string(result))

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
