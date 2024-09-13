// internal/services/user_service.go
package services

import (
    "example.com/internal/models"
    "example.com/internal/repositories"
)

type UserService struct {
    Repo repositories.UserRepository
}

func (s *UserService) CreateUser(user models.User) error {
    // Валидация или другая бизнес-логика
    if user.Name == "" || user.Email == "" {
        return errors.New("name and email cannot be empty")
    }
    return s.Repo.Save(user)
}
