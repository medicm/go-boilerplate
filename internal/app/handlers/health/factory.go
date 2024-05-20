package health

import (
	"boilerplate/internal/app/handlers/generic"
	"boilerplate/internal/interfaces"
)

func NewHandlerInstance() (interfaces.IHandler, error) {
	return &Handler{}, nil
}

func NewHealthHandlerFactory() interfaces.IHandlerFactory {
	return generic.NewHandlerFactory(NewHandlerInstance)
}
