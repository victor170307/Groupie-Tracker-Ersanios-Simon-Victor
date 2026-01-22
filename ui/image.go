package ui

import (
	"io"
	"net/http"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var imageCache = make(map[string]*fyne.StaticResource)
var cacheMutex sync.Mutex

func AsyncImage(url string, size fyne.Size) *canvas.Image {
	cacheMutex.Lock()
	cached, exists := imageCache[url]
	cacheMutex.Unlock()
	
	if exists {
		img := canvas.NewImageFromResource(cached)
		img.SetMinSize(size)
		img.FillMode = canvas.ImageFillContain
		img.ScaleMode = canvas.ImageScaleSmooth
		return img
	}
	
	img := canvas.NewImageFromResource(theme.FileImageIcon())
	img.SetMinSize(size)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScaleSmooth

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		res := fyne.NewStaticResource("img.jpg", data)
		
		cacheMutex.Lock()
		imageCache[url] = res
		cacheMutex.Unlock()
		
		img.Resource = res
		img.Refresh()
	}()

	return img
}
