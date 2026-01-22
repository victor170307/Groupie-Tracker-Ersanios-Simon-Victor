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
	img.ScaleMode = canvas.ImageScaleSmooth

	go func() {
		if url == "" {
			log.Println("URL vide")
			return
		}

		log.Println("Chargement image:", url)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Erreur création requête:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Erreur téléchargement image:", url, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			log.Println("Status code:", resp.StatusCode)
			return
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Erreur lecture données image:", err)
			return
		}

		log.Println("Image chargée, taille:", len(data), "bytes")
		res := fyne.NewStaticResource("image.jpg", data)
		img.Resource = res
		img.Refresh()
		log.Println("Image refreshée")
	}()

	return img
}
