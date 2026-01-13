package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var httpClient = &http.Client{Timeout: 5 * time.Second}

func FetchGroups() ([]Group, error) {
	const url = "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var groups []Group
	if err := json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func FetchAudioDB(name string, apiKey string) (*AudioDBArtist, error) {
	escapedName := url.QueryEscape(name)
	url := fmt.Sprintf("https://www.theaudiodb.com/api/v1/json/%s/search.php?s=%s", apiKey, escapedName)

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res AudioDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	if len(res.Artists) == 0 {
		return nil, nil
	}
	return &res.Artists[0], nil
}

func FetchAllConcerts() (map[int]Relation, error) {
	const url = "https://groupietrackers.herokuapp.com/api/relation"
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relations []Relation
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	concertMap := make(map[int]Relation)
	for i, r := range relations {
		concertMap[i] = r
	}

	return concertMap, nil
}

// ------------------------- MERGE DATA -------------------------

func MergeArtistData(group Group, audio *AudioDBArtist, concertMap map[int]Relation, index int) ArtistFull {
	artist := ArtistFull{
		ID:           group.ID,
		Name:         group.Name,
		Image:        group.Image,
		Members:      group.Members,
		CreationDate: group.CreationDate,
		FirstAlbum:   group.FirstAlbum,
	}

	if audio != nil {
		artist.Genre = audio.Genre
		artist.Country = audio.Country
		artist.BiographyFR = audio.BiographyFR
		artist.BiographyEN = audio.BiographyEN
		artist.Thumb = audio.Thumb
		artist.Banner = audio.Banner
		
		if audio.Website != "" && !strings.HasPrefix(audio.Website, "http") {
			artist.Website = "https://" + audio.Website
		} else {
			artist.Website = audio.Website
		}

		if audio.Youtube != "" && !strings.HasPrefix(audio.Youtube, "http") {
			artist.Youtube = "https://www.youtube.com/" + audio.Youtube
		} else {
			artist.Youtube = audio.Youtube
		}

		if audio.Instagram != "" && !strings.HasPrefix(audio.Instagram, "http") {
			artist.Instagram = "https://www.instagram.com/" + audio.Instagram
		} else {
			artist.Instagram = audio.Instagram
		}

		artist.Facebook = audio.Facebook
		artist.LastFM = audio.LastFM
		artist.MusicBrainz = audio.MusicBrainz
	}

	if rel, ok := concertMap[index]; ok {
		concerts := make([]Concert, len(rel.Dates))
		for i := range rel.Dates {
			location, city := "Lieu inconnu", "Ville inconnue"
			if i < len(rel.Locations) {
				location = rel.Locations[i]
			}
			if i < len(rel.Cities) {
				city = rel.Cities[i]
			}
			concerts[i] = Concert{
				Date:     rel.Dates[i],
				Location: location,
				City:     city,
			}
		}
		artist.Concerts = concerts
	}

	return artist
}