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
	// EP for Users
	app.Get(config.ConcatenateStrings(config.PathApi), routes.FirstEP)
	app.Post(config.ConcatenateStrings(config.PathApi, config.PathU), routes.CreateUser)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathU), routes.GetUsers)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), routes.GetUser)
	app.Put(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), routes.UpdateUser)
	app.Delete(config.ConcatenateStrings(config.PathApi, config.PathU, config.PathID), routes.DeleteUser)
	// Ep for Product
	app.Post(config.ConcatenateStrings(config.PathApi, config.PathP), routes.CreateProduct)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathP), routes.GetProducts)
	app.Get(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), routes.GetProduct)
	app.Put(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), routes.UpdateProduct)
	app.Delete(config.ConcatenateStrings(config.PathApi, config.PathP, config.PathID), routes.DeleteProduct)
	// EP for orders
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
