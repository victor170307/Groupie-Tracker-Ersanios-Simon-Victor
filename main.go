package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker-gui/api"
	"groupie-tracker-gui/ui"
)

func main() {
	// Fetch artists once at startup
	artists, err := api.FetchArtists()
	if err != nil {
		log.Fatal("Failed to fetch artists:", err)
	}

	// Serve static HTML and handle API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ui.ServeHTML(w, artists)
	})

	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		ui.ServeArtistsJSON(w, artists)
	})

	addr := ":8080"
	fmt.Printf("🎸 Groupie Tracker running at http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
