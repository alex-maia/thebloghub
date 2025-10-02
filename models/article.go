package models

import "time"

type Article struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255;not null"`
	Lead      string `gorm:"size:255;not null"`
	Text      string `gorm:"type:text;not null"`
	ImageID   *uint  // FK para Image (opcional)
	Image     *Image `gorm:"foreignKey:ImageID"`
	ThemeID   *uint  // FK para Theme (opcional)
	Theme     *Theme `gorm:"foreignKey:ThemeID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
