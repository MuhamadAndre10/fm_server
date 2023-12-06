package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                 string         `gorm:"PrimaryKey;size:50;not null;unique;<-:create"`
	NumberPhone        string         `gorm:"size:20;unique"`
	Email              string         `gorm:"size:50;unique"`
	Password           string         `gorm:"size:255"`
	VerificationStatus bool           `gorm:"default:false"`
	CreatedAt          time.Time      `gorm:"autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime"`
	MitraIdentity      *MitraIdentity `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE;references:ID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
