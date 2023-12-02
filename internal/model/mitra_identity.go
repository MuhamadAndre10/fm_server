package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MitraIdentity struct {
	MitraID   string    `gorm:"<-create,primaryKey;size:50;not null;unique"`
	FirstName string    `gorm:"size:50;not null"`
	LastName  string    `gorm:"size:50;not null"`
	Age       int       `gorm:"size:3;not null"`
	DateBirth time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *MitraIdentity) BeforeCreate(tx *gorm.DB) (err error) {
	u.MitraID = uuid.New().String()
	return
}
