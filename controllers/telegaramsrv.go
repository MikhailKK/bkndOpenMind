package controllers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// const (
// 	telegramToken = "6529568374:AAH62hPem7MKcmLBtit5TYb1q7EGXOYwZHY"
// 	telegramID    = "315156239"
// )

// func SendTelegramMessage(message string) error {
// 	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
// 		telegramToken, telegramID, message)

// 	// Отправка HTTP GET запроса на API Telegram
// 	go func() {
// 		resp, err := http.Get(apiURL)
// 		if err != nil {
// 			log.Printf("Failed to send Telegram message: %v", err)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		if resp.StatusCode != http.StatusOK {
// 			log.Printf("Failed to send Telegram message, status code: %d", resp.StatusCode)
// 			return
// 		}
// 	}()

// 	return nil
// }
