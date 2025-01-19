package main

import (
	monitor "alejandroblanco2001/scanneros/cmd/monitord"
	terminal "alejandroblanco2001/scanneros/internal/terminal"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(monitor.NewMonitor, zap.NewExample),
		fx.Invoke(func(*terminal.Terminal) {}),
	).Run()
}
