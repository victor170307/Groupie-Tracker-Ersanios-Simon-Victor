package main

import (
	"fmt"
	"groupie-tracker-gui/api"
)

func main() {
	artists, err := api.FetchArtists()
	if err != nil {
		fmt.Println("Erreur API:", err)
		return
	}
	fmt.Printf("API OK: %d artistes récupérés\n", len(artists))
	if len(artists) > 0 {
		fmt.Printf("Premier artiste: %s\n", artists[0].Name)
	}
}
