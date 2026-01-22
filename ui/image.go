package ui

import (
	"io"
	"log"
	"net/http"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var imageCache = make(map[string]*fyne.StaticResource)
var cacheMutex sync.Mutex

func AsyncImage(url string, size fyne.Size) *canvas.Image {
<<<<<<< HEAD
	// 1. Placeholder : On met une icône "File" ou "Question" par défaut
	// Utiliser NewImageFromResource(nil) crée un placeholder vide mais propre
	img := canvas.NewImageFromResource(theme.FileIcon())
=======
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
>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562
	img.SetMinSize(size)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScaleSmooth

	// 2. Lancement du téléchargement
	go func() {
		if url == "" {
			return
		}

		// Ajout d'un User-Agent (parfois les API bloquent les requêtes sans agent)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("Erreur création requête:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Erreur téléchargement image:", url, err)
			// Optionnel : Mettre une icône d'erreur ici
			img.Resource = theme.ErrorIcon()
			img.Refresh()
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Erreur lecture données image:", err)
			return
		}

<<<<<<< HEAD
		// 3. Création de la ressource statique (JPEG/PNG)
		// Le nom "image.jpg" aide Fyne à comprendre qu'il doit décoder du JPEG
		res := fyne.NewStaticResource("image.jpg", data)

		// 4. Mise à jour thread-safe (méthode simple)
		img.Resource = res

		// Force le redessin de l'objet image
=======
		res := fyne.NewStaticResource("img.jpg", data)
		
		cacheMutex.Lock()
		imageCache[url] = res
		cacheMutex.Unlock()
		
		img.Resource = res
>>>>>>> dbd67431a07f52576798c58a1128dcfdb49ad562
		img.Refresh()
	}()

	return img
}
