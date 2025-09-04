package main

import (
    "log"
    "net/http"
    "thebloghub/config"
    "thebloghub/routes"
)

func main() {
    // Carregar configs
    config.LoadConfig()

    // Definir rotas
    routes.RegisterRoutes()

    // Iniciar servidor
    addr := ":" + config.Cfg.AppPort
    log.Printf("Servidor %s iniciado em http://localhost%s\n", config.Cfg.AppName, addr)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        log.Fatal(err)
    }
}