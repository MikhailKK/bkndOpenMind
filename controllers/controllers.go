package controllers

import (
	"bkndOpenMind/database"
	"my-gorm-fiber-app/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}
