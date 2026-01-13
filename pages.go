package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseGlob("./static/html/*.html"))

// ---------------- TEMPLATE ----------------

func renderTemplate(w http.ResponseWriter, name string, data any) {
	if err := templates.ExecuteTemplate(w, name+".html", data); err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// ---------------- AUDIO DB ----------------

func (app *PageData) GetAudioDBArtist(name string) (*AudioDBArtist, error) {
	if artist, ok := app.AudioDBCache[name]; ok {
		return artist, nil
	}
	audio, err := FetchAudioDB(name, "123") // Clé API AudioDB
	if err != nil {
		return nil, err
	}
	app.AudioDBCache[name] = audio
	return audio, nil
}

// ---------------- HANDLER ----------------

func (app *PageData) DisplayPageHandler(pageType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allConcerts, _ := FetchAllConcerts() // Récupération des concerts

		switch pageType {

		// ---------------- HOME ----------------
		case "home":
			if len(app.Groups) == 0 {
				http.Error(w, "Aucun artiste disponible", http.StatusInternalServerError)
				return
			}

			group := app.Groups[rand.Intn(len(app.Groups))] // Artiste aléatoire
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

		// ---------------- ARTIST ----------------
		case "artist":
			idStr := r.URL.Query().Get("id")
			if idStr == "" {
				http.Error(w, "ID manquant", http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "ID invalide", http.StatusBadRequest)
				return
			}

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

		// ---------------- ARTISTS ----------------
		case "artists":
			// Paramètres GET
			filter := r.URL.Query().Get("filterArtist")
			search := r.URL.Query().Get("search")

			perPage := 10
			if v := r.URL.Query().Get("perPage"); v != "" {
				if n, err := strconv.Atoi(v); err == nil && n > 0 {
					perPage = n
				}
			}

			page := 1
			if v := r.URL.Query().Get("page"); v != "" {
				if n, err := strconv.Atoi(v); err == nil && n > 0 {
					page = n
				}
			}

			// Merge et filtre
			allArtists := make([]ArtistFull, 0, len(app.Groups))
			for i, g := range app.Groups {
				if search != "" && !strings.Contains(strings.ToLower(g.Name), strings.ToLower(search)) {
					continue
				}
				audio, _ := app.GetAudioDBArtist(g.Name)
				allArtists = append(allArtists, MergeArtistData(g, audio, allConcerts, i))
			}

			// Trier
			switch filter {
			case "alphaAZ":
				SortArtistsAZ(allArtists)
			case "alphaZA":
				SortArtistsZA(allArtists)
			case "plusVieuxMoinsVieux":
				SortArtistsOldToNew(allArtists)
			case "moinsVieuxPlusVieux":
				SortArtistsNewToOld(allArtists)
			}

			// Pagination
			total := len(allArtists)
			totalPages := (total + perPage - 1) / perPage
			if totalPages == 0 {
				totalPages = 1
			}

			if page < 1 {
				page = 1
			} else if page > totalPages {
				page = totalPages
			}

			start := (page - 1) * perPage
			end := start + perPage
			if start > total {
				start = 0
			}
			if end > total {
				end = total
			}
			pagedArtists := allArtists[start:end]

			pageNumbers := make([]int, totalPages)
			for i := 0; i < totalPages; i++ {
				pageNumbers[i] = i + 1
			}

			prevPage := 0
			if page > 1 {
				prevPage = page - 1
			}
			nextPage := 0
			if page < totalPages {
				nextPage = page + 1
			}

			renderTemplate(w, "artists", ListPage{
				Groups:      pagedArtists,
				Settings:    app.Settings,
				Filter:      filter,
				Search:      search,
				CurrentPage: page,
				PerPage:     perPage,
				TotalPages:  totalPages,
				PageNumbers: pageNumbers,
				PrevPage:    prevPage,
				NextPage:    nextPage,
			})

		default:
			http.Error(w, "Page non implémentée", http.StatusNotImplemented)
		}
	}
}