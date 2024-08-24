package controllers

import (
	"encoding/json"
	"log"

	"github.com/johneliud/Groupie-Tracker/api"
)

func GetDates() ([]api.Date, error) {
	body, err := api.FetchData(api.DateURL)
	if err != nil {
		log.Printf("Error fetching %s: %s", api.DateURL, err)
	}

	// Unmarshal the JSON into a slice of dates structs
	var dates struct {
		Index []api.Date `json:"index"`
	}
	err = json.Unmarshal(body, &dates)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %s", err)
	}
	return dates.Index, nil
}
