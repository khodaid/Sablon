package models

import "github.com/khodaid/Sablon/internal/models/base"

type UserStore struct {
	base.Id
	UserId  string `gorm:"type:char(36), not null"`
	User    *User  `gorm:"foreignKey:UserId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StoreId string `gorm:"type:char(36), not null"`
	Store   *Store `gorm:"foreignKey:StoreId; references:ID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	base.TimeStamp
}
