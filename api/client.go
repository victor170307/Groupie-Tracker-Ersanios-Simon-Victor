package api

import (
	"encoding/json"
	"net/http"

	"groupie-tracker-gui/models"
)

const baseURL = "https://groupietrackers.herokuapp.com/api/artists"

func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

