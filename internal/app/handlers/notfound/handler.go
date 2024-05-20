package notfound

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct{}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.NoRoute(h.HandleNotFound)
}

func (h *Handler) HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Sorry, the requested resource was not found.",
	})
}
