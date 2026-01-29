```
# 🎵 SoundSphère — GroupieTracker

SoundSphère est une plateforme web qui centralise artistes et concerts du monde entier. Ce projet permet aux utilisateurs de découvrir des artistes, de consulter leurs concerts passés et à venir, et de filtrer facilement les événements selon la ville ou le nom de l’artiste.

---

## 🛠️ Fonctionnalités

- Liste complète des artistes avec informations détaillées (image, biographie, réseaux sociaux).  
- Page concerts avec filtres interactifs :  
  - Recherche par artiste ou ville  
  - Tri automatique par date  
- Interface moderne et responsive, dans un style sombre et immersif.  
- Pages entièrement en **Go (Golang)** côté serveur avec **templates HTML/CSS**, sans JavaScript.  
- Navbar fixe et stylée pour une navigation fluide.  

---

## 📁 Structure du projet

```

Projet_GroupieTracker/
├─ main.go              # Point d'entrée du serveur
├─ concerts.go          # Gestion et affichage des concerts
├─ artist.go            # Gestion et affichage des artistes
├─ pages.go             # Handlers des pages principales
├─ struct.go            # Définition des structures (Concert, Artist, etc.)
├─ function.go          # Fonctions utilitaires
├─ static/
│  ├─ css/
│  │  ├─ index.css
│  │  ├─ concerts.css
│  │  └─ ... autres CSS
│  └─ html/
│     ├─ index.html
│     ├─ concerts.html
│     ├─ artists.html
│     └─ artist.html
└─ go.mod

````

---

## 🚀 Lancer le projet

1. Cloner le dépôt :  
```bash
git clone [https://github.com/Lodgia/Projet_GroupieTracker.git]
cd Projet_GroupieTracker
````

2. Installer les dépendances (si nécessaire) :

```bash
go mod tidy
```

3. Lancer le serveur :

```bash
go run .
```

4. Ouvrir le navigateur et accéder à :

```
http://localhost:8081
```

---

## 🎨 Style et UI

* Thème sombre avec dégradés et effets modernes
* Navbar fixe, responsive et élégante
* Cartes pour artistes et concerts avec hover effects
* Design responsive pour tous les écrans (desktop, tablette, mobile)

---

## 📌 Technologies utilisées

* [Golang](https://golang.org/) pour le backend
* Templates HTML côté serveur
* CSS pur pour le style et les effets modernes
* Git & GitHub pour le versioning

---

## 👥 Équipe

Projet réalisé par :

* Émilien
* Tim
* Aymeric

---

## 📜 Licence

Copyright © 2025 - Tous droits réservés par Émilien, Tim & Aymeric
