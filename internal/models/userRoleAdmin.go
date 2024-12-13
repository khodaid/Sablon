package models

import "github.com/khodaid/Sablon/internal/models/base"

type UserRoleAdmin struct {
	base.Id
	UserId string `gorm:"char:(36)"`
	User   User   `gorm:"foreignKey:UserId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
