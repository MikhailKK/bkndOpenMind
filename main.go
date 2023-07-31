package main

import (
	// "bkndOpenMind/database"
	// "bkndOpenMind/routes"

	"bkndOpenMind/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome this API")
}

func main() {
	// Initialize database connection
	database.ConnectDB()
	app := fiber.New()

	app.Get("/api", welcome)

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	// routes.SetupRoutes(app)

	// Start the server
	// app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))
}
