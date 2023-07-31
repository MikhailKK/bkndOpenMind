package main

import (
	"bkndOpenMind/database"
	"bkndOpenMind/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize database connection
	database.ConnectDB()

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":3000")
}
