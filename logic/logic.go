package logic

import (
	"encoding/json"
	"fmt"
	models "groupie/models"
	"net/http"
	"strings"
)

func GetAllArtists() []models.ResponseArtist {
	// Struct for incoming response data

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

	var responses []models.ResponseArtist
	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return responses

}

func GetArtist(w http.ResponseWriter, r *http.Request) models.ResponseArtist {
	// Split the URL path to extract the parameter
	path := strings.Split(r.URL.Path, "/")
	if len(path) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return models.ResponseArtist{}
	}
	artistName := path[2]
	fmt.Fprintf(w, "Artist name: %s", artistName)

	allArtists := GetAllArtists()

	for _, artist := range allArtists {
		if strings.EqualFold(artist.Name, artistName) {
			return artist
		}

	}
	return models.ResponseArtist{}
}
