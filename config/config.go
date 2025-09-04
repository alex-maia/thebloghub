package config

import (
    "log"
    "os"

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

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️  Nenhum .env encontrado, a usar variáveis de ambiente do sistema.")
    }

    Cfg = Config{
        AppName:   getEnv("APP_NAME", ""),
        AppEnv:    getEnv("APP_ENV", ""),
        AppPort:   getEnv("APP_PORT", ""),
        DBHost:    getEnv("DB_HOST", ""),
        DBPort:    getEnv("DB_PORT", ""),
        DBUser:    getEnv("DB_USER", ""),
        DBPass:    getEnv("DB_PASS", ""),
        DBName:    getEnv("DB_NAME", ""),
        SecretKey: getEnv("SECRET_KEY", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
