package logic

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetAllArtists(w http.ResponseWriter) []models.Artist {

	var allArtists []models.Artist
	err := ApiCall(w, "artists", &allArtists)
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
	artistId := path[2]

	actualArtistId, _ := strconv.Atoi(artistId)

	allArtists := GetAllArtists(w)

	for _, artist := range allArtists {
		if artist.ID == actualArtistId {
			return artist
		}

	}
	return models.Artist{}
}

func GetLocationsForArtist(w http.ResponseWriter, r *http.Request) models.Locations {

	var allLocations models.LocationsIndex
	err := ApiCall(w, "locations", &allLocations)
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

func GetAllLocations(w http.ResponseWriter, r *http.Request) models.LocationsIndex {
	var allLocations models.LocationsIndex
	err := ApiCall(w, "locations", &allLocations)
	if err != nil {
		fmt.Println(err)
		return models.LocationsIndex{}
	}
	return allLocations
}

func GetAllDates(w http.ResponseWriter, r *http.Request) models.DatesIndex {
	var allDates models.DatesIndex
	err := ApiCall(w, "dates", &allDates)
	if err != nil {
		fmt.Println(err)
		return models.DatesIndex{}
	}
	return allDates
}

func GetAllRelations(w http.ResponseWriter, r *http.Request) models.RelationIndex {
	var allRelations models.RelationIndex
	err := ApiCall(w, "relation", &allRelations)
	if err != nil {
		fmt.Println(err)
		return models.RelationIndex{}
	}
	return allRelations
}

func GetRelations(w http.ResponseWriter, r *http.Request) models.Relations {

	var allRelations models.RelationIndex
	err := ApiCall(w, "relation", &allRelations)
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

	var allDates models.DatesIndex
	err := ApiCall(w, "dates", &allDates)
	if err != nil {
		fmt.Println(err)
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
			return dates
		}
	}

	return models.Dates{}
}

func ApiCall(w http.ResponseWriter, url string, model any) error {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/"+url, nil)
	if err != nil {
		HandleHtml(w, "500")
		fmt.Print(err.Error())
	}
	// add headers to the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		HandleHtml(w, "500")
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil

}

func HandleHtml(w http.ResponseWriter, page string) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("./pages/" + page + ".html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading HTML file: %s", err), http.StatusInternalServerError)
		return
	}

	// Write the HTML content to the response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	w.Write(htmlFile)
}
