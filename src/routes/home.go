package routes

import (
	"fmt"
	"net/http"
)

// This function returns the homepage data
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// This function returns the URL query params

func UrlPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the url query %q", r.URL.Path)
}
