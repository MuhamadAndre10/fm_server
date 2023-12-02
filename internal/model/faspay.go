package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type FastPay struct {
	ID                string    `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	BankAccountName   string    `gorm:"size:50;not null"`
	BankAccountNumber int32     `gorm:"size:50;not null"`
	NumberDana        string    `gorm:"size:14;not null"`
	NumberShoppePay   string    `gorm:"size:14;not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
	MitraID           string
}

func (u *FastPay) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
