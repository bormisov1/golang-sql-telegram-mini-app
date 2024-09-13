// internal/repositories/user_repository.go
package repositories

import "example.com/internal/models"

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
