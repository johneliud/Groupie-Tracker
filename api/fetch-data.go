package api

import (
	"io"
	"log"
	"net/http"
)

// Fetches data from an API and returns it as a slice of bytes.
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %s", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}
	return body, nil
}
