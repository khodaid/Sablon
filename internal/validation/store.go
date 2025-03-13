package validation

import "mime/multipart"

type RegisterStoreInput struct {
	Name         string                `form:"name" binding:"required"`
	Address      string                `form:"address" binding:"required"`
	Phone        string                `form:"phone" binding:"required"`
	Email        string                `form:"email" binding:"required,email"`
	SupplierCode string                `form:"supplier_code" binding:"required"`
	Logo         *multipart.FileHeader `form:"logo" binding:"required"`
}

type UpdateStoreInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email" binding:"required"`
}

type UpdateLogoStoreInput struct {
	Logo *multipart.FileHeader `form:"logo" binding:"required"`
}
