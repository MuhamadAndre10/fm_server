package model

import "time"

type Tier struct {
	ID           string    `gorm:"primaryKey;size:50;not null;unique;<-:create"`
	Name         string    `gorm:"size:50;not null"`
	BonusPercent int       `gorm:"size:50;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
