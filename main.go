package main

import (
	"html/template"
	"log"
	"net/http"
)

func selectSocietyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/select_society.html"))
	tmpl.Execute(w, nil)
}

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", selectSocietyHandler)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
