package routes

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"
	"errors"

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

// get user by ID
func findUser(id int, user *models.User) error {
	database.DB.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user is not exist")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("There is no any users with this id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responceuser := CreateResponceUser(user)
	return c.Status(200).JSON(responceuser)
}

// update Users
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("There is no any users with this id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.Email = updateData.Email
	database.DB.Db.Save(&user)

	responceUser := CreateResponceUser(user)
	return c.Status(200).JSON(responceUser)
}

// delete user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("There is no any users with this id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Delete is succefully")
}
