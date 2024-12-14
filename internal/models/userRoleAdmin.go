package models

import "github.com/khodaid/Sablon/internal/models/base"

type UserRoleAdmin struct {
	base.Id
	UserId       string   `gorm:"char:(36); not null"`
	User         User     `gorm:"foreignKey:UserId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SupplierId   string   `gorm:"char:(36)"`
	Supplier     Supplier `gorm:"foreignKey:SupplierId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleId       string   `gorm:"char:(36); not null"`
	Role         Role     `gorm:"foreignKey:RoleId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsBackoffice bool     `gorm:"default:false; not null"`
	base.TimeStamp
}
