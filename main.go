package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Récupération des artistes depuis GroupieTracker
	groups, err := FetchGroups()
	if err != nil {
		log.Fatal("Erreur API GroupieTracker:", err)
	}

	// Map pour accès rapide par ID
	groupMap := make(map[int]Group, len(groups))
	for _, g := range groups {
		groupMap[g.ID] = g
	}

	// Initialisation des données de l'application
	app := &PageData{
		Groups:       groups,
		GroupByID:    groupMap,
		AudioDBCache: make(map[string]*AudioDBArtist),
		Settings:     Settings{DarkMode: true},
	}

	// Fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Routes
	http.HandleFunc("/", app.DisplayPageHandler("home"))
	http.HandleFunc("/artist", app.DisplayPageHandler("artist"))
	http.HandleFunc("/artists", app.DisplayPageHandler("artists"))

	log.Println("✅ Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}