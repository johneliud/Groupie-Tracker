package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/johneliud/Groupie-Tracker/api"
	"github.com/johneliud/Groupie-Tracker/controllers"
)

func ServeArtists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		log.Printf("Error parsing template: %s", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}

	artists, err := controllers.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %s", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}

	type PageData struct {
		Artists []api.Artist
	}

	data := PageData{
		Artists: artists,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
		http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}
}
