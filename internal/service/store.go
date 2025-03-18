package service

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/validation"
)

type StoreService interface {
	StoreRegister(validation.RegisterStoreInput, string) (models.Store, error)
	GetStore(string) (models.Store, error)
	GetAllStoreWithOutSoftDelete() ([]models.Store, error)
	UpdateStore(string, validation.UpdateStoreInput) (models.Store, error)
	UpdateLogoStore(storeId string, newFileName string) (models.Store, error)
	SoftDeleteStore(string) (models.Store, error)
}

type storeService struct {
	storeRepository   repositories.StoreRepository
	supplierRepositoy repositories.SupplierRepositoryForStore
}

func NewStoreService(storeRepository repositories.StoreRepository, supplierRepository repositories.SupplierRepositoryForStore) *storeService {
	return &storeService{storeRepository: storeRepository, supplierRepositoy: supplierRepository}
}

func (s *storeService) StoreRegister(input validation.RegisterStoreInput, fileName string) (models.Store, error) {
	store := models.Store{}

	supplierId, err := s.supplierRepositoy.GetIdByCode(input.SupplierCode)
	if err != nil {
		return store, err
	}

	// file := input.Logo

	store.Name = input.Name
	store.Address = input.Address
	store.Email = input.Email
	store.Phone = input.Phone
	store.SupplierId = supplierId
	store.LogoFileName = fileName

	newStore, err := s.storeRepository.Save(store)

	if err != nil {
		return store, err
	}

	return newStore, nil
}

func (s *storeService) GetStore(store_id string) (models.Store, error) {
	store, err := s.storeRepository.FindById(store_id)

	if err != nil {
		return store, err
	}

	return store, nil
}

func (s *storeService) GetAllStoreWithOutSoftDelete() ([]models.Store, error) {
	stores, err := s.storeRepository.FindAllWithOutSoftDelete()
	if err != nil {
		return stores, err
	}

	return stores, nil
}

func (s *storeService) UpdateStore(store_id string, input validation.UpdateStoreInput) (models.Store, error) {
	store, err := s.storeRepository.FindById(store_id)
	if err != nil {
		return store, err
	}

	store.Name = input.Name
	store.Address = input.Address
	store.Email = input.Email
	store.Phone = input.Phone

	result, err := s.storeRepository.Update(store)

	if err != nil {
		return store, err
	}

	return result, nil
}

func (s *storeService) UpdateLogoStore(store_id string, filename string) (models.Store, error) {
	store, err := s.storeRepository.FindById(store_id)
	if err != nil {
		return store, err
	}

	store.LogoFileName = filename

	result, err := s.storeRepository.Update(store)

	if err != nil {
		return store, err
	}

	return result, nil
}

func (s *storeService) SoftDeleteStore(store_id string) (models.Store, error) {
	store, err := s.storeRepository.FindById(store_id)

	if err != nil {
		return store, err
	}

	result, err := s.storeRepository.SoftDelete(store)

	if err != nil {
		return store, err
	}

	return result, nil
}
