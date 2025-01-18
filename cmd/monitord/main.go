package monitord

import (
	terminal "alejandroblanco2001/scanneros/internal/terminal"
	"fmt"
	"runtime"
)

func Run() {
	os := runtime.GOOS

	fmt.Println("--------------------------------------------------")
	fmt.Println("Running monitord")
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("OS: ", os)
	fmt.Println("--------------------------------------------------")

	t := terminal.NewTerminal(os)

	err := terminal.GetNetStat(t)

	if err != nil {
		fmt.Println("Error running netstat command: ", err)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("monitord finished")
	fmt.Println("--------------------------------------------------")
}
