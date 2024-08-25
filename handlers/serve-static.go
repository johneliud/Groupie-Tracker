package handlers

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("static", strings.TrimPrefix(r.URL.Path, "/static/"))

	// Prevent directory inspection
	if !strings.HasPrefix(filePath, "static/") {
		log.Printf("Invalid path: %s", r.URL.Path)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Check if path is a directory
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File not found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		} else {
			log.Printf("Error accessing file: %s", err)
			http.Error(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}
	}

	// Prevent directory listing
	if fileInfo.IsDir() {
		log.Printf("Invalid path: %s", r.URL.Path)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filePath)
}
