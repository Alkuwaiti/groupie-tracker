package main

import (
	"fmt"
	handlers "groupie/handlers"
	"net/http"
)

func main() {
	fmt.Print()
	http.HandleFunc("/artists/", handlers.ArtistHandler)
	http.HandleFunc("/locations/", handlers.LocationsHandler)
	http.HandleFunc("/dates/", handlers.DatesHandler)
	http.HandleFunc("/relations/", handlers.RelationsHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(":3000", nil)
	// TODO: relations api
}
