package services

import (
	"errors"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/models"
	"github.com/bormisov1/golang-sql-telegram-mini-app/internal/repositories"
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
