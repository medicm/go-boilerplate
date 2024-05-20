package interfaces

import (
	"github.com/gin-gonic/gin"
)

type (
	IHandler interface {
		RegisterRoutes(r *gin.Engine)
	}
	IHandlerFactory interface {
		CreateHandler() (IHandler, error)
	}
)
