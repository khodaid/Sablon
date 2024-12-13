package models

import "github.com/khodaid/Sablon/internal/models/base"

type Supplier struct {
	base.Id
	Name         string `gorm:"size:100;not null"`
	Address      string `gorm:"size:255;not null"`
	Phone        string `gorm:"size:15;not null;unique;index"`
	LogoFileName string `gorm:"size:255;unique"`
	ReferralCode string `gorm:"size:8;not null;unique"`
	base.TimeStamp
}
