# 🎧 SoundSphère

> **SoundSphère** est une application web interactive qui centralise les informations essentielles sur des artistes et groupes musicaux :
> biographies, réseaux sociaux, visuels, et **prochains concerts** — le tout en un seul endroit 🌍🎶

Projet réalisé dans le cadre du projet **Groupie Tracker**.

---

## ✨ Fonctionnalités principales

### 🎤 Artistes

* Liste complète des artistes
* Page dédiée pour chaque artiste
* Informations générales :

  * Nom
  * Genre 🎵
  * Pays 🌍
  * Année de création 📅
  * Premier album 💿
* Visuels :

  * Image principale
  * Thumbnail
  * Banner (si disponible)

### 📖 Biographie

* Biographie récupérée depuis **TheAudioDB**
* Support FR 🇫🇷 / EN 🇬🇧
* Affichage optimisé avec bouton *Voir la biographie*

### 🌐 Réseaux sociaux

* Liens cliquables vers :

  * YouTube ▶️
  * Instagram 📸
  * Facebook 📘
  * Twitter 🐦
  * Site officiel 🌐
* Icônes dynamiques (affichées seulement si disponibles)

### 🎟️ Prochains concerts

* Récupération via l’API **Groupie Tracker**
* Liste des dates avec :

  * 📅 Date
  * 📍 Lieu
  * 🏙️ Ville
* Message automatique si aucun concert n’est prévu

### 🎲 Page d’accueil

* Artiste aléatoire mis en avant
* Accès rapide à sa page détaillée

### 🌙 Dark Mode

* Mode sombre activé par défaut
* Structure prête pour un futur toggle utilisateur

---

## 🛠️ Technologies utilisées

* **Go (Golang)** 🐹
* **HTML / CSS**
* **Go Templates**
* **APIs externes** :

  * 🎶 [Groupie Tracker API](https://groupietrackers.herokuapp.com/api)
  * 🎧 [TheAudioDB API](https://www.theaudiodb.com)

---

## 📁 Architecture du projet

```bash
.
├── main.go          # Point d’entrée du serveur
├── pages.go         # Handlers & logique des pages
├── artist.go        # API, structs et fusion des données
├── models.go        # Structs globales
├── static/
│   ├── html/        # Templates HTML
│   ├── css/         # Styles
│   └── js/          # Scripts JS
```

---

## 🚀 Lancer le projet en local

```bash
go run .
```

Puis ouvrir 👉 **[http://localhost:8080](http://localhost:8080)**

---

## 🔮 Fonctionnalités à venir

* 🎧 Lecteur audio intégré (aperçus des musiques)
* 📍 Carte interactive des concerts
* 🔍 Recherche avancée par :

  * Nom
  * Genre
  * Pays
* ❤️ Système de favoris
* 🌙 Toggle Dark / Light mode
* ⚡ Cache optimisé des API

---

## 👥 Équipe

Projet réalisé par :

* Émilien
* Tim
* Aymeric

---

## 📜 Licence

Copyright © 2025 - Tous droits réservés par Émilien, Tim & Aymeric