package seed

import (
	"log"
	"thebloghub/database"
	"thebloghub/models"
)

func Run() {
	// Cria algumas imagens
	images := []models.Image{
		{URL: "/public/img/4x2/432x225-01.jpg"},
		{URL: "/public/img/4x2/432x225-02.jpg"},
		{URL: "/public/img/4x2/432x225-03.jpg"},
		{URL: "/public/img/4x2/660x267-01.jpg"},
		{URL: "/public/img/4x2/660x267-02.jpg"},
		{URL: "/public/img/4x2/660x267-03.jpg"},
		{URL: "/public/img/4x2/660x267-04.jpg"},
		{URL: "/public/img/4x2/1344x504_01.jpg"},
	}

	for i := range images {
		if err := database.DB.FirstOrCreate(&images[i], models.Image{URL: images[i].URL}).Error; err != nil {
			log.Println("Erro ao criar imagem:", err)
		}
	}

	// Cria alguns artigos associados às imagens
	articles := []models.Article{
		{
			Title:   "The Rise of AI: Transforming the Future",
			Lead:    "LEAD: The Rise of AI: Transforming the Future",
			Text:    "TEXT: Conteúdo do primeiro artigo.",
			ImageID: &images[0].ID, // índice válido
		},
	}

	for i, a := range articles {
		a.ImageID = &images[i%len(images)].ID
		database.DB.FirstOrCreate(&a, models.Article{Lead: a.Lead})
	}

	log.Println("Seed concluída com sucesso!")
}
