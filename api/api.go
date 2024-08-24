package api

type Artist struct {
	ID            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	CreationDate  int      `json:"creationDate"`
	FirstAlbum    string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	ConcertsDates string   `json:"concertDates"`
	Relations     string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

const (
	ArtistURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL = "https://groupietrackers.herokuapp.com/api/locations"
	DateURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)
