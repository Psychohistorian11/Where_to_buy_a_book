package BackendApplication

import (
	"fmt"
	"github.com/gocolly/colly"
)

func HandleFormData(nameBook string) {
	searchBuscaLibre(nameBook)
	searchLibreriaNacional(nameBook)
	searchTornamesa(nameBook)
	searchPanamericana(nameBook)
}

func searchBuscaLibre(nameBook string) {
	BL := colly.NewCollector()
	var bookTitle string
	var bookPrice string

	found := false
	found2 := false

	BL.OnHTML("h3.nombre.margin-top-10.text-align-left", func(e *colly.HTMLElement) {
		if !found {
			bookTitle = e.Text
			found = true
		}
	})

	BL.OnHTML("p.precio-ahora.hide-on-hover.margin-0.font-size-medium", func(e *colly.HTMLElement) {
		if !found2 {
			bookPrice = e.Text
			found2 = true
			fmt.Printf("Título: %s\nPrecio: %s\n", bookTitle, bookPrice)
		}
	})

	err := BL.Visit("https://www.buscalibre.com.co/libros/search?q=" + nameBook)
	if err != nil {
		return
	}
}

func searchLibreriaNacional(nameBook string) {}
func searchTornamesa(nameBook string)        {}
func searchPanamericana(nameBook string)     {}
