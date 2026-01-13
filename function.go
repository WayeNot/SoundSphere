package main

import "sort"

func SortArtistsAZ(artists []ArtistFull) {
	sort.SliceStable(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})
}

func SortArtistsZA(artists []ArtistFull) {
	sort.SliceStable(artists, func(i, j int) bool {
		return artists[i].Name > artists[j].Name
	})
}

func SortArtistsOldToNew(artists []ArtistFull) {
	sort.SliceStable(artists, func(i, j int) bool {
		return artists[i].CreationDate < artists[j].CreationDate
	})
}

func SortArtistsNewToOld(artists []ArtistFull) {
	sort.SliceStable(artists, func(i, j int) bool {
		return artists[i].CreationDate > artists[j].CreationDate
	})
}