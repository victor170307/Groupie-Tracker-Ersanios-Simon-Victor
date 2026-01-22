package ui

import (
    "fmt"
    "strings"
    "image/color"
    "unicode"

    "groupie-tracker-gui/models"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

const appTitle = "SoundTrap"

func DefaultWindowSize() fyne.Size {
    return fyne.NewSize(960, 720)
}

func NewArtistScreen(artists []models.Artist) fyne.CanvasObject {
    filtered := make([]models.Artist, len(artists))
    copy(filtered, artists)

    bg := canvas.NewLinearGradient(
        color.NRGBA{R: 15, G: 15, B: 20, A: 255},
        color.NRGBA{R: 35, G: 35, B: 45, A: 255},
        90,
    )

    title := widget.NewLabel(appTitle)
    title.TextStyle = fyne.TextStyle{Bold: true}
    title.Alignment = fyne.TextAlignCenter
    title.Wrapping = fyne.TextWrapWord

    search := widget.NewEntry()
    search.SetPlaceHolder("Rechercher par nom, membre ou annee")
    status := widget.NewLabel(formatStatus(len(filtered)))

    cards := container.NewVBox(buildArtistCards(filtered)...)
    scroll := container.NewVScroll(cards)
    scroll.SetMinSize(fyne.NewSize(0, 500))

    search.OnChanged = func(query string) {
        filtered = filterArtists(artists, query)
        status.SetText(formatStatus(len(filtered)))
        cards.Objects = buildArtistCards(filtered)
        cards.Refresh()
    }

    header := container.NewVBox(
        container.NewPadded(title),
        container.NewPadded(search),
    )

    content := container.NewBorder(header, status, nil, nil, scroll)
    padded := container.NewPadded(content)

    return container.New(layout.NewMaxLayout(), bg, padded)
}

func buildArtistCards(artists []models.Artist) []fyne.CanvasObject {
    if len(artists) == 0 {
        msg := canvas.NewText("Aucun artiste trouve", color.NRGBA{R: 220, G: 220, B: 220, A: 255})
        msg.Alignment = fyne.TextAlignCenter
        msg.TextStyle = fyne.TextStyle{Bold: true}
        return []fyne.CanvasObject{container.NewCenter(msg)}
    }

    items := make([]fyne.CanvasObject, 0, len(artists))
    for _, artist := range artists {
        items = append(items, makeArtistCard(artist))
    }
    return items
}

func makeArtistCard(artist models.Artist) fyne.CanvasObject {
    title := canvas.NewText(artist.Name, color.NRGBA{R: 240, G: 240, B: 240, A: 255})
    title.TextStyle = fyne.TextStyle{Bold: true}

    subtitle := canvas.NewText(
        fmt.Sprintf("Membres: %s", strings.Join(artist.Members, ", ")),
        color.NRGBA{R: 200, G: 200, B: 200, A: 255},
    )
    subtitle.TextSize = 12

    meta := canvas.NewText(
        fmt.Sprintf("Creation: %d  •  1er album: %s", artist.CreationDate, artist.FirstAlbum),
        color.NRGBA{R: 170, G: 170, B: 170, A: 255},
    )
    meta.TextSize = 11

    cardBody := container.NewVBox(title, subtitle, meta)

    cardBg := canvas.NewRectangle(color.NRGBA{R: 30, G: 35, B: 45, A: 220})
    cardBg.CornerRadius = 8

    padded := container.NewPadded(cardBody)
    return container.New(layout.NewMaxLayout(), cardBg, padded)
}

func formatStatus(count int) string {
    if count == 1 {
        return "1 artiste"
    }
    return fmt.Sprintf("%d artistes", count)
}

func filterArtists(artists []models.Artist, query string) []models.Artist {
    normQ := normalize(query)
    if normQ == "" {
        copied := make([]models.Artist, len(artists))
        copy(copied, artists)
        return copied
    }

    filtered := make([]models.Artist, 0, len(artists))

    for _, artist := range artists {
        if match(normQ,
            artist.Name,
            strings.Join(artist.Members, " "),
            fmt.Sprint(artist.CreationDate),
            artist.FirstAlbum,
        ) {
            filtered = append(filtered, artist)
        }
    }

    return filtered
}

func normalize(s string) string {
    var b strings.Builder
    for _, r := range strings.ToLower(s) {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            b.WriteRune(r)
        }
    }
    return b.String()
}

func match(normQuery string, fields ...string) bool {
    for _, f := range fields {
        if strings.Contains(normalize(f), normQuery) {
            return true
        }
    }
    return false
}
