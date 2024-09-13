// internal/models/telegram.go
package models

type TelegramUpdate struct {
    UpdateID int           `json:"update_id"`
    Message  TelegramMessage `json:"message"`
}

type TelegramMessage struct {
    MessageID int    `json:"message_id"`
    Chat      TelegramChat `json:"chat"`
    Text      string `json:"text"`
}

type TelegramChat struct {
    ID int `json:"id"`
}
