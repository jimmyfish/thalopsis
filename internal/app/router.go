package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmyfish/thalopsistematika/internal/app/middleware"
	"github.com/jimmyfish/thalopsistematika/internal/interfaces/health"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.HeadersMiddleware())

	healthGroup := r.Group("/health")
	{
		healthHandler := health.NewHandler()
		healthGroup.GET("/", healthHandler.CheckHealth)
	}
}
