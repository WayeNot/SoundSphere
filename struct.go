package main

type Settings struct {
	DarkMode bool
}

type Concert struct {
	Date     string
	Location string
	City     string
}

type Group struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type AudioDBArtist struct {
	Name        string `json:"strArtist"`
	BiographyFR string `json:"strBiographyFR"`
	BiographyEN string `json:"strBiographyEN"`
	Genre       string `json:"strGenre"`
	Country     string `json:"strCountry"`
	Thumb       string `json:"strArtistThumb"`
	Banner      string `json:"strArtistBanner"`

	Facebook    string `json:"strFacebook"`
	Twitter     string `json:"strTwitter"`
	Instagram   string `json:"strInstagram"`
	Website     string `json:"strWebsite"`
	Youtube     string `json:"strYoutube"`
	LastFM      string `json:"strLastFMChart"`
	MusicBrainz string `json:"strMusicBrainzID"`
}

type AudioDBResponse struct {
	Artists []AudioDBArtist `json:"artists"`
}

type Relation struct {
	ID        int      `json:"id"`
	Dates     []string `json:"dates"`
	Locations []string `json:"locations"`
	Cities    []string `json:"cities"`
}

type ArtistFull struct {
	ID           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   string

	Genre        string
	Country      string
	BiographyFR  string
	BiographyEN  string
	Thumb        string
	Banner       string

	Facebook    string
	Twitter     string
	Instagram   string
	Website     string
	Youtube     string
	LastFM      string
	MusicBrainz string

	Concerts []Concert
}

type PageData struct {
	Groups       []Group
	GroupByID    map[int]Group
	AllConcerts  map[int][]Concert
	AudioDBCache map[string]*AudioDBArtist
	Settings     Settings
}

type ListPage struct {
	Groups      []ArtistFull
	Settings    Settings
	Filter      string
	Search      string
	CurrentPage int
	PerPage     int
	TotalPages  int
	PageNumbers []int
	PrevPage    int
	NextPage    int
}

type ArtistPage struct {
	Artist   ArtistFull
	Settings Settings
}