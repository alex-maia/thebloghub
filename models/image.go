package models

import "time"

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"size:255;uniqueIndex;not null"` // caminho ou URL da imagem
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
