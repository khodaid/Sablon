package models

import "github.com/khodaid/Sablon/internal/models/base"

type User struct {
	// base.Id
	Name          string `gorm:"size:100;not null"`
	Email         string `gorm:"size:100;not null;unique;index"`
	Phone         string `gorm:"size:20;not null;unique;index"`
	Password      string `gorm:"size:100;not null;"`
	RememberToken string `gorm:"size:255;not null"`
	base.TimeStamp
}
