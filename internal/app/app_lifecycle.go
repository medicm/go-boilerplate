package app

import (
	"boilerplate/internal/pkg/middlewares/appcontext"
	"context"
	"go.uber.org/dig"
	"go.uber.org/fx"
)

type MainParams struct {
	dig.In
	Ctx *appcontext.Context
}

func StartApp(mp MainParams, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			mp.Ctx.Log.Info("Application Started!")
			return nil
		},
		OnStop: func(context.Context) error {
			mp.Ctx.Log.Info("Application Stopped!")
			return nil
		},
	})
}
