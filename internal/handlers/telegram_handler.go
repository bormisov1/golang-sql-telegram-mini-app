// internal/handlers/telegram_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "example.com/internal/models"
    "example.com/internal/services"
)

type TelegramHandler struct {
    Service services.TelegramService
}

func (h *TelegramHandler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
    var update models.TelegramUpdate
    err := json.NewDecoder(r.Body).Decode(&update)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = h.Service.HandleUpdate(update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
