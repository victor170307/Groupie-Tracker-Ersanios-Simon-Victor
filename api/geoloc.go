package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const geolocURL = "https://nominatim.openstreetmap.org/search"

// Structure pour décoder la réponse de Nominatim
type GeoLocation struct {
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
}

// GetCoordinates convertit une ville en coordonnées
func GetCoordinates(location string) (GeoLocation, error) {
	// 1. Nettoyage du nom : "new_york-usa" -> "New York, USA"
	cleanLoc := strings.ReplaceAll(location, "_", " ")
	cleanLoc = strings.ReplaceAll(cleanLoc, "-", ", ")

	// 2. Préparation de l'URL avec encodage sécurisé
	params := url.Values{}
	params.Add("q", cleanLoc)
	params.Add("format", "json")
	params.Add("limit", "1") // On ne veut que le premier résultat

	fullURL := fmt.Sprintf("%s?%s", geolocURL, params.Encode())

	// 3. Création de la requête (OBLIGATOIRE : User-Agent)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return GeoLocation{}, err
	}

	// ⚠️ MODIFIE CECI avec ton nom ou email (ex: "GroupieTracker-StudentProject")
	req.Header.Set("User-Agent", "GroupieTracker-Student/1.0 (ton-email@exemple.com)")

	// 4. Envoi
	resp, err := client.Do(req)
	if err != nil {
		return GeoLocation{}, err
	}
	defer resp.Body.Close()

	// 5. Décodage
	var results []GeoLocation
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return GeoLocation{}, err
	}

	if len(results) > 0 {
		return results[0], nil
	}

	return GeoLocation{}, fmt.Errorf("aucune coordonnée trouvée")
}
