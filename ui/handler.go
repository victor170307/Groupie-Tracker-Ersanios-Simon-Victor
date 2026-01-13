package ui

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker-gui/models"
)


func ServeHTML(w http.ResponseWriter, artists []models.Artist) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := getHTMLTemplate(artists)
	fmt.Fprint(w, html)
}

// ServeArtistsJSON returns filtered artists as JSON
func ServeArtistsJSON(w http.ResponseWriter, artists []models.Artist) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

func getHTMLTemplate(artists []models.Artist) string {
	artistsJSON, _ := json.Marshal(artists)

	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Groupie Tracker</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); min-height: 100vh; padding: 20px; }
        .container { max-width: 1200px; margin: 0 auto; }
        h1 { color: white; text-align: center; margin-bottom: 30px; font-size: 2.5em; text-shadow: 0 2px 10px rgba(0,0,0,0.3); }
        .search-box { margin-bottom: 20px; }
        #search { width: 100%; padding: 15px; font-size: 16px; border: none; border-radius: 8px; box-shadow: 0 4px 15px rgba(0,0,0,0.2); }
        .artists-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(250px, 1fr)); gap: 20px; }
        .artist-card { background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 4px 15px rgba(0,0,0,0.15); transition: transform 0.3s, box-shadow 0.3s; cursor: pointer; }
        .artist-card:hover { transform: translateY(-5px); box-shadow: 0 8px 25px rgba(0,0,0,0.25); }
        .artist-image { width: 100%; height: 200px; object-fit: cover; background: #f0f0f0; }
        .artist-info { padding: 15px; }
        .artist-name { font-size: 1.3em; font-weight: bold; color: #333; margin-bottom: 8px; }
        .artist-members { font-size: 0.9em; color: #666; margin-bottom: 5px; }
        .artist-date { font-size: 0.85em; color: #999; }
        .no-results { text-align: center; color: white; font-size: 1.2em; padding: 40px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎸 Groupie Tracker</h1>
        <div class="search-box">
            <input type="text" id="search" placeholder="Search artist, member, date, location...">
        </div>
        <div class="artists-grid" id="artists-container"></div>
        <div id="no-results" class="no-results" style="display: none;">No artists found</div>
    </div>

    <script>
        const allArtists = ` + string(artistsJSON) + `;
        const container = document.getElementById('artists-container');
        const noResults = document.getElementById('no-results');
        const searchInput = document.getElementById('search');

        function render(artists) {
            container.innerHTML = '';
            if (artists.length === 0) {
                noResults.style.display = 'block';
                return;
            }
            noResults.style.display = 'none';
            artists.forEach(artist => {
                const card = document.createElement('div');
                card.className = 'artist-card';
                card.innerHTML = ` + "`" + `
                    <img src="${artist.image}" alt="${artist.name}" class="artist-image" onerror="this.src='https://via.placeholder.com/250x200?text=No+Image'">
                    <div class="artist-info">
                        <div class="artist-name">${artist.name}</div>
                        <div class="artist-members"><strong>Members:</strong> ${artist.members.join(', ')}</div>
                        <div class="artist-date"><strong>Created:</strong> ${artist.creationDate}</div>
                    </div>
                ` + "`" + `;
                container.appendChild(card);
            });
        }

        function filter(query) {
            const q = query.toLowerCase();
            return allArtists.filter(artist => {
                if (artist.name.toLowerCase().includes(q)) return true;
                if (artist.members.some(m => m.toLowerCase().includes(q))) return true;
                if (artist.creationDate.toString().includes(q)) return true;
                return false;
            });
        }

        searchInput.addEventListener('input', (e) => {
            const filtered = filter(e.target.value);
            render(filtered);
        });

        // Initial render
        render(allArtists);
    </script>
</body>
</html>`
}
