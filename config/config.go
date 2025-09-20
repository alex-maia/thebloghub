package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	AppEnv    string
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	SecretKey string
}

var Cfg Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Nenhum .env encontrado, a usar variáveis de ambiente do sistema.")
	}

	Cfg = Config{
		AppName:   getEnv("APP_NAME", "TheBlogHub"),
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", ""),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", ""),
		SecretKey: getEnv("SECRET_KEY", ""),
	}

	// Validar configuração
	if err := ValidateConfig(); err != nil {
		return fmt.Errorf("erro na validação da configuração: %w", err)
	}

	log.Printf("Configuração carregada com sucesso - Ambiente: %s, Porta: %s", Cfg.AppEnv, Cfg.AppPort)
	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// ValidateConfig valida se todas as configurações obrigatórias estão definidas
func ValidateConfig() error {
	var missingVars []string

	// Variáveis obrigatórias para produção
	if Cfg.AppEnv == "production" {
		requiredVars := map[string]string{
			"DB_USER":    Cfg.DBUser,
			"DB_PASS":    Cfg.DBPass,
			"DB_NAME":    Cfg.DBName,
			"SECRET_KEY": Cfg.SecretKey,
		}

		for varName, varValue := range requiredVars {
			if strings.TrimSpace(varValue) == "" {
				missingVars = append(missingVars, varName)
			}
		}
	}

	// Validações gerais
	if Cfg.AppPort == "" {
		missingVars = append(missingVars, "APP_PORT")
	}

	// Validações específicas
	if Cfg.AppEnv != "development" && Cfg.AppEnv != "production" && Cfg.AppEnv != "testing" {
		return errors.New("APP_ENV deve ser 'development', 'production' ou 'testing'")
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("variáveis de ambiente obrigatórias não definidas: %s", strings.Join(missingVars, ", "))
	}

	// Validações de formato
	if Cfg.SecretKey != "" && len(Cfg.SecretKey) < 32 {
		log.Println("SECRET_KEY deve ter pelo menos 32 caracteres para maior segurança")
	}

	return nil
}

// IsDevelopment retorna true se a aplicação está em modo de desenvolvimento
func IsDevelopment() bool {
	return Cfg.AppEnv == "development"
}

// IsProduction retorna true se a aplicação está em modo de produção
func IsProduction() bool {
	return Cfg.AppEnv == "production"
}

// IsTesting retorna true se a aplicação está em modo de teste
func IsTesting() bool {
	return Cfg.AppEnv == "testing"
}
