package main

import (
	terminal "alejandroblanco2001/scanneros/internal/terminal"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			terminal.Module,
			fx.NopLogger,
		),
	).Run()
}
