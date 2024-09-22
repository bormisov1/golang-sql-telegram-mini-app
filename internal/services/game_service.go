package services

import (
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/models"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	allowedDelta   = 0.1
	updateInterval = 200 * time.Millisecond
	screenWidth    = 400
)

type GameService struct {
	Players map[string]*models.Player
}

func NewGameService() *GameService {
	return &GameService{
		Players: make(map[string]*models.Player),
	}
}

func (g *GameService) RegisterPlayer(playerID string, conn *websocket.Conn) {
	g.Players[playerID] = &models.Player{ID: playerID, Position: 0, Direction: true, Speed: 1.0}
}

func (g *GameService) StartGameLoop(playerID string, conn *websocket.Conn) {
	player := g.Players[playerID]
	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				g.sendDirection(conn, player)
			}
		}
	}()

	for {
		var message map[string]interface{}
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading JSON:", err)
			return
		}

		switch message["type"] {
		case "position":
			position := message["position"].(float64)
			if !validatePosition(player.Position, position) {
				log.Println("Player speed violation")
				continue
			}
			player.Position = position
		case "direction":
			directionMultiplier := float64(boolToDirection(message["direction"].(bool)))
			timeDelta := float64(updateInterval / time.Millisecond)

			player.Position = player.Position + timeDelta*directionMultiplier*player.Speed
		}
	}
}

func (g *GameService) sendDirection(conn *websocket.Conn, player *models.Player) {
	err := conn.WriteJSON(map[string]interface{}{
		"type":      "direction",
		"direction": player.Direction,
	})
	if err != nil {
		log.Println("Error sending direction:", err)
	}
}

func validatePosition(currentPosition float64, updatedPosition float64) bool {
	if updatedPosition < 0 || updatedPosition >= 400 || abs(updatedPosition-currentPosition) > allowedDelta {
		log.Println("Player speed violation")
		return false
	}
	return true
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func boolToDirection(b bool) int {
	if b {
		return 1
	}
	return -1
}
