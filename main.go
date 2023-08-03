package main

import (
	"bkndOpenMind/auth"
	"bkndOpenMind/config"
	"bkndOpenMind/controllers"
	"bkndOpenMind/database"
	"bkndOpenMind/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database connection
	database.ConnectDB()
	app := fiber.New()
	// Применение BasicAuthMiddleware для всего приложения
	// app.Use(auth.BasicAuthMiddleware)

	// Применение BasicAuthMiddleware только к конкретному эндпоинту
	app.Delete(config.PathApi+config.PathU+config.PathID, auth.BasicAuthMiddleware, auth.ProtectedEndpoint)
	app.Delete(config.PathApi+config.PathP+config.PathID, auth.BasicAuthMiddleware, auth.ProtectedEndpoint)
	app.Get(config.PathApi+config.PathU, auth.BasicAuthMiddleware, auth.ProtectedEndpoint)
	app.Get(config.PathApi+config.PathO, auth.BasicAuthMiddleware, auth.ProtectedEndpoint)

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
