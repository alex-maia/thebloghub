package services

import (
	"thebloghub/models"

	"gorm.io/gorm"
)

func GetAllArticles(db *gorm.DB) ([]models.Article, error) {
	var articles []models.Article
	if err := db.Preload("Image").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
