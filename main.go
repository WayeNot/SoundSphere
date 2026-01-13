package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Initialisation aléatoire pour choisir un artiste aléatoire
	rand.Seed(time.Now().UnixNano())

	// Récupération des groupes depuis l'API GroupieTracker
	groups, err := FetchGroups()
	if err != nil {
		log.Fatal("Erreur API GroupieTracker:", err)
	}

	// Map pour accéder rapidement aux groupes par ID
	groupMap := make(map[int]Group, len(groups))
	for _, g := range groups {
		groupMap[g.ID] = g
	}

	// Structure principale de l'application
	app := &PageData{
		Groups:       groups,
		GroupByID:    groupMap,
		AudioDBCache: make(map[string]*AudioDBArtist),
		Settings:     Settings{DarkMode: true},
	}

	// --- SERVEUR HTTP ---

	// Servir les fichiers statiques
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes principales
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			app.DisplayPageHandler("home")(w, r)
		case "/artist":
			app.DisplayPageHandler("artist")(w, r)
		case "/artists":
			app.DisplayPageHandler("artists")(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("✅ Serveur lancé sur http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
