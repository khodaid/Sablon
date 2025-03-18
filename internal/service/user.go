package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
	"github.com/khodaid/Sablon/internal/validation"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Login(validation.LoginUserInput) (models.User, error)
	Register(validation.RegisterUserStoreAdminInput) (models.User, error)
	CreateEmployeeStoreByUserRootStore(input validation.CreateUserEmployeeStore, userIdAuth string) (models.User, error)
	GetUserById(id string) (models.User, error)
	UpdateUserById(userId string, userInput validation.UpdateUserStore) (models.User, error)
	SoftDeleteUserById(id string) (models.User, error)
	GetAllWithOutSoftDelete() ([]models.User, error)
	GetAllUserByStore(string) ([]models.User, error)
	GetAllUsersWithSoftDelete() ([]models.User, error)
	GetUserWithSoftDelete(string) (models.User, error)
	HardDeleteUser(string) (models.User, error)
}

type userService struct {
	db                     *gorm.DB
	repository             repositories.UserRepository
	userStoreRepository    repositories.UserStoreRepository
	userRoleAdminRepositor repositories.UserRoleAdminRepository
	roleRepository         repositories.RoleRepository
}

func NewUserService(db *gorm.DB, repository repositories.UserRepository, roleRepository repositories.RoleRepository, userStoreRepository repositories.UserStoreRepository, userRoleAdminRepository repositories.UserRoleAdminRepository) *userService {
	return &userService{db: db, repository: repository, roleRepository: roleRepository, userStoreRepository: userStoreRepository, userRoleAdminRepositor: userRoleAdminRepository}
}

func (s *userService) Login(input validation.LoginUserInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == uuid.Nil.String() {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) Register(input validation.RegisterUserStoreAdminInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	// cek kesamaan password dengan confrimed
	if input.Password != input.ConfirmPassword {
		return user, errors.New("password no equals")
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password_hash)

	new_user, err := s.repository.Save(user)
	if err != nil {
		return new_user, err
	}

	return new_user, nil
}

func (s *userService) CreateEmployeeStoreByUserRootStore(input validation.CreateUserEmployeeStore, userIdAuth string) (models.User, error) {
	trx := s.db.Begin()
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	if input.Password != input.ConfirmPassword {
		return user, errors.New("password no equals")
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password_hash)

	new_user, err := s.repository.WithTrx(trx).Save(user)
	if err != nil {
		trx.Rollback()
		return new_user, err
	}

	userRoleAdmin := models.UserRoleAdmin{
		UserId: new_user.ID,
		RoleId: input.RoleId,
	}

	_, err = s.userRoleAdminRepositor.WithTrx(trx).Save(userRoleAdmin)
	if err != nil {
		trx.Rollback()
		return new_user, err
	}

	user, err = s.repository.FindById(userIdAuth)
	if err != nil {
		trx.Rollback()
		return new_user, err
	}

	userStore := models.UserStore{
		UserId:  new_user.ID,
		StoreId: user.UserStore.StoreId,
	}

	_, err = s.userStoreRepository.WithTrx(trx).Save(userStore)
	if err != nil {
		trx.Rollback()
		return new_user, err
	}

	if err := trx.Commit().Error; err != nil {
		return new_user, err
	}

	return new_user, nil
}

func (s *userService) GetUserById(id string) (models.User, error) {
	user, err := s.repository.FindById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) UpdateUserById(userId string, userInput validation.UpdateUserStore) (models.User, error) {
	userData, err := s.repository.FindById(userId)

	if err != nil {
		return userData, err
	}

	roleData, err := s.roleRepository.FindById(userInput.RoleId)

	if err != nil {
		return userData, err
	}

	userData.Name = userInput.Name
	userData.Email = userInput.Email
	userData.Phone = userInput.Phone
	userData.UserRoleAdmin.RoleId = roleData.ID

	result, err := s.repository.Update(userData)

	if err != nil {
		return userData, err
	}

	return result, err
}

func (s *userService) SoftDeleteUserById(id string) (models.User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	result, err := s.repository.SoftDelete(user)

	if err != nil {
		return user, err
	}

	return result, nil
}

func (s *userService) GetAllWithOutSoftDelete() ([]models.User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) GetAllUserByStore(id string) ([]models.User, error) {
	users, err := s.repository.FindAllUserByStore(id)

	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) GetAllUsersWithSoftDelete() ([]models.User, error) {
	users, err := s.repository.FindAllWithSoftDelete()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) GetUserWithSoftDelete(id string) (models.User, error) {
	user, err := s.repository.FindSoftDeleteById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) HardDeleteUser(id string) (models.User, error) {
	user, err := s.repository.FindSoftDeleteById(id)

	if err != nil {
		return user, err
	}

	result, err := s.repository.HardDelete(user)

	if err != nil {
		return user, err
	}

	return result, nil
}
