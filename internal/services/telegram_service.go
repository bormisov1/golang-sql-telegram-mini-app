// internal/services/telegram_service.go
package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "example.com/internal/models"
)

type TelegramService interface {
    SendMessage(chatID int, text string) error
    HandleUpdate(update models.TelegramUpdate) error
}

type BotService struct {
    BotToken string
}

func (s *BotService) SendMessage(chatID int, text string) error {
    url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.BotToken)
    body := map[string]interface{}{
        "chat_id": chatID,
        "text":    text,
    }

    jsonBody, _ := json.Marshal(body)
    _, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
    return err
}

func (s *BotService) HandleUpdate(update models.TelegramUpdate) error {
    // Логика обработки входящего апдейта, например, команда "/start"
    if update.Message.Text == "/start" {
        return s.SendMessage(update.Message.Chat.ID, "Привет! Это Telegram бот!")
    }
    return nil
}
