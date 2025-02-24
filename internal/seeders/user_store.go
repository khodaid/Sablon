package seeders

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/internal/models/base"
)

func SeedUserStore() []models.UserStore {
	userStore := []models.UserStore{
		{
			Id:      base.Id{"b55541ea-f698-4c05-87ec-0542304361ea"},
			UserId:  "f3c5792f-d368-4142-9533-674e98b685db",
			StoreId: "8d0897ea-1e75-47d8-b3ef-a001bdf9b8a1",
		},
	}
	return userStore
}
