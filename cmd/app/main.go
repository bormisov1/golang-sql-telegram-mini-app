package main

import (
	"github.com/bormisov1/golang-sql-telegram-mini-app/config"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/handlers"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/repositories"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	userRepo := &repositories.PostgresUserRepository{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userHandler := &handlers.UserHandler{Service: userService}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable not set")
	}
	print(botToken)

	//gameService := services.NewGameService()
	//wsHandler := handlers.NewWebSocketHandler(gameService)

	botService := &services.BotService{BotToken: botToken}
	telegramHandler := &handlers.TelegramHandler{Service: *botService}
	htmlHandler := &handlers.HTMLHandler{}

	//http.HandleFunc("/ws", wsHandler.HandleWS)
	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/telegram/webhook", telegramHandler.WebhookHandler) // Webhook для Telegram бота
	http.HandleFunc("/telegram/miniapp", htmlHandler.ServeMiniApp)       // HTML страница мини-приложения

	go func() {
		log.Println("Starting long polling...")
		botService.PollUpdates()
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
