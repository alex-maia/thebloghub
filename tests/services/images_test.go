package services_test

import (
	"testing"
	"time"

	"thebloghub/config"
	"thebloghub/database"
	"thebloghub/models"
	"thebloghub/services"

	"gorm.io/gorm"
)

// setupTestDB prepara um banco limpo para cada teste
func setupTestDBImage(t *testing.T) *gorm.DB {
	// Carregar config de teste
	err := config.LoadTestConfig()
	if err != nil {
		t.Fatalf("Erro ao carregar config de teste: %v", err)
	}

	// Conectar ao banco de teste
	database.ConnectTest()
	db := database.DB

	// Resetar schema
	if err := db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;").Error; err != nil {
		t.Fatalf("Erro ao resetar schema: %v", err)
	}

	// Migrar tabelas
	if err := db.AutoMigrate(&models.Image{}); err != nil {
		t.Fatalf("Erro ao migrar DB: %v", err)
	}

	image := models.Image{
		URL:       "/public/img/4x2/432x225-01.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&image)

	return db
}

// Teste do service GetAllArticles
func TestGetAllImages(t *testing.T) {
	db := setupTestDBImage(t)

	image, err := services.GetAllImages(db)
	if err != nil {
		t.Fatalf("Erro ao obter artigos: %v", err)
	}

	if len(image) != 1 {
		t.Fatalf("Esperava 1 artigo, mas obteve %d", len(image))
	}

	if image[0].URL != "/public/img/4x2/432x225-01.jpg" {
		t.Errorf("Esperava '/public/img/4x2/432x225-01.jpg', mas obteve %s", image[0].URL)
	}
}
