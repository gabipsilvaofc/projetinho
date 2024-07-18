package main

import (
	"teste/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)
	app.Run("localhost:5000")
}