package main

import (
	"bkndOpenMind/config"
	"bkndOpenMind/controllers"
	"bkndOpenMind/database"
	"bkndOpenMind/middleware"
	"bkndOpenMind/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	_ "bkndOpenMind/docs"
)

func main() {
	// Initialize database connection
	database.ConnectDB()

	app := fiber.New()
	// Применение BasicAuthMiddleware для всего приложения
	// app.Use(middleware.BasicAuthMiddleware)

	// Используем CorsMiddleware из файла middlewares.go
	app.Use(middleware.CorsMiddleware)

	// Применение BasicAuthMiddleware только к конкретному эндпоинту
	app.Delete(config.PathApi+config.PathU+config.PathID, middleware.BasicAuthMiddleware, middleware.ProtectedEndpoint)
	// app.Get(config.PathApi+config.PathU, middleware.BasicAuthMiddleware, middleware.ProtectedEndpoint)

	// Defer closing the database connection when the application exits
	// defer database.DB.Close()

	// Setup routes
	routes.SetupRoutes(app)

	// initial telegram bot
	app.Get("/sendTelegramMessage", sendTelegramMessage)

	// Start the server
	// app.Listen(":3000")
	log.Fatal(app.Listen(":3000"))

}

// Send message to Telegramm
func sendTelegramMessage(c *fiber.Ctx) error {
	message := c.Query("message")
	if message == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Message not provided",
		})
	}

	err := controllers.SendTelegramMessage(message)
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send Telegram message",
		})
	}

	// Возврат успешного ответа
	return c.SendStatus(http.StatusOK)
}
