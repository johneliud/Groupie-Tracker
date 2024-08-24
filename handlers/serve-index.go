package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %s", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %s", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
	}
}
