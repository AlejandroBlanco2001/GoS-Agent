package main

import (
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			// Remove the default logger from the fx module (it's not needed)
			fx.NopLogger,
		),
	).Run()
}
