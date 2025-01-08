package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type SupplierRepositoryForStore interface {
	GetIdByCode(string) (string, error)
}

type supplierRepositories struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *supplierRepositories {
	return &supplierRepositories{db}
}

func (r *supplierRepositories) GetIdByCode(code string) (string, error) {
	var supplier models.Supplier
	err := r.db.Where("referral_code = ?", code).Find(&supplier).Error

	if err != nil {
		return supplier.ID, err
	}

	return supplier.ID, nil
}
