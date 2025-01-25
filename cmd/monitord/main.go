package main

import (
	terminal "alejandroblanco2001/scanneros/internal/terminal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			terminal.Module,
			// Remove the default logger from the fx module (it's not needed)
			fx.NopLogger,
		),
	).Run()
}
