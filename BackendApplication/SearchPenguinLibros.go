package BackendApplication

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func SearchPenguinLibros(nameBook string, maxResults int) []Book {
	PL := colly.NewCollector()
	var books []Book
	resultCount := 0

	PL.OnHTML("article.x-mot-result", func(e *colly.HTMLElement) {
		if resultCount < maxResults {
			link := e.ChildAttr("a", "href")
			fmt.Println("este es el link: " + link)
			resultCount++
		}
	})

	newName := replaceSpaces_20(nameBook)

	link := "https://www.penguinlibros.com/co/?srsltid=AfmBOorG7pLJ46kdwpG6FfCVvw-nq0LGZJotVq3Fbdo3_26u3NRykoyK&mot_q=" + newName
	err := PL.Visit(link)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books
}
