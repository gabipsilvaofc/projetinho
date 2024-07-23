package main

import (
	"log"
	
	"teste/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.AppRoutes(app)

	log.Fatal(app.Listen(":5000"))
}