package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Save(models.Store) (models.Store, error)
	FindAllWithOutSoftDelete() ([]models.Store, error)
	FindById(string) (models.Store, error)
	Update(models.Store) (models.Store, error)
	SoftDelete(models.Store) (models.Store, error)
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *storeRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) Save(store models.Store) (models.Store, error) {
	err := r.db.Create(&store).Error

	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *storeRepository) FindAllWithOutSoftDelete() ([]models.Store, error) {
	var stores []models.Store
	err := r.db.Find(&stores).Error

	if err != nil {
		return nil, err
	}

	return stores, nil
}

func (r *storeRepository) FindById(id string) (models.Store, error) {
	var store models.Store

	err := r.db.Where("id = ?", id).First(&store).Error
	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *storeRepository) Update(store models.Store) (models.Store, error) {
	err := r.db.Save(&store).Error
	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *storeRepository) SoftDelete(store models.Store) (models.Store, error) {
	err := r.db.Delete(&store).Error
	if err != nil {
		return store, err
	}

	return store, nil
}

func (r *storeRepository) FindAllBySupplierId(supplier_id string) ([]models.Store, error) {
	var stores []models.Store

	err := r.db.Unscoped().Where("supplier_id = ?", supplier_id).Find(&stores).Error

	if err != nil {
		return stores, err
	}

	return stores, nil
}
