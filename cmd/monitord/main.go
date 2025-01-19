package monitord

import (
	terminal "alejandroblanco2001/scanneros/internal/terminal"
	"context"
	"runtime"

	"go.uber.org/fx"
)

func StartMonitor(lc fx.Lifecycle) *terminal.Terminal {
	os := runtime.GOOS

	t := terminal.NewTerminal(os)

	if t == nil {
		panic("Error creating terminal")
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			t.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			t.Stop()
			return nil
		},
	})

	return t
}
