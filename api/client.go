package api

import (
	"encoding/json"
	"groupie-tracker-gui/models"
	"io"
	"net/http"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

type relationIndex struct {
	Index []models.Relation `json:"index"`
}

func FetchData() ([]models.Artist, map[int]models.Relation, error) {
	// 1. Artistes
	respArt, err := http.Get(artistsURL)
	if err != nil {
		return nil, nil, err
	}
	defer respArt.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(respArt.Body).Decode(&artists); err != nil {
		return nil, nil, err
	}

	// 2. Relations
	respRel, err := http.Get(relationsURL)
	if err != nil {
		return artists, nil, err
	}
	defer respRel.Body.Close()

	bodyBytes, _ := io.ReadAll(respRel.Body)
	var relIndex relationIndex
	if err := json.Unmarshal(bodyBytes, &relIndex); err != nil {
		return artists, nil, err
	}

	// 3. Transformation en Map
	relMap := make(map[int]models.Relation)
	for _, rel := range relIndex.Index {
		relMap[rel.ID] = rel
	}

	return artists, relMap, nil
}
