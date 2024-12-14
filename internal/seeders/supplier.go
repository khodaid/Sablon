package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
	"github.com/khodaid/Sablon/pkg/helpers"
)

func SeedSuplier() []models.Supplier {
	suppliers := []models.Supplier{
		{
			Id:           base.Id{ID: "b35654a0-2e56-4e87-bcd6-c687986bce08"},
			Name:         "khoda supplier",
			Phone:        "934593539",
			ReferralCode: string(helpers.GenerateSupplierCode()),
		},
	}
	return suppliers
}
