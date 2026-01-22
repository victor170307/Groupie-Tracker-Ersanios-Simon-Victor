package ui

import (
	"io"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

func AsyncImage(url string, size fyne.Size) *canvas.Image {
	img := canvas.NewImageFromResource(theme.FileIcon())
	img.SetMinSize(size)
	img.FillMode = canvas.ImageFillContain

	go func() {
		if url == "" {
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Erreur création requête:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Erreur téléchargement image:", url, err)
			img.Resource = theme.ErrorIcon()
			img.Refresh()
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Erreur lecture données image:", err)
			return
		}

		res := fyne.NewStaticResource("image.jpg", data)
		img.Resource = res
		img.Refresh()
	}()

	return img
}
