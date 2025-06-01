package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"new-blog/boot"
	"new-blog/cmd"
)

func main() {
	cmd.Execute()
	fx.New(
		boot.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
	).Run()
}
