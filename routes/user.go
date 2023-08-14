package routes

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	// this is not the model of user it's serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
}

// simple EP
func FirstEP(c *fiber.Ctx) error {
	return c.SendString("The API")
}

func CreateResponceUser(userModel models.Users) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, Phone: userModel.Phone}
}

// create new user
func CreateUser(c *fiber.Ctx) error {
	var user models.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Create(&user)
	responceUser := CreateResponceUser(user)

	return c.Status(200).JSON(responceUser)
}

// get all users
func GetUsers(c *fiber.Ctx) error {
	users := []models.Users{}

	database.DB.Db.Find(&users)
	responceUsers := []UserSerializer{}
	for _, user := range users {
		responceUser := CreateResponceUser(user)
		responceUsers = append(responceUsers, responceUser)
	}

	return c.Status(200).JSON(responceUsers)
}

// get user by ID
func findUser(id int, user *models.Users) error {
	database.DB.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user is not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.Users
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

	var user models.Users
	if err != nil {
		return c.Status(400).JSON("There is no any users with this id")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		Phone     string `json:"phone"`
	}
	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.Phone = updateData.Phone
	database.DB.Db.Save(&user)

	responceUser := CreateResponceUser(user)
	return c.Status(200).JSON(responceUser)
}

// delete user
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.Users
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
