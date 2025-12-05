package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type PageData struct {
	Groups   []Group
	Settings Settings
}

type PageArtist struct {
	Artist   Group
	Settings Settings
}

type Settings struct {
	DarkMode bool
}

type Group struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func ManageApi() ([]Group, error) {
	const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var groups []Group
	err = json.Unmarshal(body, &groups)
	return groups, err
}

func DefaultSettings() Settings {
	return Settings{DarkMode: true}
}

func (app *PageData) findArtistByID(id int) *Group {
	for _, g := range app.Groups {
		if g.ID == id {
			return &g
		}
	}
	return nil
}

func renderTemplate(w http.ResponseWriter, file string, data any) {
	tmpl, err := template.ParseFiles("./static/html/" + file + ".html")
	if err != nil {
		http.Error(w, "Template error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func (app *PageData) DisplayPage(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, name, PageData{
			Groups:   app.Groups,
			Settings: app.Settings,
		})
	}
}

func (app *PageData) DisplayArtist(random bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if random {
			randomIndex := rand.Intn(len(app.Groups))
			artist := app.Groups[randomIndex]

			renderTemplate(w, "index", PageArtist{
				Artist:   artist,
				Settings: app.Settings,
			})
			return
		}

		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, "Missing ID", http.StatusBadRequest)
			return
		}

		var id int
		fmt.Sscanf(idStr, "%d", &id)

		artist := app.findArtistByID(id)
		if artist == nil {
			http.NotFound(w, r)
			return
		}

		renderTemplate(w, "artist", PageArtist{
			Artist:   *artist,
			Settings: app.Settings,
		})
	}
}

func main() {
	groups, err := ManageApi()
	if err != nil {
		log.Fatal("Erreur API :", err)
	}

	app := &PageData{
		Groups:   groups,
		Settings: DefaultSettings(),
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", app.DisplayArtist(true))
	http.HandleFunc("/artists", app.DisplayPage("artists"))
	http.HandleFunc("/concerts", app.DisplayPage("concerts"))
	http.HandleFunc("/map", app.DisplayPage("map"))
	http.HandleFunc("/about", app.DisplayPage("about"))
	http.HandleFunc("/artist", app.DisplayArtist(false))

	log.Println("Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}