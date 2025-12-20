package main

type PageData struct {
	Groups       []Group
	GroupByID    map[int]Group
	AudioDBCache map[string]*AudioDBArtist
	Settings     Settings
}

type Settings struct {
	DarkMode bool
}

type ArtistPage struct {
	Artist   ArtistFull
	Settings Settings
}