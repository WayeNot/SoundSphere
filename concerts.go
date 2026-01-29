package main

import (
	"net/http"
	"strings"
)

type ConcertDisplay struct {
	Artist string
	Date   string
	City   string
}

type ConcertPage struct {
	Concerts         []ConcertDisplay
	AvailableArtists []string
	AvailableCities  []string
	FilterArtist     string
	FilterCity       string
	Search           string
}

func (app *PageData) DisplayConcertsHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        search := strings.ToLower(r.URL.Query().Get("search"))
        filterCity := r.URL.Query().Get("city")
        filterArtist := r.URL.Query().Get("artist")

        concerts := []ConcertDisplay{}
        artistSet := map[string]struct{}{}
        citySet := map[string]struct{}{}

        for _, g := range app.Groups {
            groupConcerts, ok := app.AllConcerts[g.ID]
            if !ok {
                continue
            }
            for _, c := range groupConcerts {
                if filterArtist != "" && g.Name != filterArtist {
                    continue
                }
                if filterCity != "" && c.City != filterCity {
                    continue
                }
                if search != "" &&
                    !strings.Contains(strings.ToLower(g.Name), search) &&
                    !strings.Contains(strings.ToLower(c.City), search) {
                    continue
                }
                concerts = append(concerts, ConcertDisplay{
                    Artist: g.Name,
                    Date:   c.Date,
                    City:   c.City,
                })
                artistSet[g.Name] = struct{}{}
                citySet[c.City] = struct{}{}
            }
        }

        availableArtists := []string{}
        for a := range artistSet {
            availableArtists = append(availableArtists, a)
        }
        availableCities := []string{}
        for c := range citySet {
            availableCities = append(availableCities, c)
        }

        page := ConcertPage{
            Concerts:         concerts,
            AvailableArtists: availableArtists,
            AvailableCities:  availableCities,
            FilterArtist:     filterArtist,
            FilterCity:       filterCity,
            Search:           search,
        }

        renderTemplate(w, "concerts", page)
    }
}
