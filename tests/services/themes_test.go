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
func setupTestDBTheme(t *testing.T) *gorm.DB {
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
	if err := db.AutoMigrate(&models.Theme{}); err != nil {
		t.Fatalf("Erro ao migrar DB: %v", err)
	}

	theme := models.Theme{
		Theme:     "thecnolagy",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&theme)

	return db
}

// Teste do service GetAllArticles
func TestGetAllThemes(t *testing.T) {
	db := setupTestDBTheme(t)

	theme, err := services.GetAllThemes(db)
	if err != nil {
		t.Fatalf("Erro ao obter artigos: %v", err)
	}

	if len(theme) != 1 {
		t.Fatalf("Esperava 1 artigo, mas obteve %d", len(theme))
	}

	if theme[0].Theme != "thecnolagy" {
		t.Errorf("Esperava 'thecnolagy', mas obteve %s", theme[0].Theme)
	}
}
