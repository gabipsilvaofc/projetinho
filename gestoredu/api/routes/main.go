package routes

import (

	"teste/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(router *fiber.App) *fiber.App {
	v1 := router.Group("/v1")
	userController := controllers.NewUserController()
	{
		v1.Delete("/user/:id", userController.Delete)
	}
	return router
}