package controllers

import (
	"encoding/json"
	"log"

	"github.com/johneliud/Groupie-Tracker/api"
)

func GetLocations() ([]api.Location, error) {
	body, err := api.FetchData(api.LocationURL)
	if err != nil {
		log.Printf("Error fetching %s: %s", api.LocationURL, err)
	}

	// Unmarshal the JSON into a slice of locations structs
	var locations struct {
		Index []api.Location `json:"index"`
	}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %s", err)
	}
	return locations.Index, nil
}
