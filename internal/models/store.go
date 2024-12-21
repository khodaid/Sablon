package models

import "github.com/khodaid/Sablon/internal/models/base"

type Store struct {
	base.Id
	Name         string   `gorm:"size:100; not null"`
	Address      string   `gorm:"size:255; not null"`
	Phone        string   `gorm:"size:20; not null; unique; index"`
	Email        string   `gorm:"size:100; not null; unique; index"`
	LogoFileName string   `gorm:"size:255;"`
	SupplierId   string   `gorm:"type:char(36); not null"`
	Supplier     Supplier `gorm:"foreignKey:SupplierId; references:ID; constarint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	base.TimeStamp
}
