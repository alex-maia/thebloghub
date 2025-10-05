package services

import (
	"thebloghub/models"

	"gorm.io/gorm"
)

func GetAllImages(db *gorm.DB) ([]models.Image, error) {
	var images []models.Image
	if err := db.Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
