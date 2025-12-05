---

# 🎵 SoundSphère — Groupie Tracker

SoundSphère est un projet web développé en Go permettant d’explorer facilement des artistes, leurs concerts, leurs lieux et leurs dates à partir de l’API Groupie Tracker.
Le site propose une interface moderne, dynamique et agréable, avec un design centré autour de la découverte musicale.

---

## 🚀 Fonctionnalités

### 🎲 Artiste aléatoire

* Un artiste différent s’affiche à chaque chargement de la page.
* Bouton **“Tirer un nouvel artiste”** (rechargement actuellement via une requête Go côté serveur).

### 👤 Page artiste

* Affichage détaillé d’un artiste :

  * Nom, image, membres
  * Dates de concerts
  * Localisations
  * Informations de l’API Groupie Tracker

### 📅 Liste des artistes

* Page listant tous les artistes récupérés depuis l’API.

### 🎤 Page concerts

* Liste des concerts, dates et lieux associés.

### 💄 Interface moderne

* Hero section avec image d’arrière-plan assombrie
* Dégradés, text-shadow, boutons cohérents
* Polices modernes (Orbitron, Exo, etc.)

---

## 🛠️ Technologies utilisées

* **Go (Golang)** — backend, routes et logique serveur
* **HTML / CSS** — rendu côté client
* **API Groupie Tracker** — récupération des données
* **net/http** — serveur web en Go

---

## 📦 Installation

### 1. Cloner le projet

```bash
git clone https://github.com/Lodgia/Projet_GroupieTracker.git
cd soundsphere
```

### 2. Lancer le serveur Go

```bash
go run main.go
```

### 3. Ouvrir dans le navigateur

Accéder au site via :
➡ [http://localhost:8080](http://localhost:8080)

---

## 📁 Structure du projet

```
soundsphere/
│
├── static/               # HTML / CSS / images
│   ├── html/            
│   ├── css/              
│   └── img/              
├── main.go               # Point d'entrée du serveur
├── api/                  # Récupération & parsing des données Groupie Tracker
└── README.md             # Documentation
```

---

## 🔌 API utilisée

Le projet s’appuie sur :
➡ **[https://groupietrackers.herokuapp.com/api](https://groupietrackers.herokuapp.com/api)**

---

## 🧠 Idées d’améliorations futures

* 🔄 Charger un artiste aléatoire **sans recharger la page** (fetch AJAX)
* 🔍 Ajout d’une recherche par artiste
* 🎨 Mode sombre / mode clair
* ⚡ Ajouter des animations lors du changement d’artiste
* 🗺️ Intégration d’une carte interactive pour les concerts

---

## 📜 Licence

Projet réalisé dans le cadre du sujet **Groupie Tracker**.

---