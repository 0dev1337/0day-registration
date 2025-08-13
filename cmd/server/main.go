package main

import (
	"0day-backend/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":3001")
}
