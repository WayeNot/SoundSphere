package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	groups, err := FetchGroups()
	if err != nil {
		log.Fatal("❌ Erreur API GroupieTracker:", err)
	}

	groupMap := make(map[int]Group, len(groups))
	for _, g := range groups {
		groupMap[g.ID] = g
	}

	allConcerts, err := FetchAllConcerts()
	if err != nil {
		log.Println("⚠ Impossible de récupérer les concerts:", err)
		allConcerts = make(map[int][]Concert)
	}

	app := &PageData{
		Groups:       groups,
		GroupByID:    groupMap,
		AllConcerts:  allConcerts,
		AudioDBCache: make(map[string]*AudioDBArtist),
		Settings:     Settings{DarkMode: true},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.DisplayPageHandler("home"))
	http.HandleFunc("/artists", app.DisplayPageHandler("artists"))
	http.HandleFunc("/artist", app.DisplayPageHandler("artist"))

	http.HandleFunc("/concerts", app.DisplayConcertsHandler())

	log.Println("✅ Serveur lancé sur http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}