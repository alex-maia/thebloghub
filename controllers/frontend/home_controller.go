package frontend

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Caminhos dos templates
	master := filepath.Join("resources", "views", "frontend", "layouts", "masterpage.html")
	content := filepath.Join("resources", "views", "frontend", "homepage.html")
	header := filepath.Join("resources", "views", "frontend", "shared", "header.html")
	footer := filepath.Join("resources", "views", "frontend", "shared", "footer.html")
	headline := filepath.Join("resources", "views", "frontend", "homepage", "headline.html")
	article := filepath.Join("resources", "views", "frontend", "components", "article.html")

	// Parse de todos os templates
	tmpl, err := template.ParseFiles(master, content, header, footer, headline, article)
	if err != nil {
		log.Println("Erro ao carregar templates:", err)
		http.Error(w, "Erro ao carregar templates", http.StatusInternalServerError)
		return
	}

	// Dados para o template
	data := map[string]string{
		"Title": "Página Inicial",
	}

	// Executa o template do layout master
	err = tmpl.ExecuteTemplate(w, "masterpage.html", data)
	if err != nil {
		log.Println("Erro ao renderizar template:", err)
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		return
	}
}
