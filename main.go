package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("✅ Serveur lancé sur le port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}