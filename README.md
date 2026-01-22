# SoundTrap - Groupie Tracker

> Une application graphique moderne pour découvrir et explorer les artistes musicaux, leurs membres, leurs albums et leurs tournées mondiales.

## 📋 Table des matières

- [À propos](#à-propos)
- [Fonctionnalités](#fonctionnalités)
- [Architecture](#architecture)
- [Installation](#installation)
- [Utilisation](#utilisation)
- [Structure du projet](#structure-du-projet)
- [Technologies](#technologies)
- [Contributeurs](#contributeurs)

## 🎵 À propos

**SoundTrap** est une application desktop développée en Go qui consomme l'API Groupie Trackers pour afficher une galerie interactive d'artistes musicaux. L'application offre une interface utilisateur intuitive et responsable avec deux niveaux de consultation : une galerie d'accueil avec recherche et une vue détaillée par artiste.

## ✨ Fonctionnalités

### Écran Principal - Galerie
- 📸 **Galerie d'artistes** : Affichage en grille (3 colonnes) des artistes avec leurs photos
- 🔍 **Recherche en temps réel** : Filtrez les artistes par nom, membre, date, ou concert
- 📅 **Métadonnées** : Année de création affichée sur chaque carte
- 🖱️ **Navigation fluide** : Cliquez sur une carte pour voir les détails de l'artiste

### Écran Détails - Artiste
- 🎤 **Vue complète** : Photo haute résolution et informations détaillées
- 👥 **Liste des membres** : Tous les membres du groupe listés
- 🗺️ **Concerts & Lieux** : Affichage interactif des dates et lieux de concert
- 📍 **Intégration OpenStreetMap** : Accès direct aux coordonnées GPS des lieux de concert
- ⬅️ **Navigation** : Retour facile à la galerie principale

## 🏗️ Architecture

```
groupie-tracker-gui/
├── api/           # Couche d'accès aux données
│   ├── client.go  # Client API Groupie Trackers
│   └── geoloc.go  # Service de géolocalisation
├── models/        # Structures de données
│   └── artist.go  # Modèles Artist et Relation
├── ui/            # Interface utilisateur Fyne
│   ├── handler.go # Logique de rendu des écrans
│   └── image.go   # Gestion asynchrone des images
├── main.go        # Point d'entrée de l'application
├── go.mod         # Dépendances Go
└── README.md      # Cette documentation
```

## 🚀 Installation

### Prérequis

- **Go** 1.25 ou supérieur
- **Windows**, **macOS** ou **Linux**
- Connexion Internet (pour charger l'API Groupie Trackers)

### Étapes

1. **Cloner le dépôt**
   ```bash
   git clone https://github.com/votre-repo/groupie-tracker-gui.git
   cd groupie-tracker-gui
   ```

2. **Installer les dépendances**
   ```bash
   go mod download
   ```

3. **Compiler l'application**
   ```bash
   go build -o SoundTrap.exe
   ```

4. **Lancer l'application**
   ```bash
   ./SoundTrap.exe
   ```

   Ou en mode développement :
   ```bash
   go run main.go
   ```

## 💻 Utilisation

### Galerie Principale
1. L'application se lance sur la galerie complète des 52 artistes
2. Utilisez la **barre de recherche** pour filtrer :
   - Par **nom d'artiste** : "Queen", "Pink Floyd"
   - Par **membre** : "Freddie Mercury"
   - Par **année** : "1970"
   - Par **concert/lieu** : "Paris", "New York"

### Vue Détails
1. Cliquez sur une carte d'artiste pour ouvrir sa vue détaillée
2. Consultez l'image, les informations et la liste des membres
3. Pour chaque lieu de concert, cliquez sur le bouton **📍 Lieu** pour ouvrir OpenStreetMap
4. Cliquez sur **Retour à la liste** pour revenir à la galerie

## 📦 Structure du projet

### `api/`
- **client.go** : Requêtes HTTP vers l'API Groupie Trackers
- **geoloc.go** : Conversion adresses/coordinates (Nominatim)

### `models/`
- **artist.go** : Structures `Artist` et `Relation` (sérialisation JSON)

### `ui/`
- **handler.go** : Création des écrans (liste & détails)
- **image.go** : Chargement asynchrone des images avec cache

### `main.go`
- Initialisation de l'application Fyne
- Chargement des données au démarrage
- Gestion de la fenêtre principale

## 🛠️ Technologies

| Composant | Version | Usage |
|-----------|---------|-------|
| **Go** | 1.25 | Langage principal |
| **Fyne** | v2.7.2 | Framework GUI |
| **API** | Groupie Trackers | Données artistes |
| **Géolocalisation** | OpenStreetMap/Nominatim | Coordonnées |

## 🔄 Flux de données

```
API Groupie Trackers
        ↓
   api/client.go
        ↓
  models/artist.go
        ↓
   ui/handler.go (rendu)
        ↓
   ui/image.go (chargement images)
        ↓
   Application Fyne
```

## 🎨 Interface Utilisateur

- **Thème** : Dark mode optimisé pour la musique
- **Couleurs** : Palette sombre (gris/bleu) pour une ambiance musicale
- **Police** : UI système par défaut (accessible)
- **Responsive** : Redimensionnable et adaptable

## ⚙️ Configuration

Aucune configuration requise. L'application fonctionne directement avec :
- URL API : `https://groupietrackers.herokuapp.com/api`
- Cache images : Mémoire de l'application
- Fenêtre par défaut : 1200x800 pixels

## 🐛 Dépannage

### Les images ne s'affichent pas
- Vérifiez votre connexion Internet
- Attendez quelques secondes (chargement asynchrone)
- Les images sont en cache après le premier chargement

### L'application freeze au démarrage
- C'est normal, l'API se charge en arrière-plan
- Attendez 3-5 secondes pour que les données arrivent

### Erreur "connection refused"
- Vérifiez que l'API Groupie Trackers est accessible
- Vérifiez votre connexion Internet

## 📝 Licence

Ce projet est fourni à titre éducatif.

## 👥 Contributeurs

- **Ersanios** - Contributeur
- **Simon** - Contributeur
- **Victor** - Contributeur

---

**Dernière mise à jour** : Janvier 2026  
**Version** : 1.0.0