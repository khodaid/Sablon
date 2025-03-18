package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type UserRoleAdminRepository interface {
	WithTrx(trxHandle *gorm.DB) *userRoleAdminRepository
	Save(userRoleAdmin models.UserRoleAdmin) (models.UserRoleAdmin, error)
}

type userRoleAdminRepository struct {
	db *gorm.DB
}

func NewUserRoleAdminRepository(db *gorm.DB) *userRoleAdminRepository {
	return &userRoleAdminRepository{db: db}
}

func (r *userRoleAdminRepository) WithTrx(trxHandle *gorm.DB) *userRoleAdminRepository {
	r.db = trxHandle
	return r
}

func (r *userRoleAdminRepository) Save(userRoleAdmin models.UserRoleAdmin) (models.UserRoleAdmin, error) {
	if err := r.db.Create(&userRoleAdmin).Error; err != nil {
		return userRoleAdmin, err
	}

	return userRoleAdmin, nil
}
