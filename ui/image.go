package ui

import (
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func AsyncImage(url string, size fyne.Size) *canvas.Image {
	// Icône par défaut pendant le chargement
	img := canvas.NewImageFromResource(theme.FileIcon())
	img.SetMinSize(size)
	img.FillMode = canvas.ImageFillContain

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		// Mise à jour avec la vraie image
		res := fyne.NewStaticResource("cover.jpg", data)
		img.Resource = res
		img.Refresh() // Très important pour que l'image apparaisse !
	}()

	return img
}
