package service

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/validation"
)

type StoreService interface {
	StoreRegister(validation.RegisterStoreInput) (models.Store, error)
}

type storeService struct {
	storeRepository   repositories.StoreRepository
	supplierRepositoy repositories.SupplierRepositoryForStore
}

func NewStoreService(storeRepository repositories.StoreRepository, supplierRepository repositories.SupplierRepositoryForStore) *storeService {
	return &storeService{storeRepository: storeRepository, supplierRepositoy: supplierRepository}
}

func (s *storeService) StoreRegister(input validation.RegisterStoreInput) (models.Store, error) {
	store := models.Store{}

	supplierId, err := s.supplierRepositoy.GetIdByCode(input.SupplierCode)
	if err != nil {
		return store, err
	}

	file := input.Logo

	store.Name = input.Name
	store.Address = input.Address
	store.Email = input.Email
	store.Phone = input.Phone
	store.SupplierId = supplierId
	store.LogoFileName = file.Filename

	newStore, err := s.storeRepository.Save(store)

	if err != nil {
		return store, err
	}

	return newStore, nil
}
