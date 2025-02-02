package main

import (
	"alejandroblanco2001/scanneros/cmd/shell/internal/commandhandler"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			commandhandler.Module,
			// Remove the default logger from the fx module (it's not needed)
			fx.NopLogger,
		),
	).Run()
}
