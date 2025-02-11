package dto

import "github.com/khodaid/Sablon/internal/models"

type SupplierDetailFormatter struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	UrlLogo      string `json:"image_url"`
	ReferralCode string `json:"referral_code"`
}

func FormatDetailSupplier(supplier models.Supplier) SupplierDetailFormatter {
	supplierDetail := SupplierDetailFormatter{
		Id:           supplier.ID,
		Name:         supplier.Name,
		Address:      supplier.Address,
		Phone:        supplier.Phone,
		ReferralCode: supplier.ReferralCode,
	}

	return supplierDetail
}
