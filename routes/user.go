package routes

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"

	// "bkndOpenMind/handlers"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	// this is not the model of user it's serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// simple EP
func FirstEP(c *fiber.Ctx) error {
	return c.SendString("The API")
}

func CreateResponceUser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Email: userModel.Email}
}

// create new user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Create(&user)
	responceUser := CreateResponceUser(user)

	return c.Status(200).JSON(responceUser)
}

// get all users
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Db.Find(&users)
	responceUsers := []UserSerializer{}
	for _, user := range users {
		responceUser := CreateResponceUser(user)
		responceUsers = append(responceUsers, responceUser)
	}

	return c.Status(200).JSON(responceUsers)
}
