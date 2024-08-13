package BackendApplication

import (
	"github.com/gocolly/colly"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Book struct {
	Title   string
	Author  string
	Price   string
	Details string
	Img     string
	Link    string
}

type BooksData struct {
	BooksFromBuscaLibre       []Book
	BooksFromLibreriaNacional []Book
	BooksFromTornamesa        []Book
	BooksFromPanamericana     []Book
}

func HandleFormData(w http.ResponseWriter, nameBook string, maxResults int) {
	booksData := BooksData{
		BooksFromBuscaLibre: searchBuscaLibre(nameBook, maxResults),
		//BooksFromLibreriaNacional: searchLibreriaNacional(nameBook, maxResults),
		//BooksFromTornamesa:        searchTornamesa(nameBook, maxResults),
		//BooksFromPanamericana:     searchPanamericana(nameBook, maxResults),
	}

	renderHTML(w, booksData)
}

func searchBuscaLibre(nameBook string, maxResults int) []Book {
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

	// Visita la página y maneja errores
	err := BL.Visit("https://www.buscalibre.com.co/libros/search?q=" + nameBook)
	if err != nil {
		log.Fatalf("Error visiting page: %v", err)
	}

	return books
}

/**func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Printf("Título: %s\nAutor: %s\nPrecio: %s\nDetalles: %s\nImagen: %s\nEnlace: %s\n\n",
			book.Title, book.Author, book.Price, book.Details, book.Img, book.Link)
	}
}**/

func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", "")
	return text
}

func searchLibreriaNacional(nameBook string, maxResults int) {}
func searchTornamesa(nameBook string, maxResults int)        {}
func searchPanamericana(nameBook string, maxResults int)     {}

func renderHTML(w http.ResponseWriter, data BooksData) {
	tmpl, err := template.ParseFiles("Interface/BooksInStock.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Fatalf("Error parsing template: %v", err)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Fatalf("Error executing template: %v", err)
		return
	}

	log.Println("HTML rendered and sent successfully!")
}
