package logic

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetAllArtists(w http.ResponseWriter) ([]models.Artist, error) {

	var allArtists []models.Artist
	err := ApiCall(w, "artists", &allArtists)
	if err != nil {
		fmt.Println(err)
		return []models.Artist{}, err
	}

	return allArtists, err

}

func GetArtist(w http.ResponseWriter, r *http.Request) (models.Artist, error) {
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
		}, nil
	}
	artistId := path[2]

	actualArtistId, _ := strconv.Atoi(artistId)

	allArtists, err := GetAllArtists(w)

	for _, artist := range allArtists {
		if artist.ID == actualArtistId {
			return artist, err
		}

	}
	return models.Artist{}, err
}

func GetLocationsForArtist(w http.ResponseWriter, r *http.Request) (models.Locations, error) {

	var allLocations models.LocationsIndex
	err := ApiCall(w, "locations", &allLocations)
	if err != nil {
		fmt.Println(err)
		return models.Locations{}, err
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Locations{}, err
	}

	artist, err := GetArtist(w, r)

	for _, locations := range allLocations.Index {
		if artist.ID == locations.ID {
			return locations, err
		}
	}

	return models.Locations{}, err
}

func GetAllLocations(w http.ResponseWriter, r *http.Request) (models.LocationsIndex, error) {
	var allLocations models.LocationsIndex
	err := ApiCall(w, "locations", &allLocations)
	if err != nil {
		fmt.Println(err)
		return models.LocationsIndex{}, err
	}
	return allLocations, err
}

func GetAllDates(w http.ResponseWriter, r *http.Request) (models.DatesIndex, error) {
	var allDates models.DatesIndex
	err := ApiCall(w, "dates", &allDates)
	if err != nil {
		fmt.Println(err)
		return models.DatesIndex{}, err
	}
	return allDates, err
}

func GetAllRelations(w http.ResponseWriter, r *http.Request) (models.RelationIndex, error) {
	var allRelations models.RelationIndex
	err := ApiCall(w, "relation", &allRelations)
	if err != nil {
		fmt.Println(err)
		return models.RelationIndex{}, err
	}
	return allRelations, err
}

func GetRelations(w http.ResponseWriter, r *http.Request) (models.Relations, error) {

	var allRelations models.RelationIndex
	err := ApiCall(w, "relation", &allRelations)
	if err != nil {
		fmt.Println(err)
		return models.Relations{}, err
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Relations{}, err
	}

	artist, err := GetArtist(w, r)

	for _, relations := range allRelations.Index {
		if artist.ID == relations.ID {
			return relations, err
		}
	}

	return models.Relations{}, err
}

func GetDates(w http.ResponseWriter, r *http.Request) (models.Dates, error) {

	var allDates models.DatesIndex
	err := ApiCall(w, "dates", &allDates)
	if err != nil {
		fmt.Println(err)
		return models.Dates{}, err
	}

	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.Dates{}, err
	}

	artist, err := GetArtist(w, r)

	for _, dates := range allDates.Index {
		if artist.ID == dates.ID {
			return dates, err
		}
	}

	return models.Dates{}, err
}

func ApiCall(w http.ResponseWriter, url string, model any) error {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/"+url, nil)

	// add headers to the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Some Error Occurred"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return err
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
