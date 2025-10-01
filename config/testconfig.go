package config

import "os"

// LoadTestConfig define as variáveis de ambiente para testes e carrega a config
func LoadTestConfig() error {
	os.Setenv("APP_ENV", "testing")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "testUser")
	os.Setenv("DB_PASS", "testPass")
	os.Setenv("DB_NAME", "thebloghub_test")
	os.Setenv("APP_PORT", "8081")

	return LoadConfig()
}
