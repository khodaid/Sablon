package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(validation.Login) (models.User, error)
	Register(validation.RegisterUserStoreAdminInput) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Login(input validation.Login) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == uuid.Nil.String() {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) Register(input validation.RegisterUserStoreAdminInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	// cek kesamaan password dengan confrimed
	if input.Password != input.ConfirmPassword {
		return user, errors.New("password no equals")
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password_hash)

	new_user, err := s.repository.Save(user)
	if err != nil {
		return new_user, err
	}

	return new_user, nil
}
