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
	"fyne.io/fyne/v2/driver/desktop"
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

	// Grille adaptative
	grid := container.NewGridWithColumns(3)

	// Fonction de rafraîchissement de la grille
	refreshGrid := func() {
		grid.Objects = nil // On vide
		for _, artist := range filtered {
			a := artist // Capture de variable pour la closure

			// On crée la carte avec notre nouveau système sans bouton superposé
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

	// Titre
	title := canvas.NewText("SoundTrap Collection", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 28
	title.Alignment = fyne.TextAlignCenter

	return container.NewBorder(
		container.NewVBox(container.NewPadded(title), container.NewPadded(search)),
		nil, nil, nil,
		container.NewVScroll(container.NewPadded(grid)),
	)
}

// makeBigCard : Crée une carte visuelle et l'emballe dans notre ClickableCard
func makeBigCard(artist models.Artist, onClick func()) fyne.CanvasObject {
<<<<<<< HEAD
	// 1. Image
	img := AsyncImage(artist.Image, fyne.NewSize(180, 180))
=======
	img := AsyncImage(artist.Image, fyne.NewSize(220, 220))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillContain

>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562
	imgContainer := container.NewCenter(img)
	imgContainer.Resize(fyne.NewSize(240, 240))

<<<<<<< HEAD
	// 2. Textes
	nameLabel := widget.NewLabel(artist.Name)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Wrapping = fyne.TextTruncate
=======
	nameLabel := widget.NewLabel(artist.Name)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Wrapping = fyne.TextWrapWord
>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562

	dateLabel := widget.NewLabel(fmt.Sprintf("Créé en %d", artist.CreationDate))
	dateLabel.Alignment = fyne.TextAlignCenter
	dateLabel.TextStyle = fyne.TextStyle{Italic: true}

<<<<<<< HEAD
	// 3. Empilement Vertical
=======
>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562
	contentVBox := container.NewVBox(
		imgContainer,
		nameLabel,
		dateLabel,
	)

	bg := canvas.NewRectangle(color.NRGBA{R: 60, G: 65, B: 80, A: 255})
	bg.CornerRadius = 12

<<<<<<< HEAD
	// 5. Construction visuelle (Fond + Contenu)
	// On ne met PAS de bouton ici pour éviter le voile gris au survol
	visualCard := container.NewMax(
		bg,
		container.NewPadded(contentVBox),
	)

	// 6. On rend le tout cliquable proprement
	return NewClickableCard(visualCard, onClick)
=======
	btn := widget.NewButton("", onClick)

	paddedContent := container.NewPadded(contentVBox)

	return container.NewMax(bg, paddedContent, btn)
>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562
}

// --- ECRAN 2 : DETAILS ---

func buildDetailScreen(win fyne.Window, artist models.Artist, relation models.Relation, allArtists []models.Artist, allRelations map[int]models.Relation) fyne.CanvasObject {

	btnBack := widget.NewButtonWithIcon("Retour à la liste", theme.NavigateBackIcon(), func() {
		win.SetContent(buildListScreen(win, allArtists, allRelations))
	})

	img := AsyncImage(artist.Image, fyne.NewSize(300, 300))

	name := canvas.NewText(artist.Name, color.White)
	name.TextStyle = fyne.TextStyle{Bold: true}
	name.TextSize = 24
	name.Alignment = fyne.TextAlignCenter

	meta := widget.NewLabel(fmt.Sprintf("Création: %d  •  1er Album: %s", artist.CreationDate, artist.FirstAlbum))
	meta.Alignment = fyne.TextAlignCenter

	concertsContainer := container.NewVBox()
	concertsContainer.Add(widget.NewLabelWithStyle("Concerts & Lieux", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))

	if len(relation.DatesLocations) > 0 {
		for loc, dates := range relation.DatesLocations {
			locationName := loc
			prettyName := strings.Title(strings.ReplaceAll(strings.ReplaceAll(locationName, "_", " "), "-", ", "))

			btnGeo := widget.NewButton("📍 "+prettyName, func() {
				go func() {
					log.Println("GPS Request:", locationName)
					coords, err := api.GetCoordinates(locationName)
					if err != nil {
						log.Println("Erreur GPS:", err)
						return
					}
					mapURL := fmt.Sprintf("https://www.openstreetmap.org/?mlat=%s&mlon=%s#map=12/%s/%s", coords.Lat, coords.Lon, coords.Lat, coords.Lon)
					u, _ := url.Parse(mapURL)
					fyne.CurrentApp().OpenURL(u)
				}()
			})

			lblDates := widget.NewLabel("📅 " + strings.Join(dates, ", "))
			lblDates.TextStyle = fyne.TextStyle{Italic: true}
			lblDates.Wrapping = fyne.TextWrapWord

			concertsContainer.Add(btnGeo)
			concertsContainer.Add(lblDates)
			concertsContainer.Add(layout.NewSpacer())
		}
	} else {
		concertsContainer.Add(widget.NewLabel("Aucune date prévue."))
	}

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

// --- WIDGET CLIQUABLE PERSONNALISÉ ---
// Ce widget remplace le bouton invisible. Il capture le clic sans modifier le visuel (pas de hover gris).

type ClickableCard struct {
	widget.BaseWidget
	content fyne.CanvasObject
	onTap   func()
}

func NewClickableCard(content fyne.CanvasObject, onTap func()) *ClickableCard {
	c := &ClickableCard{content: content, onTap: onTap}
	c.ExtendBaseWidget(c)
	return c
}

// Tapped implémente l'interface fyne.Tappable (déclenche le clic)
func (c *ClickableCard) Tapped(_ *fyne.PointEvent) {
	if c.onTap != nil {
		c.onTap()
	}
}

// CreateRenderer dit à Fyne comment dessiner le widget (juste afficher le contenu)
func (c *ClickableCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.content)
}

// Cursor change la souris en "main" au survol (Interface desktop.Cursorable)
func (c *ClickableCard) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
