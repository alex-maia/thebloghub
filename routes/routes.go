package routes

import (
    "net/http"
	"thebloghub/controllers/frontend"
)

func RegisterRoutes() {
    // Frontend
    http.HandleFunc("/", frontend.HomeHandler)
}