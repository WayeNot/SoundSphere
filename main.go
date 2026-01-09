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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/", app.DisplayPageHandler("home"))
	http.HandleFunc("/artist", app.DisplayPageHandler("artist"))
	http.HandleFunc("/artists", app.DisplayPageHandler("artists"))

	log.Println("✅ Serveur lancé sur http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}