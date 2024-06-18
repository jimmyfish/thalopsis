package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jimmyfish/thalopsistematika/internal/app"
)

func main() {

	r := gin.Default()

	app.RegisterRoutes(r)

	r.Run(":8080")
}
