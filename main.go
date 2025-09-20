package main

import (
	"log"
	"net/http"
	"thebloghub/config"
	"thebloghub/routes"
)

func main() {
	// Carregar e validar configs
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Definir rotas
	routes.RegisterRoutes()

	// Servir ficheiros estáticos em /assets/
	staticFiles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", staticFiles))

	// Iniciar servidor
	addr := ":" + config.Cfg.AppPort
	log.Printf("Servidor %s iniciado em http://localhost%s\n", config.Cfg.AppName, addr)
	log.Printf("Ambiente: %s\n", config.Cfg.AppEnv)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
