package health

import (
	"time"

	"github.com/gin-gonic/gin"
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
