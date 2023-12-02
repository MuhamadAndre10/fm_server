package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ReferralCode struct {
	ID        string    `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	Code      string    `gorm:"size:10;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	MitraID   string
}

func (u *ReferralCode) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
