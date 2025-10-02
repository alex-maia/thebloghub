package models

import (
	"time"
)

type Theme struct {
	ID        uint   `gorm:"primaryKey"`
	Theme     string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
