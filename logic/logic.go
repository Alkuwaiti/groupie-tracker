package logic

import (
	"encoding/json"
	"fmt"
	models "groupie/models"
	"net/http"
	"strings"
)

func GetAllArtists() []models.Artist {

	var allArtists []models.Artist
	err := ApiCall("artists", &allArtists)
	if err != nil {
		fmt.Println(err)
		return []models.Artist{}
	}

	return allArtists

}

func GetArtist(w http.ResponseWriter, r *http.Request) models.Artist {
	// Split the URL path to extract the parameter
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Artist{
			ID:           0,
			Image:        "",
			Name:         "",
			Members:      []string{},
			CreationDate: 0,
			FirstAlbum:   "",
		}
	}
	artistName := path[2]

	allArtists := GetAllArtists()

	for _, artist := range allArtists {
		if strings.EqualFold(artist.Name, artistName) {
			return artist
		}

	}
	return models.Artist{}
}

func GetLocationsForArtist(w http.ResponseWriter, r *http.Request) models.Locations {

	var allLocations models.LocationsIndex
	err := ApiCall("locations", &allLocations)
	if err != nil {
		fmt.Println(err)
		return models.Locations{}
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Locations{}
	}

	artist := GetArtist(w, r)

	for _, locations := range allLocations.Index {
		if artist.ID == locations.ID {
			return locations
		}
	}

	return models.Locations{}
}

func GetRelations(w http.ResponseWriter, r *http.Request) models.Relations {

	var allRelations models.RelationIndex
	err := ApiCall("relation", &allRelations)
	if err != nil {
		fmt.Println(err)
		return models.Relations{}
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Relations{}
	}

	artist := GetArtist(w, r)

	for _, relations := range allRelations.Index {
		if artist.ID == relations.ID {
			return relations
		}
	}

	return models.Relations{}
}

func GetDates(w http.ResponseWriter, r *http.Request) models.Dates {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/dates", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	// add headers to the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	var allDates models.DatesIndex
	if err := json.NewDecoder(resp.Body).Decode(&allDates); err != nil {
		fmt.Print(err.Error())
		return models.Dates{}
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Dates{}
	}

	artist := GetArtist(w, r)

	for _, dates := range allDates.Index {
		if artist.ID == dates.ID {
			fmt.Println(dates)
			return dates
		}
	}

	return models.Dates{}
}

func ApiCall(url string, model any) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/"+url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	// add headers to the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil

}
