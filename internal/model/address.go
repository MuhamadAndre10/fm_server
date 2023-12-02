package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID            string `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	City          string `gorm:"size:50"`
	StreetAddress string `gorm:"size:255"`
	ZipCode       string `gorm:"size:10"`
	State         string `gorm:"size:50;default:Indonesian"`
	CreatedAt     string `gorm:"autoCreateTime"`
	UpdatedAt     string `gorm:"autoUpdateTime"`
	MitraID       string
}

func (u *Address) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
