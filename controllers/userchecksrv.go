package controllers

import (
	"bkndOpenMind/database"
	"bkndOpenMind/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

// // func GetAllUsers() ([]models.User, error) {
// // 	var users []models.User
// // 	result := database.DB.Find(&users)
// // 	return users, result.Error
// // }

// func MonitorHandler(c *fiber.Ctx) error {
// 	log.Println("Handling monitor request...")

// 	var users []models.User

// 	database.DB.Db.Where("id > ?", getLastRecordID()).Find(&users)

// 	for _, user := range users {
// 		SendTelegramMessage(user)
// 		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>ВОТ ТАКОЙ ЮЗЕР", &user)
// 	}

//		return c.SendStatus(fiber.StatusOK)
//	}
var prevUserCount int

func init() {
	// Запускаем горутину, которая будет каждые 5 секунд делать выборку и сравнивать количество пользователей
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				currentUserCount := GetUserCount()
				if currentUserCount > prevUserCount {
					s := currentUserCount - prevUserCount
					SendTelegramMessage(fmt.Sprint("Новых пользователей", s))
				}
				prevUserCount = currentUserCount
			}
		}
	}()
}

// GetUserCount селектит количество записей в таблице Users
func GetUserCount() int {
	var count int64
	database.DB.Db.Model(&models.User{}).Count(&count)
	return int(count)
}

const (
	telegramToken = "6529568374:AAH62hPem7MKcmLBtit5TYb1q7EGXOYwZHY"
	telegramID    = "315156239"
)

// отправка в телеграмм
func SendTelegramMessage(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		telegramToken, telegramID, message)

	// Отправка HTTP GET запроса на API Telegram
	go func() {
		resp, err := http.Get(apiURL)
		if err != nil {
			log.Printf("Failed to send Telegram message: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to send Telegram message, status code: %d", resp.StatusCode)
			return
		}
	}()

	return nil
}
