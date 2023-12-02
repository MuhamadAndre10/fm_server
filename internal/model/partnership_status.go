package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type PartnershipStatus struct {
	ID        string    `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	TosStatus bool      `gorm:"default:false"`
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	MitraID   string
}

func (u *PartnershipStatus) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
