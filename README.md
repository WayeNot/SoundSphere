---

 🎧 SoundSphère — GroupieTracker

**SoundSphère** est une plateforme web immersive dédiée à la découverte musicale.
Elle centralise artistes et concerts du monde entier en proposant une expérience riche, moderne et entièrement rendue côté serveur en **Go (Golang)**.

Le projet se distingue par le **croisement de plusieurs APIs musicales** afin d’offrir une quantité d’informations détaillées et cohérentes pour chaque artiste.

🔗 **Site en ligne**
👉 [https://soundsphere-0yjv.onrender.com/](https://soundsphere-0yjv.onrender.com/)

---

## 🌍 Présentation du projet

SoundSphère permet aux utilisateurs de :

* découvrir des artistes internationaux,
* explorer leurs concerts passés et à venir,
* accéder à des informations musicales approfondies,
* naviguer dans une interface sombre, élégante et immersive.

Le projet adopte volontairement une approche **backend-driven** :

* aucun framework frontend,
* aucun JavaScript côté client,
* tout le rendu est assuré par des **templates HTML en Go**.

---

## 🔗 Croisement d’APIs & enrichissement des données

SoundSphère repose sur la **fusion intelligente de deux APIs complémentaires**, afin de proposer des pages artistes très complètes.

### 🎵 GroupieTracker API

Utilisée comme base principale pour :

* l’identité des artistes / groupes
* les membres
* l’année de création
* le premier album
* les concerts (dates, lieux, villes)

### 🎧 TheAudioDB API

Utilisée pour enrichir les données musicales avec :

* biographies détaillées (**FR & EN**)
* genre musical
* pays d’origine
* visuels avancés (thumbnail, bannière)
* liens vers les réseaux sociaux et plateformes musicales

👉 Les données sont **fusionnées côté serveur en Go**, pour créer une structure unifiée représentant un artiste complet.

---

## 📊 Données disponibles par artiste

Grâce à ce croisement, chaque artiste peut contenir :

### 🎤 Identité & musique

* Nom de l’artiste
* Image principale
* Membres du groupe
* Année de création
* Premier album
* Genre musical
* Pays d’origine

### 📝 Contenu éditorial

* Biographie en français
* Biographie en anglais

### 🖼️ Visuels

* Thumbnail officiel
* Bannière artistique

### 🌐 Réseaux & plateformes

* Facebook
* Twitter / X
* Instagram
* Site officiel
* YouTube
* Last.fm
* MusicBrainz

### 🎫 Concerts

* Liste complète des concerts
* Dates
* Lieux
* Villes

---

## ✨ Fonctionnalités actuelles

### 🎤 Artistes

* Liste complète des artistes
* Page dédiée par artiste avec informations détaillées
* Données enrichies via le croisement des APIs

### 🎫 Concerts

* Page concerts centralisée
* Recherche par :

  * nom d’artiste
  * ville
* Tri automatique par date
* Gestion des données manquantes

### 🎨 Interface & UX

* Thème sombre immersif inspiré des plateformes musicales
* UI moderne, lisible et cohérente
* Navigation fluide
* Navbar persistante
* Mise en page responsive (desktop / tablette / mobile)

### ⚙️ Architecture

* Backend **100 % Golang**
* Templates HTML rendus côté serveur
* CSS pur (sans framework)
* Aucune dépendance JavaScript
* Cache interne pour les données AudioDB
* Code structuré et maintenable

---

## 🚀 Fonctionnalités à venir (Roadmap)

### 🎧 Expérience musicale

* ▶️ Écoute d’un extrait audio directement depuis la page artiste
* 🔗 Lien direct vers la page Spotify officielle de l’artiste
* Intégration d’un lecteur audio léger

### 📱 UI & UX

* Finalisation complète du responsive
* Animations et transitions plus avancées
* Amélioration de l’accessibilité

### 🔍 Navigation & recherche

* Recherche globale artistes + concerts
* Suggestions dynamiques
* Pagination intelligente

### ⭐ Personnalisation (long terme)

* Système de favoris
* Recommandations d’artistes similaires (genre / pays)
* Historique de navigation
* Mode clair / sombre configurable

---

## 📁 Structure du projet

```
Projet_GroupieTracker/
├─ main.go              # Point d’entrée du serveur
├─ pages.go             # Handlers des pages principales
├─ artist.go            # Gestion des artistes
├─ concerts.go          # Gestion des concerts
├─ struct.go            # Structures de données
├─ function.go          # Fonctions utilitaires
├─ static/
│  ├─ css/
│  │  ├─ index.css
│  │  ├─ artists.css
│  │  ├─ concerts.css
│  │  └─ ...
│  └─ html/
│     ├─ home.html
│     ├─ artists.html
│     ├─ artist.html
│     └─ concerts.html
├─ go.mod
└─ go.sum
```

---

## ▶️ Lancer le projet en local

### 1️⃣ Cloner le dépôt

```bash
git clone https://github.com/WayeNot/SoundSphere.git
cd Projet_GroupieTracker
```

### 2️⃣ Installer les dépendances

```bash
go mod tidy
```

### 3️⃣ Lancer le serveur

```bash
go run .
```

### 4️⃣ Accéder au site

```
http://localhost:8081
```

---

## 🛠️ Technologies utilisées

* **Golang (Go)** — Backend & serveur HTTP
* **HTML Templates** — Rendu côté serveur
* **CSS pur** — UI et animations
* **APIs REST** — GroupieTracker & TheAudioDB
* **Git & GitHub** — Versioning
* **Render** — Déploiement

---

## 👥 Équipe

Projet réalisé par :

* **Aymeric**
* Émilien
* Tim

---

## 📜 Licence

© 2025 — SoundSphère
Projet réalisé dans un cadre pédagogique.
Tous droits réservés.

---