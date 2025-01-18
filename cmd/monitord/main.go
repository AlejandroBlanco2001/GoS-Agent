package monitord

import (
	terminal "alejandroblanco2001/scanneros/internal/terminal"
	"fmt"
	"runtime"
	"time"
)

func Run() {
	os := runtime.GOOS

	fmt.Println("--------------------------------------------------")
	fmt.Println("Running monitord")
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
	fmt.Println("OS: ", os)
	fmt.Println("--------------------------------------------------")

	t := terminal.NewTerminal(os)
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go monitor(t, ticker, quit)

	time.Sleep(10 * time.Second)

	close(quit)

	fmt.Println("--------------------------------------------------")
	fmt.Println("monitord finished")
	fmt.Println("--------------------------------------------------")
}

func monitor(t *terminal.Terminal, ticker *time.Ticker, quit chan struct{}) {
	fmt.Println("Checking open connections")
	for {
		select {
		case <-ticker.C:
			if err := t.GetOpenConnections(); err != nil {
				fmt.Println("Error checking open connections: ", err)
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}
