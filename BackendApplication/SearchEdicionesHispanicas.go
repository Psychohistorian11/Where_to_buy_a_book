package BackendApplication

import (
	"github.com/gocolly/colly"
	"log"
)

func SearchEdicionesHispanicas(nameBook string, maxResults int) []Book {
	EH := colly.NewCollector()
	var books []Book
	resultCount := 0

	EH.OnHTML("div.product-wrapper", func(e *colly.HTMLElement) {
		if resultCount < maxResults {
			link := e.ChildAttr("a", "href")
			book := SearchInternalFeaturesEH(link)
			books = append(books, book)
			resultCount++

		}
	})
	newName := replaceSpaces(nameBook)
	link := "https://edicioneshispanicas.com/?s=" + newName + "&post_type=product"
	err := EH.Visit(link)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books

}

func SearchInternalFeaturesEH(link string) Book {
	EH2 := colly.NewCollector()
	var details string
	var book Book

	EH2.OnHTML("div.row.product-image-summary-inner", func(e *colly.HTMLElement) {

		title := cleanText(e.ChildText("h1.product_title.entry-title.wd-entities-title"))
		author := cleanText(e.DOM.Find("span:contains('Gabriel García Márquez')").Text())
		price := cleanText(e.ChildText("span.woocommerce-Price-amount.amount"))
		img := e.ChildAttr("img", "src")

		editorial := cleanText(e.DOM.Find("th:contains('Editorial')").Next().Text())
		supplier := cleanText(e.DOM.Find("th:contains('Proveedor')").Next().Text())
		language := cleanText(e.DOM.Find("th:contains('Idioma')").Next().Text())
		presentation := cleanText(e.DOM.Find("th:contains('Presentación')").Next().Text())
		numPages := cleanText(e.DOM.Find("th:contains('Número de páginas')").Next().Text())

		// Concatenando todos los detalles en un solo string
		details = "Editorial: " + editorial + ", Proveedor: " + supplier + ", Idioma: " + language + ", Cubierta: " + presentation + ", Páginas: " + numPages

		book = Book{
			Title:   title,
			Author:  author,
			Price:   price,
			Img:     img,
			Link:    link,
			Details: details,
		}
	})

	err := EH2.Visit(link)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}
	return book
}
