package dto

import "github.com/khodaid/Sablon/internal/models"

type loginDetailFormatter struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	IsBackoffice bool   `json:"is_backoffice"`
	Role         string `json:"role"`
}

func FormatDetailUserLogin(user models.User) loginDetailFormatter {
	data := loginDetailFormatter{
		Name:         user.Name,
		Email:        user.Email,
		IsBackoffice: user.UserRoleAdmin.IsBackoffice,
		Role:         user.UserRoleAdmin.Role.Name,
	}
	return data
}
