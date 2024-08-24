package controllers

import (
	"encoding/json"
	"log"

	"github.com/johneliud/Groupie-Tracker/api"
)

func GetRelations() ([]api.Relation, error) {
	body, err := api.FetchData(api.RelationURL)
	if err != nil {
		log.Printf("Error fetching %s: %s", api.RelationURL, err)
	}

	// Unmarshal the JSON into a slice of Relation structs
	var relations struct {
		Index []api.Relation `json:"index"`
	}
	err = json.Unmarshal(body, &relations)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %s", err)
	}
	return relations.Index, nil
}
