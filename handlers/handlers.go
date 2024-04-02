package handlers

import (
	"fmt"
	logic "groupie/logic"
	"html/template"
	"net/http"
	"os"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// return an error if wrong path
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	response := logic.GetAllArtists()

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

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	HandleHtml(w, "404")
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

func HandleHtml(w http.ResponseWriter, page string) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("./pages/" + page + ".html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading HTML file: %s", err), http.StatusInternalServerError)
		return
	}

	// Write the HTML content to the response
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}
