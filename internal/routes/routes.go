package routes

import (
	"0day-backend/internal/middleware"
	routes "0day-backend/internal/routes/public"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.Logging)

	// Public API Routes
	publicGroup := app.Group("/api/v1/public")
	publicGroup.Post("/register", routes.Register)
}
