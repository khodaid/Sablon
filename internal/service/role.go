package service

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/repositories"
)

type RoleService interface {
	GetAllRole() ([]models.Role, error)
	GetRole(id string) (models.Role, error)
}

type roleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(roleRepository repositories.RoleRepository) *roleService {
	return &roleService{roleRepository}
}

func (s *roleService) GetAllRole() ([]models.Role, error) {
	roles, err := s.roleRepository.FindAll()

	if err != nil {
		return roles, err
	}

	return roles, nil
}
