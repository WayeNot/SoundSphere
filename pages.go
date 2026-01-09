package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var templates = template.Must(template.ParseGlob("./static/html/*.html"))

type ListPage struct {
	Groups   []ArtistFull
	Settings Settings
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	if err := templates.ExecuteTemplate(w, name+".html", data); err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (app *PageData) GetAudioDBArtist(name string) (*AudioDBArtist, error) {
	if artist, ok := app.AudioDBCache[name]; ok {
		return artist, nil
	}
	audio, err := FetchAudioDB(name, "123")
	if err != nil {
		return nil, err
	}
	app.AudioDBCache[name] = audio
	return audio, nil
}

func (app *PageData) DisplayPageHandler(pageType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		allConcerts, _ := FetchAllConcerts()

		switch pageType {

		case "home":
			group := app.Groups[rand.Intn(len(app.Groups))]

			audio, _ := app.GetAudioDBArtist(group.Name)

			index := -1
			for i, g := range app.Groups {
				if g.ID == group.ID {
					index = i
					break
				}
			}

			artist := MergeArtistData(group, audio, allConcerts, index)

			renderTemplate(w, "index", ArtistPage{
				Artist:   artist,
				Settings: app.Settings,
			})

		case "artist":
			idStr := r.URL.Query().Get("id")
			if idStr == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			var id int
			fmt.Sscanf(idStr, "%d", &id)
			group, ok := app.GroupByID[id]
			if !ok {
				http.NotFound(w, r)
				return
			}

			audio, _ := app.GetAudioDBArtist(group.Name)

			index := -1
			for i, g := range app.Groups {
				if g.ID == id {
					index = i
					break
				}
			}

			artist := MergeArtistData(group, audio, allConcerts, index)

			renderTemplate(w, "artist", ArtistPage{
				Artist:   artist,
				Settings: app.Settings,
			})

		case "artists":
			artists := make([]ArtistFull, len(app.Groups))
			for i, g := range app.Groups {
				audio, _ := app.GetAudioDBArtist(g.Name)
				artists[i] = MergeArtistData(g, audio, allConcerts, i)
			}

			renderTemplate(w, "artists", ListPage{
				Groups:   artists,
				Settings: app.Settings,
			})

		default:
			http.Error(w, "Page not implemented", http.StatusNotImplemented)
		}
	}
}
