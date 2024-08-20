package BackendApplication

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func SearchLibreriaNacional(nameBook string, maxResults int) []Book {
	LN := colly.NewCollector() // Activa la depuración

	var books []Book

	// Verifica si puedes capturar cualquier cosa del cuerpo

	// Ajuste en el selector
	LN.OnHTML("div.row.mx-0.mx-sm-2.mx-md-3", func(e *colly.HTMLElement) {
		fmt.Println("Página cargada con éxito.")
	})

	// Imprime la URL para estar seguro de que es correcta
	finalURL := "https://www.librerianacional.com/categoria/libros/busqueda/texto/Cien%20a%C3%B1os%20de%20soledad"
	err := LN.Visit(finalURL)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books
}
