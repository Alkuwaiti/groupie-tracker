package handlers

import (
	"fmt"
	logic "groupie/logic"
	"net/http"
	"os"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// return an error if wrong path
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	logic.GetAllArtists()

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
