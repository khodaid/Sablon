package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
)

func SeedStore() []models.Store {
	stores := []models.Store{
		{
			Id:           base.Id{ID: "8d0897ea-1e75-47d8-b3ef-a001bdf9b8a1"},
			Name:         "khoda store",
			Address:      "kalijaran",
			Phone:        "90090980980",
			Email:        "khodastore@gmail.com",
			LogoFileName: "download.jpeg",
			SupplierId:   "b35654a0-2e56-4e87-bcd6-c687986bce08",
		},
	}
	return stores
}
