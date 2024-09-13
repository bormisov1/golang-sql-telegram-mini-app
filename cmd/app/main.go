// cmd/app/main.go
package main

import (
    "log"
    "net/http"
    "example.com/config"
    "example.com/internal/handlers"
    "example.com/internal/repositories"
    "example.com/internal/services"
)

func main() {
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Cannot connect to DB:", err)
    }

    userRepo := &repositories.PostgresUserRepository{DB: db}
    userService := &services.UserService{Repo: userRepo}
    userHandler := &handlers.UserHandler{Service: userService}

    // Telegram Bot
    botToken := "YOUR_TELEGRAM_BOT_TOKEN"
    telegramService := &services.BotService{BotToken: botToken}
    telegramHandler := &handlers.TelegramHandler{Service: telegramService}

    // HTML handler for mini app
    htmlHandler := &handlers.HTMLHandler{}

    http.HandleFunc("/users", userHandler.CreateUser)
    http.HandleFunc("/telegram/webhook", telegramHandler.WebhookHandler)  // Webhook для Telegram бота
    http.HandleFunc("/telegram/miniapp", htmlHandler.ServeMiniApp)        // HTML страница мини-приложения

    log.Fatal(http.ListenAndServe(":8080", nil))
}
