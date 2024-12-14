package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
	"github.com/khodaid/Sablon/pkg/helpers"
)

func SeedUsers() []models.User {
	users := []models.User{
		{
			Id:       base.Id{ID: "75c4a8f3-208c-4c6d-9e83-90d42443b233"},
			Name:     "khoda",
			Email:    "khodaid@gmail.com",
			Phone:    "08123123131",
			Password: string(helpers.GeneratePassword("password")),
		},
		{
			Id:       base.Id{ID: "ccc5a49d-fa89-4dec-8035-a69b1dbed62f"},
			Name:     "khoda supplier",
			Email:    "khodasupplier@gmail.com",
			Phone:    "08123123121",
			Password: string(helpers.GeneratePassword("password")),
		},
		{
			Id:       base.Id{ID: "f3c5792f-d368-4142-9533-674e98b685db"},
			Name:     "khoda store",
			Email:    "khodastore@gmail.com",
			Phone:    "08123123135",
			Password: string(helpers.GeneratePassword("password")),
		},
	}
	return users
}
