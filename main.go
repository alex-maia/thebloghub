package main

import (
	"log"
	"net/http"
	"thebloghub/config"
	"thebloghub/database"
	"thebloghub/models"
	"thebloghub/routes"
	"thebloghub/seed"
)

func main() {
	// Carregar e validar configs
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Iniciar a base de dados
	database.Connect()

	// AutoMigrate cria/atualiza tabelas
	if err := database.DB.AutoMigrate(&models.Image{}, &models.Article{}); err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	log.Println("Tabelas migradas com sucesso!")

	// Rodar seed
	seed.Run()

	// Definir rotas
	routes.RegisterRoutes()

	// Servir ficheiros estáticos em /public/
	staticFiles := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", staticFiles))

	// Iniciar servidor
	addr := ":" + config.Cfg.AppPort
	log.Printf("Servidor %s iniciado em http://localhost%s\n", config.Cfg.AppName, addr)
	log.Printf("Ambiente: %s\n", config.Cfg.AppEnv)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
