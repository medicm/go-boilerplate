package generic

import "boilerplate/internal/interfaces"

type HandlerFactory struct {
	createFunc func() (interfaces.IHandler, error)
}

func NewHandlerFactory(createFunc func() (interfaces.IHandler, error)) *HandlerFactory {
	return &HandlerFactory{
		createFunc: createFunc,
	}
}

func (f *HandlerFactory) CreateHandler() (interfaces.IHandler, error) {
	return f.createFunc()
}
