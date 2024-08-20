package main

import (
	"Where_to_buy_a_book/BackendApplication"
	"net/http"
	"strconv"
)

func main() {
	fs := http.FileServer(http.Dir("./Interface"))
	http.Handle("/", fs)

	http.HandleFunc("/Books", processForm)

	// Iniciar el servidor en el puerto 4000 del localhost
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		return
	}
}
func processForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		amount := r.FormValue("amount")
		amountInt, _ := strconv.Atoi(amount)

		BackendApplication.HandleFormData(w, name, amountInt)
	}
}
