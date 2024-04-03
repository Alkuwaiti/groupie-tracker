package handlers

import (
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

	response, err := logic.GetAllArtists(w)

	if err == nil {
		TemplateExecution(w, "index", response)

	}

}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artist, err := logic.GetArtist(w, r)

	if err == nil {
		TemplateExecution(w, "details", artist)

	}

}

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	location, err := logic.GetLocationsForArtist(w, r)

	if err == nil {
		TemplateExecution(w, "locations", location)

	}

}

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := logic.GetDates(w, r)

	if err == nil {
		TemplateExecution(w, "dates", dates)

	}

}

func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := logic.GetRelations(w, r)

	if err == nil {
		TemplateExecution(w, "relations", relations)

	}

}

func AllLocationsHandler(w http.ResponseWriter, r *http.Request) {
	allLocations, err := logic.GetAllLocations(w, r)

	if err == nil {
		TemplateExecution(w, "allLocations", allLocations)

	}

}

func AllDatesHandler(w http.ResponseWriter, r *http.Request) {
	allDates, err := logic.GetAllDates(w, r)

	if err == nil {
		TemplateExecution(w, "allDates", allDates)

	}

}

func AllRelationsHandler(w http.ResponseWriter, r *http.Request) {
	allRelations, err := logic.GetAllRelations(w, r)

	if err == nil {
		TemplateExecution(w, "allRelations", allRelations)

	}

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
