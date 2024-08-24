package BackendApplication

import (
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
	BooksFromBuscaLibre          []Book
	BooksFromEdicionesHispanicas []Book
	BooksFromTornamesa           []Book
}

func HandleFormData(w http.ResponseWriter, nameBook string, maxResults int) {
	booksData := BooksData{
		BooksFromBuscaLibre:          SearchBuscaLibre(nameBook, maxResults),
		BooksFromEdicionesHispanicas: SearchEdicionesHispanicas(nameBook, maxResults),
		BooksFromTornamesa:           SearchTornamesa(nameBook, maxResults),
	}

	renderHTML(w, booksData)
}

/**func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Printf("Título: %s\nAutor: %s\nPrecio: %s\nDetalles: %s\nImagen: %s\nEnlace: %s\n\n",
			book.Title, book.Author, book.Price, book.Details, book.Img, book.Link)
	}
}**/

func searchTornamesa(nameBook string, maxResults int) {}

func replaceSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "+")
}

func cleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", "")
	return text
}
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
