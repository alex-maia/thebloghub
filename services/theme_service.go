package services

import (
	"thebloghub/models"

	"gorm.io/gorm"
)

func GetAllThemes(db *gorm.DB) ([]models.Theme, error) {
	var themes []models.Theme
	if err := db.Find(&themes).Error; err != nil {
		return nil, err
	}
	return themes, nil
}
