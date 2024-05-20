package handlers

import (
	"boilerplate/internal/app/handlers/health"
	"boilerplate/internal/app/handlers/notfound"
	"boilerplate/internal/interfaces"
	"boilerplate/internal/pkg/middlewares/appcontext"
	"go.uber.org/fx"
	"sync"
)

const (
	HealthHandler   = "health"
	NotFoundHandler = "not_found"
)

type HandlerRegistry struct {
	registry map[string]interfaces.IHandlerFactory
	mux      sync.Mutex
	Ctx      *appcontext.Context
}

func NewHandlerRegistry(ctx *appcontext.Context) *HandlerRegistry {
	return &HandlerRegistry{
		registry: make(map[string]interfaces.IHandlerFactory),
		Ctx:      ctx,
	}
}

func (hr *HandlerRegistry) RegisterHandler(name string, factory interfaces.IHandlerFactory) {
	hr.mux.Lock()
	defer hr.mux.Unlock()
	hr.registry[name] = factory
}

func (hr *HandlerRegistry) CreateHandlers() ([]interfaces.IHandler, error) {
	hr.mux.Lock()
	defer hr.mux.Unlock()

	handlers := make([]interfaces.IHandler, 0, len(hr.registry))
	for _, factory := range hr.registry {
		handler, err := factory.CreateHandler()
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}
	return handlers, nil
}

func (hr *HandlerRegistry) SetupHandlers() {
	hr.Ctx.Log.Info("Setting up Handlers...")
	hr.RegisterHandler(HealthHandler, health.NewHealthHandlerFactory())
	hr.RegisterHandler(NotFoundHandler, notfound.NewNotFoundHandlerFactory())
}

var Module = fx.Options(fx.Provide(NewHandlerRegistry))
