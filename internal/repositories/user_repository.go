package repositories

import (
	"database/sql"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/models"
)

type UserRepository interface {
	Save(user models.User) error
}

// Реализация UserRepository для базы данных (например, PostgreSQL)
type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) Save(user models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, user.Name, user.Email)
	return err
}
