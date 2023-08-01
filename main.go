package main

import (
	// "bkndOpenMind/routes"

	"bkndOpenMind/config"
	"bkndOpenMind/database"
	"bkndOpenMind/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get(config.ConcatenateStrings(config.PathApi), routes.FirstEP)
	app.Post(config.ConcatenateStrings(config.PathApi, config.PathU), routes.CreateUser)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathU), routes.GetUsers)

	// Add more routes here

}

func main() {
	// Initialize database connection
	database.ConnectDB()
	app := fiber.New()

	// app.Get("/api", routes.Welcome)

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	SetupRoutes(app)

	// Start the server
	// app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))
}
