package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(string) (models.User, error)
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
	FindAllUserByStore(string) ([]models.User, error)
	FindById(string) (models.User, error)
	Update(user models.User) (models.User, error)
	SoftDelete(user models.User) (models.User, error)
	FindAllWithSoftDelete() ([]models.User, error)
	FindSoftDeleteById(id string) (models.User, error)
	HardDelete(user models.User) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ? ", email).Preload("UserRoleAdmin.Role").Preload("UserStore.Store").Find(&user).Error

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

func (r *repository) FindById(id string) (models.User, error) {
	var user models.User

	err := r.db.Where("id", id).Preload("UserRoleAdmin.Role").First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SoftDelete(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) HardDelete(user models.User) (models.User, error) {
	err := r.db.Unscoped().Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAllWithSoftDelete() ([]models.User, error) {
	var users []models.User

	err := r.db.Unscoped().Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil

}

func (r *repository) FindSoftDeleteById(id string) (models.User, error) {
	var user models.User

	err := r.db.Unscoped().First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAllUserByStore(id string) ([]models.User, error) {
	var users []models.User
	var user models.User

	err := r.db.Preload("UserStore").Where("id = ?", id).First(&user).Error

	if err != nil {
		return users, err
	}

	err = r.db.Preload("UserRoleAdmin.Role").
		Joins("JOIN user_stores ON users.id = user_stores.user_id").
		Where("user_stores.store_id = ?", user.UserStore.StoreId).
		Group("users.id").
		Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
