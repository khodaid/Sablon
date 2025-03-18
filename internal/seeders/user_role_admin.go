package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/pkg/helpers"
)

func SeedUserRoleAdmin() []models.UserRoleAdmin {
	userRoles := []models.UserRoleAdmin{
		{
			UserId:       "75c4a8f3-208c-4c6d-9e83-90d42443b233",
			SupplierId:   nil,
			RoleId:       "64fcda52-c67b-4502-945c-5089e2ecccd5",
			IsBackoffice: true,
		},
		{
			UserId:       "ccc5a49d-fa89-4dec-8035-a69b1dbed62f",
			SupplierId:   helpers.PointerString("b35654a0-2e56-4e87-bcd6-c687986bce08"),
			RoleId:       "f95667d1-d9a2-4501-9bdf-713472d1885d",
			IsBackoffice: true,
		},
		{
			UserId:     "f3c5792f-d368-4142-9533-674e98b685db",
			SupplierId: nil,
			RoleId:     "d8d23b40-0578-4c02-8f78-5b6c2763d30e",
		},
	}
	return userRoles
}
