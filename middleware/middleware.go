package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CorsMiddleware(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	return c.Next()
}
