package handlers

import (
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/services"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketHandler struct {
	GameService *services.GameService
}

func NewWebSocketHandler(gameService *services.GameService) *WebSocketHandler {
	return &WebSocketHandler{GameService: gameService}
}

func (h *WebSocketHandler) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	playerID := r.URL.Query().Get("playerID")
	h.GameService.RegisterPlayer(playerID, conn)
	h.GameService.StartGameLoop(playerID, conn)
}
