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

type GroupID struct {
	ID int `json:"id"`
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
	if err := json.Unmarshal(body, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func DefaultSettings() Settings {
	return Settings {
		DarkMode: true,
	}
}

func (app *PageData) DisplayPage(templateName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Groups:   app.Groups,
			Settings: app.Settings,

		}

		tmpl, err := template.ParseFiles("./static/html/" + templateName + ".html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
	}
}

func (app *PageData) DisplayArtist(GetRandom bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if GetRandom {
			idStr := fmt.Sprint(rand.Intn(len(app.Groups)))

			if idStr == "" {
				http.Error(w, "Missing id", http.StatusBadRequest)
				return
			}

			var id int
			fmt.Sscanf(idStr, "%d", &id)

			var found *Group
			for _, g := range app.Groups {
				if g.ID == id {
					found = &g
					break
				}
			}

			if found == nil {
				http.NotFound(w, r)
				return
			}

			data := PageArtist{
				Artist:   *found,
				Settings: app.Settings,
			}

			tmpl, err := template.ParseFiles("./static/html/index.html")

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			tmpl.Execute(w, data)
		} else {
			idStr := r.URL.Query().Get("id")

			if idStr == "" {
				http.Error(w, "Missing id", http.StatusBadRequest)
				return
			}

			var id int
			fmt.Sscanf(idStr, "%d", &id)

			var found *Group
			for _, g := range app.Groups {
				if g.ID == id {
					found = &g
					break
				}
			}

			if found == nil {
				http.NotFound(w, r)
				return
			}

			data := PageArtist{
				Artist:   *found,
				Settings: app.Settings,
			}

			tmpl, err := template.ParseFiles("./static/html/artist.html")

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			tmpl.Execute(w, data)
		}
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
