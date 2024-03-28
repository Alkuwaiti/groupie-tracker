package main

import (
	"fmt"
	handlers "groupie/handlers"
	"net/http"
)

func main() {
	fmt.Print()
	http.HandleFunc("/", handlers.RootHandler)

	http.ListenAndServe(":3000", nil)
}
