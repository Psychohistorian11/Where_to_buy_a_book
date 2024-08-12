package main

import (
	"Where_to_buy_a_book/BackendApplication"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./Interface"))
	http.Handle("/", fs)

	http.HandleFunc("/process", processForm)

	// Iniciar el servidor en el puerto 4000 del localhost
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		return
	}
}
func processForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")

		BackendApplication.HandleFormData(name)
	}
}
