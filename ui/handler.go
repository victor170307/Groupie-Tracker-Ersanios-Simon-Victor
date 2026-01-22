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

func buildListScreen(win fyne.Window, artists []models.Artist, relations map[int]models.Relation) fyne.CanvasObject {
	var filtered = artists

	// Grille adaptative
	grid := container.NewGridWithColumns(3)

	// Fonction de rafraîchissement de la grille
	refreshGrid := func() {
		grid.Objects = nil // On vide
		for _, artist := range filtered {
			a := artist

			card := makeBigCard(a, func() {
				win.SetContent(buildDetailScreen(win, a, relations[a.ID], artists, relations))
			})
			grid.Add(card)
		}
		grid.Refresh()
	}

	refreshGrid()

	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste, une date, un concert...")
	search.OnChanged = func(s string) {
		filtered = filterArtists(artists, s)
		refreshGrid()
	}

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

func makeBigCard(artist models.Artist, onClick func()) fyne.CanvasObject {
	// 1. Image
	img := AsyncImage(artist.Image, fyne.NewSize(180, 180))
	imgContainer := container.NewCenter(img)

	nameLabel := widget.NewLabel(artist.Name)
	nameLabel.Alignment = fyne.TextAlignCenter
	nameLabel.TextStyle = fyne.TextStyle{Bold: true}
	nameLabel.Wrapping = fyne.TextTruncate

	dateLabel := widget.NewLabel(fmt.Sprintf("Est. %d", artist.CreationDate))
	dateLabel.Alignment = fyne.TextAlignCenter
	dateLabel.TextStyle = fyne.TextStyle{Italic: true}

	contentVBox := container.NewVBox(
		imgContainer,
		nameLabel,
		dateLabel,
	)

	bg := canvas.NewRectangle(color.NRGBA{R: 60, G: 65, B: 80, A: 255})
	bg.CornerRadius = 12
	visualCard := container.NewMax(
		bg,
		container.NewPadded(contentVBox),
	)
	return NewClickableCard(visualCard, onClick)
}

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

func (c *ClickableCard) Tapped(_ *fyne.PointEvent) {
	if c.onTap != nil {
		c.onTap()
	}
}

func (c *ClickableCard) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.content)
}

func (c *ClickableCard) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}
