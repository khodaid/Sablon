package repositories

import (
	"github.com/khodaid/Sablon/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(string) (models.User, error)
	Save(models.User) (models.User, error)
	FindAll() ([]models.User, error)
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

func (r *repositories) FindById(id string) (models.User, error) {
	var user models.User

	err := r.db.Where("id", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositories) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositories) SoftDelete(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositories) HardDelete(user models.User) (models.User, error) {
	err := r.db.Unscoped().Delete(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositories) FindAllWithSoftDelete() ([]models.User, error) {
	var users []models.User

	err := r.db.Unscoped().Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil

}

func (r *repositories) GetAllUserByStore(email string) ([]models.User, error) {
	var users []models.User

	err := r.db.Model(&models.User{}).
		Joins("JOIN user_stores us1 ON users.id = us1.user_id").
		Joins("JOIN user_stores us2 ON us1.store_id = us2.store_id").
		Joins("JOIN users u ON us2.user_id = u.id").
		Group("u.id").
		Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
