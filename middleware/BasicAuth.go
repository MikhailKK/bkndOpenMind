package middleware

import (
	"bkndOpenMind/config"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
)

var username = config.GetUserName()
var password = "password"

func BasicAuthMiddleware(c *fiber.Ctx) error {
	// Получить значение заголовка "Authorization"
	authHeader := c.Get("Authorization")

	// Проверить, что значение заголовка начинается с "Basic "
	if len(authHeader) > 6 && authHeader[:6] == "Basic " {
		// Извлечь зашифрованные учетные данные (часть после "Basic ")
		encodedCredentials := authHeader[6:]

		// Расшифровать учетные данные из base64
		credentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid Authorization header",
			})
		}

		// Разделить учетные данные на имя пользователя и пароль
		userPass := string(credentials)
		if userPass != username+":"+password {
			// Если учетные данные не совпадают, вернуть ошибку "Unauthorized"
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Если учетные данные верны, продолжаем выполнение следующего Middleware или обработчика
		return c.Next()
	}

	// Если заголовок Authorization отсутствует или имеет некорректный формат, вернуть ошибку "Unauthorized"
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func ProtectedEndpoint(c *fiber.Ctx) error {
	return c.SendString("/api/v1")
}
