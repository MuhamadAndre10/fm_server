package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ReceivedBalance struct {
	ID           string    `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	Balance      int64     `gorm:"size:50;not null"`
	Amount       int64     `gorm:"size:50;not null"`
	DateReceived time.Time `gorm:"size:50;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	MitraID      string
}

func (u *ReceivedBalance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
