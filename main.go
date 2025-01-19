package main

import (
	monitor "alejandroblanco2001/scanneros/cmd/monitord"
	terminal "alejandroblanco2001/scanneros/internal/terminal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(monitor.StartMonitor),
		fx.Invoke(func(*terminal.Terminal) {}),
	).Run()
}
