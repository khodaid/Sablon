package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]models.Role, error)
	FindById(id string) (models.Role, error)
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindAll() ([]models.Role, error) {
	var roles []models.Role

	err := r.db.Find(&roles).Error

	if err != nil {
		return roles, err
	}

	return roles, nil
}

func (r *roleRepository) FindById(id string) (models.Role, error) {
	var role models.Role
	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}
