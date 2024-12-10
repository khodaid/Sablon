package service

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(validation.Login) (models.User, error)
}

type service struct {
	repository repositories.Repository
}

func NewUserService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) Login(input validation.Login) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	// kurang pengujian untuk user apakah ada atau tidak

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
