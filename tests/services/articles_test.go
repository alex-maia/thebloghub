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
func setupTestDB(t *testing.T) *gorm.DB {
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
	if err := db.AutoMigrate(&models.Image{}, &models.Theme{}, &models.Article{}); err != nil {
		t.Fatalf("Erro ao migrar DB: %v", err)
	}

	// Inserir seed mínima para teste
	image := models.Image{URL: "/public/img/test.jpg"}
	theme := models.Theme{Theme: "techonlogy"}
	db.Create(&image)
	db.Create(&theme)

	article := models.Article{
		Title:     "Artigo de Teste",
		Lead:      "Lead Teste",
		Text:      "Texto do artigo de teste",
		ImageID:   &image.ID,
		ThemeID:   &theme.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&article)

	return db
}

// Teste do service GetAllArticles
func TestGetAllArticles(t *testing.T) {
	db := setupTestDB(t)

	articles, err := services.GetAllArticles(db)
	if err != nil {
		t.Fatalf("Erro ao obter artigos: %v", err)
	}

	if len(articles) != 1 {
		t.Fatalf("Esperava 1 artigo, mas obteve %d", len(articles))
	}

	if articles[0].Title != "Artigo de Teste" {
		t.Errorf("Esperava 'Artigo de Teste', mas obteve %s", articles[0].Title)
	}
}
