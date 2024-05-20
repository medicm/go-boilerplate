package appcontext

import (
	"boilerplate/internal/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Context struct {
	Log *logger.Logger
}

type AppContext struct {
	*gin.Context
	AppContext *Context
}

const ContextName = "AppContext"

func InjectContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.NewLogger()
		appContext := &Context{
			Log: log,
		}
		ginContext := &AppContext{
			Context:    c,
			AppContext: appContext,
		}
		c.Set(ContextName, ginContext)
		c.Next()
	}
}

func GetContext(c *gin.Context) (*Context, error) {
	val, exists := c.Get(ContextName)
	if !exists {
		return nil, fmt.Errorf("appContext not found in gin.Context")
	}

	// Debugging line
	log.Printf("Type of val: %T\n", val)

	ctx, ok := val.(*Context)
	if !ok {
		return nil, fmt.Errorf("Context has an incorrect type")
	}

	return ctx, nil
}
