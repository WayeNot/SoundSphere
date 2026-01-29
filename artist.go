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

func FetchAudioDB(name string, apiKey string) (*AudioDBArtist, error) {
	escapedName := url.QueryEscape(name)
	apiURL := fmt.Sprintf(
		"https://www.theaudiodb.com/api/v1/json/%s/search.php?s=%s",
		apiKey,
		escapedName,
	)

	resp, err := httpClient.Get(apiURL)
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

type RelationAPIResponse struct {
	Index []RelationAPI `json:"index"`
}

type RelationAPI struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchAllConcerts() (map[int][]Concert, error) {
	const apiURL = "https://groupietrackers.herokuapp.com/api/relation"

	resp, err := httpClient.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp RelationAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	concertMap := make(map[int][]Concert)

	for _, rel := range apiResp.Index {
		var concerts []Concert

		for location, dates := range rel.DatesLocations {
			for _, date := range dates {
				concerts = append(concerts, Concert{
					Date:     date,
					Location: location,
					City:     extractCity(location),
				})
			}
		}

		concertMap[rel.ID] = concerts
	}

	return concertMap, nil
}

func extractCity(location string) string {
	parts := strings.Split(location, "-")
	if len(parts) == 0 {
		return ""
	}
	return strings.ReplaceAll(parts[0], "_", " ")
}

func MergeArtistData(
	group Group,
	audio *AudioDBArtist,
	concertMap map[int][]Concert,
	index int,
) ArtistFull {

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

		artist.Website = normalizeURL(audio.Website, "https://")
		artist.Youtube = normalizeURL(audio.Youtube, "https://www.youtube.com/")
		artist.Instagram = normalizeURL(audio.Instagram, "https://www.instagram.com/")

		artist.Facebook = audio.Facebook
		artist.LastFM = audio.LastFM
		artist.MusicBrainz = audio.MusicBrainz
	}

	if concerts, ok := concertMap[index]; ok {
		artist.Concerts = concerts
	}

	return artist
}

func normalizeURL(value, prefix string) string {
	if value == "" {
		return ""
	}
	if strings.HasPrefix(value, "http") {
		return value
	}
	return prefix + value
}

func FetchGroups() ([]Group, error) {
	const apiURL = "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := httpClient.Get(apiURL)
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