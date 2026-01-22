package ui

import (
	"fmt"
	"image/color"
	"log"
	"net/url"
	"strings"

	"groupie-tracker-gui/api"
	"groupie-tracker-gui/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Point d'entrée de l'interface
func NewArtistScreen(win fyne.Window, artists []models.Artist, relations map[int]models.Relation) fyne.CanvasObject {
	return buildListScreen(win, artists, relations)
}

func DefaultWindowSize() fyne.Size {
	return fyne.NewSize(1200, 800)
}

// --- ECRAN 1 : GALERIE D'ARTISTES ---

func buildListScreen(win fyne.Window, artists []models.Artist, relations map[int]models.Relation) fyne.CanvasObject {
	var filtered = artists

	// On utilise une grille avec 3 colonnes.
	grid := container.NewGridWithColumns(3)

	// Fonction pour remplir la grille
	refreshGrid := func() {
		grid.Objects = nil // On vide la grille
		for _, artist := range filtered {
			a := artist

			// Création de la carte robuste
			card := makeBigCard(a, func() {
				win.SetContent(buildDetailScreen(win, a, relations[a.ID], artists, relations))
			})
			grid.Add(card)
		}
		grid.Refresh()
	}

	refreshGrid()

	// Barre de recherche
	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste, une date, un concert...")
	search.OnChanged = func(s string) {
		filtered = filterArtists(artists, s)
		refreshGrid()
	}

	// Titre stylisé
	title := canvas.NewText("SoundTrap Collection", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 28
	title.Alignment = fyne.TextAlignCenter

	return container.NewBorder(
		container.NewVBox(container.NewPadded(title), container.NewPadded(search)),
		nil, nil, nil,
		container.NewVScroll(container.NewPadded(grid)), // Scroll vertical
	)
}

// makeBigCard : Version Corrigée (Layout VBox)
// Cette version empile l'image et le texte pour éviter qu'ils disparaissent
func makeBigCard(artist models.Artist, onClick func()) fyne.CanvasObject {
	// 1. Image
	// On force une taille fixe pour que l'alignement soit joli
	img := AsyncImage(artist.Image, fyne.NewSize(180, 180))

	// On centre l'image dans son conteneur
	imgContainer := container.NewCenter(img)

	// 2. Textes
	// On utilise widget.Label car c'est plus stable que canvas.Text
	nameLabel := widget.NewLabel(artist.Name)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	// Astuce : Si le texte est trop long, il sera coupé proprement
	nameLabel.Wrapping = fyne.TextTruncate

	dateLabel := widget.NewLabel(fmt.Sprintf("Est. %d", artist.CreationDate))
	dateLabel.Alignment = fyne.TextAlignCenter
	dateLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 3. EMPILEMENT VERTICAL (C'est ici que la magie opère)
	// Image -> Nom -> Date
	contentVBox := container.NewVBox(
		imgContainer,
		nameLabel,
		dateLabel,
	)

	// 4. Fond de la carte
	bg := canvas.NewRectangle(color.NRGBA{R: 60, G: 65, B: 80, A: 255})
	bg.CornerRadius = 12

	// 5. Bouton invisible (pour le clic)
	btn := widget.NewButton("", onClick)

	// 6. Assemblage final
	// On ajoute du Padding (marge) pour que le contenu ne touche pas les bords
	paddedContent := container.NewPadded(contentVBox)

	// On superpose : Fond -> Contenu -> Bouton
	return container.NewMax(bg, paddedContent, btn)
}

// --- ECRAN 2 : DETAILS ---

func buildDetailScreen(win fyne.Window, artist models.Artist, relation models.Relation, allArtists []models.Artist, allRelations map[int]models.Relation) fyne.CanvasObject {

	btnBack := widget.NewButtonWithIcon("Retour à la liste", theme.NavigateBackIcon(), func() {
		win.SetContent(buildListScreen(win, allArtists, allRelations))
	})

	// Image géante
	img := AsyncImage(artist.Image, fyne.NewSize(300, 300))

	name := widget.NewLabel(artist.Name)
	name.TextStyle = fyne.TextStyle{Bold: true}
	name.Alignment = fyne.TextAlignCenter

	meta := widget.NewLabel(fmt.Sprintf("Création: %d  •  1er Album: %s", artist.CreationDate, artist.FirstAlbum))
	meta.Alignment = fyne.TextAlignCenter

	// Concerts
	concertsContainer := container.NewVBox()
	concertsContainer.Add(widget.NewLabelWithStyle("Concerts & Lieux", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))

	if len(relation.DatesLocations) > 0 {
		for loc, dates := range relation.DatesLocations {
			locationName := loc
			prettyName := strings.Title(strings.ReplaceAll(strings.ReplaceAll(locationName, "_", " "), "-", ", "))

			// Bouton Map
			btnGeo := widget.NewButton("📍 "+prettyName, func() {
				go func() {
					log.Println("GPS:", locationName)
					coords, err := api.GetCoordinates(locationName)
					if err != nil {
						return
					}
					mapURL := fmt.Sprintf("https://www.openstreetmap.org/?mlat=%s&mlon=%s#map=12/%s/%s", coords.Lat, coords.Lon, coords.Lat, coords.Lon)
					u, _ := url.Parse(mapURL)
					fyne.CurrentApp().OpenURL(u)
				}()
			})

			lblDates := widget.NewLabel("📅 " + strings.Join(dates, ", "))
			lblDates.TextStyle = fyne.TextStyle{Italic: true}

			concertsContainer.Add(btnGeo)
			concertsContainer.Add(lblDates)
			concertsContainer.Add(layout.NewSpacer())
		}
	} else {
		concertsContainer.Add(widget.NewLabel("Aucune date prévue."))
	}

	// Layout Détails
	textScroll := container.NewVScroll(container.NewVBox(
		name,
		meta,
		widget.NewSeparator(),
		widget.NewLabelWithStyle("Membres:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel(strings.Join(artist.Members, "\n")),
		widget.NewSeparator(),
		concertsContainer,
	))

	split := container.NewHSplit(container.NewPadded(img), container.NewPadded(textScroll))
	split.SetOffset(0.35)

	return container.NewBorder(btnBack, nil, nil, nil, split)
}

// --- UTILITAIRES ---

func filterArtists(artists []models.Artist, query string) []models.Artist {
	normQ := strings.ToLower(query)
	if normQ == "" {
		return artists
	}
	var res []models.Artist
	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), normQ) ||
			strings.Contains(strings.ToLower(strings.Join(a.Members, " ")), normQ) ||
			strings.Contains(fmt.Sprint(a.CreationDate), normQ) {
			res = append(res, a)
		}
	}
	return res
}
