package main

import (
	"context"

	"boilerplate/internal/app"
	"boilerplate/internal/app/handlers"
	"boilerplate/internal/pkg/config"
	"boilerplate/internal/pkg/logger"
	"boilerplate/internal/pkg/middlewares/appcontext"
	"boilerplate/internal/pkg/webserver"
	"go.uber.org/fx"
)

func main() {
	appContext := &appcontext.Context{
		Log: logger.NewLogger(),
	}

	fxApp := fx.New(
		config.Module,
		fx.Provide(func() *appcontext.Context {
			return appContext
		}),
		webserver.Module,
		handlers.Module,
		fx.Invoke(app.StartApp, handlers.SetupHandlers, webserver.RegisterHooks),
		fx.WithLogger(logger.NewFxLoggerWrapper),
	)

	startErr := fxApp.Start(context.Background())
	if startErr != nil {
		panic(startErr)
	}

	<-fxApp.Done()

	stopErr := fxApp.Stop(context.Background())
	if stopErr != nil {
		panic(stopErr)
	}
}
