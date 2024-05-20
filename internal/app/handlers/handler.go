package handlers

import (
	"boilerplate/internal/pkg/middlewares/appcontext"
	"go.uber.org/dig"
)

type HandlerParams struct {
	dig.In
	Hr      *HandlerRegistry
	Context *appcontext.Context
}

func SetupHandlers(hr *HandlerRegistry) {
	hr.SetupHandlers()
}
