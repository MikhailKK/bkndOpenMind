package main

import (
	"my-gorm-fiber-app/database"
	"my-gorm-fiber-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize database connection
	database.ConnectDB()
	defer database.DB.Close()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
