package terminal

import (
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

func (t *Terminal) run(includeOutput bool, command []string) error {
	if len(command) == 0 {
		fmt.Println("No command provided")
		return nil
	}

	out, err := exec.Command(command[0], command[1:]...).Output()

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	if includeOutput {
		fmt.Println("Output: ", string(out))
	}

	return nil
}

func (t *Terminal) GetOpenConnections() error {
	return t.run(true, OpenConnections)
}

func (t *Terminal) Start() {
	for {
		t.GetOpenConnections()
	}
}

func (t *Terminal) Stop() {
	fmt.Println("Stopping terminal")
	panic("Stop")
}
