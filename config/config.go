package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)

	return sql.Open("postgres", connStr)
}
