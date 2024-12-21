package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(string) (models.User, error)
	// FindAll() []models.User
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ? ", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Model(&users).Error

	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
