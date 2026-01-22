package main

import (
	"log"

	"groupie-tracker-gui/api"
	"groupie-tracker-gui/ui"

	"fyne.io/fyne/v2/app"
)

func main() {

	log.Println("Chargement des données API...")
	artists, relations, err := api.FetchData()
	if err != nil {

		log.Println("Erreur lors du chargement des données:", err)
	} else {
		log.Println("Données chargées avec succès !")
	}

	// 2. Création de l'application Fyne
	application := app.New()

	application.Settings().SetTheme(nil)

	window := application.NewWindow("SoundTrap - Groupie Tracker")

	content := ui.NewArtistScreen(window, artists, relations)
	window.SetContent(content)

	window.Resize(ui.DefaultWindowSize())
	window.CenterOnScreen()
	window.ShowAndRun()
}
