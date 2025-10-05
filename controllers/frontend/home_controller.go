package frontend

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
	master := filepath.Join("resources", "views", "frontend", "layouts", "masterpage.gohtml")
	content := filepath.Join("resources", "views", "frontend", "homepage.gohtml")
	header := filepath.Join("resources", "views", "frontend", "shared", "header.gohtml")
	footer := filepath.Join("resources", "views", "frontend", "shared", "footer.gohtml")
	headline := filepath.Join("resources", "views", "frontend", "homepage", "headline.gohtml")
	article := filepath.Join("resources", "views", "frontend", "components", "article.gohtml")

	// Parse de todos os templates
	tmpl, err := template.ParseFiles(master, content, header, footer, headline, article)
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

	// Dados para o template
	data := map[string]interface{}{
		"Title":    "Página Inicial",
		"Articles": articles,
	}

	// Executa o template do layout master
	err = tmpl.ExecuteTemplate(w, "masterpage.gohtml", data)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		return
	}
}
