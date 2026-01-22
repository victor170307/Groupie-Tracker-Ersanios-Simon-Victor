package main

import (
	"log"

	"groupie-tracker-gui/api"
	"groupie-tracker-gui/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	artists, err := api.FetchArtists()
	if err != nil {
		log.Fatal("Erreur pas d'artiste:", err)
	}

	application := app.New()
	window := application.NewWindow("SoundTrap")
	window.SetContent(ui.NewArtistScreen(artists))
	window.Resize(ui.DefaultWindowSize())
	window.ShowAndRun()
}
