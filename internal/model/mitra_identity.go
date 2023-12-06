package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MitraIdentity struct {
	MitraID           string              `gorm:"<-create,primaryKey;size:50;not null;unique"`
	FirstName         string              `gorm:"size:50"`
	LastName          string              `gorm:"size:50"`
	Age               int                 `gorm:"size:3"`
	DateBirth         time.Time           `gorm:"-"`
	CreatedAt         time.Time           `gorm:"autoCreateTime"`
	UpdatedAt         time.Time           `gorm:"autoUpdateTime"`
	FastPay           []FastPay           `gorm:"foreignKey:MitraID;references:MitraID"`
	ReceivedBalance   []ReceivedBalance   `gorm:"foreignKey:MitraID;references:MitraID"`
	UnReceivedBalance []UnReceivedBalance `gorm:"foreignKey:MitraID;references:MitraID"`
	Address           Address             `gorm:"foreignKey:MitraID;references:MitraID"`
	Tier              Tier                `gorm:"foreignKey:MitraID;references:MitraID"`
	ReferralCode      ReferralCode        `gorm:"foreignKey:MitraID;references:MitraID"`
	PartnershipStatus PartnershipStatus   `gorm:"foreignKey:MitraID;references:MitraID"`
	UserID            string
}

func (u *MitraIdentity) BeforeCreate(tx *gorm.DB) (err error) {
	u.MitraID = uuid.New().String()
	return
}
