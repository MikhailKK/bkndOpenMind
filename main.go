package main

import (
	"bkndOpenMind/database"
	"bkndOpenMind/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome this API")
}

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	// Initialize database connection
	database.ConnectDB()

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	// app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))
}
