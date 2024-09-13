// config/config.go
package config

import (
    "database/sql"
    _ "github.com/lib/pq" // Используем драйвер PostgreSQL
)

func ConnectDB() (*sql.DB, error) {
    connStr := "user=youruser dbname=yourdb sslmode=disable"
    return sql.Open("postgres", connStr)
}
