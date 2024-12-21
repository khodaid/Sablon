package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Save(models.Store) (models.Store, error)
}

type repositories struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *repositories {
	return &repositories{db: db}
}

func (r *repositories) Save(store models.Store) (models.Store, error) {
	err := r.db.Create(&store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}
