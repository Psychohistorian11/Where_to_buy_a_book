package BackendApplication

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func SearchBuscaLibre(nameBook string, maxResults int) []Book {
	BL := colly.NewCollector()
	var books []Book
	resultCount := 0

	// Crear un nuevo Book por cada resultado y almacenarlo en el slice
	BL.OnHTML("div.box-producto", func(e *colly.HTMLElement) {
		if resultCount < maxResults {
			author := cleanText(e.ChildText("div.autor"))
			details := cleanText(e.ChildText("div.autor.color-dark-gray.metas.hide-on-hover"))

			if strings.Contains(author, details) {
				author = strings.Replace(author, details, "", -1)
			}

			book := Book{
				Title:   cleanText(e.ChildText("a h3.nombre")),
				Author:  author,
				Price:   cleanText(e.ChildText("p.precio-ahora.hide-on-hover.margin-0.font-size-medium")),
				Details: details,
				Img:     e.ChildAttr("img", "data-src"),
				Link:    e.ChildAttr("a", "href"),
			}
			books = append(books, book)
			resultCount++
		}
	})

	err := BL.Visit("https://www.buscalibre.com.co/libros/search?q=" + nameBook)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books
}
