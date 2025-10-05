package routes

import (
	"net/http"
	"thebloghub/controllers/backend"
	"thebloghub/controllers/frontend"
)

func RegisterRoutes() {
	// Frontend
	http.HandleFunc("/", frontend.HomeHandler)
	http.HandleFunc("/admin", backend.HomeHandler)
}
