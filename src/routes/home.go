package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaestroJolly/go-be-apis-scaffold/src/middlewares"
)

// This function returns the homepage data
func HomePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := middlewares.GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	response := map[string]string{"token": validToken, "greetings": "Welcome to the HomePage!"}

	json.NewEncoder(w).Encode(response)

	fmt.Println("Endpoint Hit: homePage")
}

// This function returns the URL query params

func UrlPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the url query %q", r.URL.Path)
}
