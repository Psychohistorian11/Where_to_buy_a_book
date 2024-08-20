package BackendApplication

import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func SearchTornamesa(nameBook string, maxResults int) []Book {
	TO := colly.NewCollector()
	var books []Book
	resultCount := 0

	TO.OnHTML("div.portada", func(e *colly.HTMLElement) {
		if resultCount < maxResults {
			link := e.ChildAttr("a", "href")
			completeLink := "https://www.tornamesa.co/" + link
			book := SearchInternalFeatures(completeLink)
			books = append(books, book)
			resultCount++
		}
	})

	nameBookWithoutSpaces := replaceSpaces(nameBook)

	err := TO.Visit("https://www.tornamesa.co/busqueda/listaLibros.php?tipoBus=full&palabrasBusqueda=" + nameBookWithoutSpaces)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books

}

func replaceSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "+")
}

func SearchInternalFeatures(link string) Book {
	TO2 := colly.NewCollector()
	var details string
	var book Book

	TO2.OnHTML("div.content.col-md-12", func(e *colly.HTMLElement) {
		title := cleanText(e.ChildText("h1#titulo"))
		author := cleanText(e.ChildText("p#autor"))
		LongPrice := cleanText(e.ChildText("span.despues"))
		price := strings.Split(LongPrice, " ")[1]
		img := cleanText(e.ChildAttr("a", "href"))

		// Capturando detalles específicos
		editorial := cleanText(e.DOM.Find("dd").Eq(0).Text())
		year := cleanText(e.DOM.Find("dd").Eq(1).Text())
		materia := cleanText(e.DOM.Find("dd").Eq(2).Text())
		idioma := cleanText(e.DOM.Find("dd").Eq(4).Text())
		cubierta := cleanText(e.DOM.Find("dd").Eq(5).Text())

		// Concatenando todos los detalles en un solo string
		details = "Editorial: " + editorial + ", Año de edición: " + year + ", Materia: " + materia + ", Idioma: " + idioma + ", Cubierta: " + cubierta

		book = Book{
			Title:   title,
			Author:  author,
			Price:   price,
			Img:     img,
			Link:    link,
			Details: details,
		}

	})
	err := TO2.Visit(link)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}
	return book
}
