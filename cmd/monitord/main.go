// Purpose: Main entry point for the monitord application.
package main

import (
	zero "alejandroblanco2001/scanneros/cmd/monitord/internal/zeromq"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			zero.Module,
			// Remove the default logger from the fx module (it's not needed),
			fx.NopLogger,
		),
	).Run()
}
