package notfound

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
