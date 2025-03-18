package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type UserStoreRepository interface {
	WithTrx(trxHandle *gorm.DB) *userStoreRepository
	Save(userStore models.UserStore) (models.UserStore, error)
}

type userStoreRepository struct {
	db *gorm.DB
}

func NewUserStoreRepository(db *gorm.DB) *userStoreRepository {
	return &userStoreRepository{db: db}
}

func (r *userStoreRepository) WithTrx(trxHandle *gorm.DB) *userStoreRepository {
	r.db = trxHandle
	return r
}

func (r *userStoreRepository) Save(userStore models.UserStore) (models.UserStore, error) {
	if err := r.db.Create(&userStore).Error; err != nil {
		return userStore, err
	}

	return userStore, nil
}
