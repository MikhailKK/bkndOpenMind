package handlers

import (
	"my-gorm-fiber-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := controllers.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving users",
		})
	}
	return c.JSON(users)
}
