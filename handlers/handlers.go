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

	// Parse the HTML/EJS template
	tmpl, err := template.ParseFiles("./pages/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing the data to it
	err = tmpl.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// space
// space
// space
// space
// space
// space
// space
// space

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	HandleHtml(w, "404")

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
