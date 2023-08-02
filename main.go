package main

import (
	"bkndOpenMind/auth"
	"bkndOpenMind/database"
	"bkndOpenMind/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database connection
	database.ConnectDB()
	app := fiber.New()
	// Применение BasicAuthMiddleware для всего приложения
	app.Use(auth.BasicAuthMiddleware)

	// Применение BasicAuthMiddleware только к конкретному эндпоинту
	// app.Post(config.PathApi+config.PathU, auth.BasicAuthMiddleware, auth.ProtectedEndpoint)

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	// app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))
}
