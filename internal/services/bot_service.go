package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/models"
)

type BotService struct {
	BotToken string
}

// PollUpdates fetches updates from Telegram using long polling
func (s *BotService) PollUpdates() {
	offset := 0
	for {
		updates, err := s.GetUpdates(offset)
		if err != nil {
			log.Printf("Error getting updates: %v\n", err)
			time.Sleep(1 * time.Second) // Retry after a short delay
			continue
		}

		for _, update := range updates {
			log.Printf("Received update: %+v\n", update)

			if update.Message.Text == "/start" {
				err := s.SendMessage(update.Message.Chat.ID, "Привет! Это Telegram бот!")
				if err != nil {
					log.Printf("Error sending message: %v\n", err)
				}
			}

			offset = update.UpdateID + 1 // Move the offset to the next update
		}

		time.Sleep(2 * time.Second) // Poll every 2 seconds
	}
}

func (s *BotService) GetUpdates(offset int) ([]models.TelegramUpdate, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", s.BotToken, offset)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error, status code: %d", resp.StatusCode)
	}

	var response struct {
		OK          bool                    `json:"ok"`
		Result      []models.TelegramUpdate `json:"result"`
		Description string                  `json:"description"` // Add if API returns errors in description
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("decoding error: %w", err)
	}

	if !response.OK {
		return nil, fmt.Errorf("API response error: %s", response.Description)
	}

	return response.Result, nil
}

// SendMessage sends a message to the Telegram chat
func (s *BotService) SendMessage(chatID int, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.BotToken)
	data := map[string]string{
		"chat_id": strconv.Itoa(chatID),
		"text":    text,
	}

	// Marshal the data into JSON
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Create the request body
	requestBody := bytes.NewBuffer(payload)

	// Send the POST request with the payload
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}
