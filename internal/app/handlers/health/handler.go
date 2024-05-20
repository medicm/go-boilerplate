package health

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct{}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/health", h.HandleHealthCheck)
}

func (h *Handler) HandleHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    "UP",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
