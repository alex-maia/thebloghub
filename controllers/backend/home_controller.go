package backend

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"thebloghub/database"
	"thebloghub/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Caminhos dos templates
	master := filepath.Join("resources", "views", "backend", "layouts", "masterpage.html")
	content := filepath.Join("resources", "views", "backend", "homepage.html")
	header := filepath.Join("resources", "views", "backend", "shared", "header.html")
	aside := filepath.Join("resources", "views", "backend", "shared", "aside.html")
	footer := filepath.Join("resources", "views", "backend", "shared", "footer.html")
	article := filepath.Join("resources", "views", "backend", "components", "article.html")

	// Parse de todos os templates
	tmpl, err := template.ParseFiles(master, content, header, aside, footer, article)
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		http.Error(w, "Erro ao carregar templates", http.StatusInternalServerError)
		return
	}

	articles, err := services.GetAllArticles(database.DB)
	if err != nil {
		http.Error(w, "Erro ao buscar artigos", http.StatusInternalServerError)
		return
	}
	images, err := services.GetAllImages(database.DB)
	if err != nil {
		http.Error(w, "Erro ao buscar artigos", http.StatusInternalServerError)
		return
	}
	themes, err := services.GetAllThemes(database.DB)
	if err != nil {
		http.Error(w, "Erro ao buscar artigos", http.StatusInternalServerError)
		return
	}

	// Dados para o template
	data := map[string]interface{}{
		"Title":         "Página Inicial",
		"Articles":      articles,
		"ArticlesCount": len(articles),
		"ImagesCount":   len(images),
		"ThemesCount":   len(themes),
	}

	// Executa o template do layout master
	err = tmpl.ExecuteTemplate(w, "masterpage.html", data)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		return
	}
}
