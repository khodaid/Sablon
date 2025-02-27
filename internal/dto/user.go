package dto

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type loginDetailFormatter struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	IsBackoffice bool   `json:"is_backoffice"`
	Role         string `json:"role"`
	LogoUrl      string `json:logo_url`
}

type allUserStoreFormatter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
}

func FormatDetailUserLogin(host string, user models.User) loginDetailFormatter {
	data := loginDetailFormatter{
		Name:         user.Name,
		Email:        user.Email,
		IsBackoffice: user.UserRoleAdmin.IsBackoffice,
		Role:         user.UserRoleAdmin.Role.Name,
		LogoUrl:      helpers.UrlLogo(host, user.UserStore.Store.LogoFileName),
	}
	return data
}

func FormatUserStore(user models.User) allUserStoreFormatter {
	return allUserStoreFormatter{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.UserRoleAdmin.Role.Name,
	}
}

func FormatUsersStore(users []models.User) []allUserStoreFormatter {
	usersData := make([]allUserStoreFormatter, 0, len(users))
	for _, v := range users {
		usersData = append(usersData, allUserStoreFormatter{
			Name:  v.Name,
			Email: v.Email,
			Phone: v.Phone,
			Role:  v.UserRoleAdmin.Role.Name,
		})
	}

	return usersData
}
