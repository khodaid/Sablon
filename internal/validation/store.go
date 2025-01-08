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
