package routes

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"

	"github.com/gofiber/fiber/v2"
)

type RequestSerializer struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

func CreateResponceRequest(requestModel models.Request) RequestSerializer {
	return RequestSerializer{ID: requestModel.ID, Description: requestModel.Description}
}

// create new request
func CreateRequest(c *fiber.Ctx) error {
	var request models.Request

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Create(&request)
	responceRequest := CreateResponceRequest(request)

	return c.Status(200).JSON(responceRequest)
}
