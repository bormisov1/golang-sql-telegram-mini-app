package handlers

import (
	"encoding/json"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/models"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/services"
	"net/http"
)

type TelegramHandler struct {
	Service services.BotService
}

func (h *TelegramHandler) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var update models.TelegramUpdate
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//err = h.Service.HandleUpdate(update)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
}
