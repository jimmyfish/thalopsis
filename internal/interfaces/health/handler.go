package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
