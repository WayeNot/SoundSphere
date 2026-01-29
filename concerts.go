package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type ConcertDisplay struct {
	Date   string
	Artist string
	City   string
}

type ConcertPage struct {
	Concerts        []ConcertDisplay
	Settings        Settings
	FilterCity      string
	FilterArtist    string
	Search          string
	AvailableCities  []string
	AvailableArtists []string
}

func (app *PageData) DisplayConcertsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filterCity := r.URL.Query().Get("city")
		filterArtist := r.URL.Query().Get("artist")
		search := strings.TrimSpace(r.URL.Query().Get("search"))

		var concerts []ConcertDisplay
		citySet := make(map[string]struct{})
		artistSet := make(map[string]struct{})

		for _, g := range app.Groups {
			artistName := g.Name
			if _, ok := artistSet[artistName]; !ok {
				artistSet[artistName] = struct{}{}
			}

			for _, c := range app.AllConcerts[g.ID] {
				city := c.City
				if _, ok := citySet[city]; !ok {
					citySet[city] = struct{}{}
				}

				// Filtres
				if filterCity != "" && !strings.EqualFold(filterCity, city) {
					continue
				}
				if filterArtist != "" && !strings.EqualFold(filterArtist, artistName) {
					continue
				}
				if search != "" && !strings.Contains(strings.ToLower(artistName+" "+city), strings.ToLower(search)) {
					continue
				}

				concerts = append(concerts, ConcertDisplay{
					Date:   c.Date,
					City:   city,
					Artist: artistName,
				})
			}
		}

		sort.Slice(concerts, func(i, j int) bool {
			t1, err1 := time.Parse("02-01-2006", concerts[i].Date)
			t2, err2 := time.Parse("02-01-2006", concerts[j].Date)
			if err1 != nil || err2 != nil {
				return concerts[i].Date > concerts[j].Date
			}
			return t2.Before(t1)
		})

		var cities, artists []string
		for c := range citySet {
			cities = append(cities, c)
		}
		for a := range artistSet {
			artists = append(artists, a)
		}
		sort.Strings(cities)
		sort.Strings(artists)

		page := ConcertPage{
			Concerts:        concerts,
			Settings:        app.Settings,
			FilterCity:      filterCity,
			FilterArtist:    filterArtist,
			Search:          search,
			AvailableCities:  cities,
			AvailableArtists: artists,
		}

		tmpl := template.Must(template.ParseFiles("./static/html/concerts.html"))
		if err := tmpl.Execute(w, page); err != nil {
			log.Println("Template error:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}