package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
	"github.com/khodaid/Sablon/pkg/utils"
)

func SeedUsers() []models.User {
	users := []models.User{
		{
			Id:       base.Id{ID: "75c4a8f3-208c-4c6d-9e83-90d42443b233"},
			Name:     "khoda",
			Email:    "khodaid@gmail.com",
			Phone:    "08123123131",
			Password: string(utils.GeneratePassword("password")),
		},
	}
	return users
}
