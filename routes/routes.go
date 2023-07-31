package routes

import (
	"my-gorm-fiber-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", handlers.GetAllUsers)
	// Add more routes here
}
