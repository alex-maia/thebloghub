package database

import (
	"fmt"
	"log"

	"thebloghub/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect conecta de acordo com config.Cfg
func Connect() {
	if config.Cfg.AppEnv == "testing" {
		connectToDB("DB de teste")
	} else {
		connectToDB("DB principal")
	}
}

// ConnectTest força conexão com o banco de teste
func ConnectTest() {
	if !config.IsTesting() {
		log.Fatal("ConnectTest deve ser chamado apenas no modo testing")
	}
	connectToDB("DB de teste")
}

// Função interna de conexão
func connectToDB(label string) {
	connectionDB := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Cfg.DBHost,
		config.Cfg.DBUser,
		config.Cfg.DBPass,
		config.Cfg.DBName,
		config.Cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(connectionDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao %s: %v", label, err)
	}

	DB = db
	log.Printf("%s conectada com sucesso", label)
}
