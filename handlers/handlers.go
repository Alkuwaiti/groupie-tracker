package handlers

import (
	"fmt"
	"groupie/logic"
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// return an error if wrong path
	if r.URL.Path != "/" && r.URL.Path != "/artists" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	response := logic.GetAllArtists(w)

	TemplateExecution(w, "index", response)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist := logic.GetArtist(w, r)

	TemplateExecution(w, "details", artist)
}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	location := logic.GetLocationsForArtist(w, r)

	TemplateExecution(w, "locations", location)
}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates := logic.GetDates(w, r)

	TemplateExecution(w, "dates", dates)
}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations := logic.GetRelations(w, r)

	TemplateExecution(w, "relations", relations)
}

func AllLocationsHandler(w http.ResponseWriter, r *http.Request) {
	allLocations := logic.GetAllLocations(w, r)

	fmt.Println(allLocations)

	TemplateExecution(w, "allLocations", allLocations)
}

func AllDatesHandler(w http.ResponseWriter, r *http.Request) {
	allDates := logic.GetAllDates(w, r)

	TemplateExecution(w, "allDates", allDates)
}

func AllRelationsHandler(w http.ResponseWriter, r *http.Request) {
	allRelations := logic.GetAllRelations(w, r)

	TemplateExecution(w, "allRelations", allRelations)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	logic.HandleHtml(w, "404")
}

func TemplateExecution(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles("./pages/" + page + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing the data to it
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
