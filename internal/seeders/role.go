package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
)

func SeedRole() []models.Role {
	roles := []models.Role{
		{
			Id:       base.Id{ID: "64fcda52-c67b-4502-945c-5089e2ecccd5"},
			Name:     "User Root",
			ForLogin: models.BackofficeUser,
			Value:    "user root",
		},
		{
			Id:       base.Id{ID: "f95667d1-d9a2-4501-9bdf-713472d1885d"},
			Name:     "User Root",
			ForLogin: models.SupplieUser,
			Value:    "user root",
		},
		{
			Id:       base.Id{ID: "d8d23b40-0578-4c02-8f78-5b6c2763d30e"},
			Name:     "User Root",
			ForLogin: models.StoreUser,
			Value:    "user root",
		},
	}
	return roles
}
