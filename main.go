package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Group struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Settings struct {
	DarkMode bool
}

type PageData struct {
	Groups        []Group
	GroupByID     map[int]Group
	AudioDBCache  map[string]*AudioDBArtist
	Settings      Settings
}

type PageArtist struct {
	Artist   Group
	AudioDB  *AudioDBArtist
	Settings Settings
}

type AudioDBResponse struct {
	Artists []AudioDBArtist `json:"artists"`
}

type AudioDBArtist struct {
	Name        string `json:"strArtist"`
	BiographyFR string `json:"strBiographyFR"`
	BiographyEN string `json:"strBiographyEN"`
	Genre       string `json:"strGenre"`
	Country     string `json:"strCountry"`
	Thumb       string `json:"strArtistThumb"`
	Banner      string `json:"strArtistBanner"`
}

var (
	httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	templates = template.Must(
		template.ParseGlob("./static/html/*.html"),
	)
)

func FetchGroups() ([]Group, error) {
	const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := httpClient.Get(apiURL)
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

func (app *PageData) GetArtistFromAudioDB(name string) (*AudioDBArtist, error) {
	if artist, ok := app.AudioDBCache[name]; ok {
		return artist, nil
	}

	apiKey := "123"
	escapedName := url.QueryEscape(name)

	endpoint := fmt.Sprintf(
		"https://www.theaudiodb.com/api/v1/json/%s/search.php?s=%s",
		apiKey,
		escapedName,
	)

	resp, err := httpClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AudioDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Artists) == 0 {
		app.AudioDBCache[name] = nil
		return nil, nil
	}

	app.AudioDBCache[name] = &result.Artists[0]
	return &result.Artists[0], nil
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	err := templates.ExecuteTemplate(w, name+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

		var artist Group

		if random {
			artist = app.Groups[rand.Intn(len(app.Groups))]
		} else {
			idStr := r.URL.Query().Get("id")
			if idStr == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			var id int
			fmt.Sscanf(idStr, "%d", &id)

			var ok bool
			artist, ok = app.GroupByID[id]
			if !ok {
				http.NotFound(w, r)
				return
			}
		}

		audioArtist, err := app.GetArtistFromAudioDB(artist.Name)
		if err != nil {
			log.Println("AudioDB error:", err)
		}

		renderTemplate(w, "artist", PageArtist{
			Artist:   artist,
			AudioDB:  audioArtist,
			Settings: app.Settings,
		})
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	groups, err := FetchGroups()
	if err != nil {
		log.Fatal("Erreur API GroupieTracker:", err)
	}

	groupMap := make(map[int]Group, len(groups))
	for _, g := range groups {
		groupMap[g.ID] = g
	}

	app := &PageData{
		Groups:       groups,
		GroupByID:    groupMap,
		AudioDBCache: make(map[string]*AudioDBArtist),
		Settings:     Settings{DarkMode: true},
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static")),
		),
	)

	http.HandleFunc("/", app.DisplayArtist(true))
	http.HandleFunc("/artists", app.DisplayPage("artists"))
	http.HandleFunc("/concerts", app.DisplayPage("concerts"))
	http.HandleFunc("/map", app.DisplayPage("map"))
	http.HandleFunc("/about", app.DisplayPage("about"))
	http.HandleFunc("/artist", app.DisplayArtist(false))

	log.Println("✅ Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
