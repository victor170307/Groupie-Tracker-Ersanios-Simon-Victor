package main

import (
	"log"

	"groupie-tracker-gui/api"
	"groupie-tracker-gui/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	// 1. Initialisation des données
	// On récupère les Artistes ET les Relations (dates/lieux) au démarrage
	log.Println("Chargement des données API...")
	artists, relations, err := api.FetchData()
	if err != nil {
		// On log l'erreur mais on ne crash pas (l'app s'ouvrira vide ou partielle)
		log.Println("Erreur lors du chargement des données:", err)
	} else {
		log.Println("Données chargées avec succès !")
	}

	// 2. Création de l'application Fyne
	application := app.New()

	// Tu peux définir un ID unique pour sauvegarder les préférences (optionnel)
	application.Settings().SetTheme(nil) // Utilise le thème système par défaut

	// 3. Création de la fenêtre principale
	window := application.NewWindow("SoundTrap - Groupie Tracker")

	// 4. Configuration du contenu
	// On passe 'window' pour permettre le changement de page (navigation)
	// On passe 'artists' et 'relations' pour remplir l'interface
	content := ui.NewArtistScreen(window, artists, relations)
	window.SetContent(content)

	// 5. Définition de la taille et lancement
	window.Resize(ui.DefaultWindowSize())
	window.CenterOnScreen()
	window.ShowAndRun()
}
