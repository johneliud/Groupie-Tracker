package controllers

import (
	"encoding/json"
	"log"

	"github.com/johneliud/Groupie-Tracker/api"
)

func GetArtists() ([]api.Artist, error) {
	body, err := api.FetchData(api.ArtistURL)
	if err != nil {
		log.Printf("Error fetching %s: %s", api.ArtistURL, err)
	}

	// Unmarshal the JSON into a slice of Artist structs
	var artists []api.Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %s", err)
	}
	return artists, nil
}
