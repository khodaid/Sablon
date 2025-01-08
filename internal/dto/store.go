package dto

import (
	"github.com/khodaid/Sablon/internal/models"
	"github.com/khodaid/Sablon/pkg/helpers"
)

type StoreFormatter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

func FormatStoreRegister(host string, store models.Store) StoreFormatter {
	storeFormatter := StoreFormatter{}

	storeFormatter.Id = store.ID
	storeFormatter.Name = store.Name
	storeFormatter.Address = store.Address
	storeFormatter.Email = store.Email
	storeFormatter.ImageUrl = helpers.UrlLogo(host, store.LogoFileName)

	return storeFormatter
}
