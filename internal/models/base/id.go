package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Id struct {
	ID string `gorm:"char:(36);primaryKey"`
}

func (base *Id) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == "" {
		base.ID = uuid.New().String()
	}
	return nil
}
