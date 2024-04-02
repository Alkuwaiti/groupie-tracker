package logic

import (
	"encoding/json"
	"fmt"
	models "groupie/models"
	"net/http"
	"strings"
)

func GetAllArtists() []models.Artist {

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/artists", nil)
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

	var responses []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return responses

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
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/locations", nil)
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

	var allLocations models.LocationsIndex
	if err := json.NewDecoder(resp.Body).Decode(&allLocations); err != nil {
		fmt.Print(err.Error())
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

func ApiCall(url string, model interface{}) interface{} {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/"+url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	// add headers to the request
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(model); err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return model
}
