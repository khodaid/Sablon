package models

import "github.com/khodaid/Sablon/internal/models/base"

type UserRole string

const (
	BackofficeUser UserRole = "backoffice"
	SupplieUser    UserRole = "supplier"
	StoreUser      UserRole = "store"
)

// daftar enum untuk coloum for_login untuk table role
var UserRoles = []string{"backoffice", "supplier", "store"}

type Role struct {
	base.Id
	Name     string   `gorm:"size:20;not null"`
	ForLogin UserRole `gorm:"type:user_role_enum;not null"`
	Value    string   `gorm:"size:20;not null"`
	base.TimeStamp
}
