package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllArtists() {
	// Struct for incoming response data
	type Response struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
	}

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

	var responses []Response
	if err := json.NewDecoder(resp.Body).Decode(&responses); err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println(responses)
}
