package main

import (
	"log"
	"net/http"

	"github.com/johneliud/Groupie-Tracker/handlers"
)

func main() {
	http.HandleFunc("/static/", handlers.ServeStatic)
	http.HandleFunc("/", handlers.ServeIndex)

	port := ":9000"

	log.Printf("Server started at http://localhost%v", port)
	http.ListenAndServe(port, nil)
}
