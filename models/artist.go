package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"` // "Image" avec I majuscule
	Name         string   `json:"name"`  // "Name" avec N majuscule
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
