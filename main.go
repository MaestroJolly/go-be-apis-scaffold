package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaestroJolly/go-be-apis-scaffold/src/middlewares"
	"github.com/MaestroJolly/go-be-apis-scaffold/src/routes"
	"github.com/MaestroJolly/go-be-apis-scaffold/src/services/greetings"
	"github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	homePage := routes.HomePage
	urlPath := routes.UrlPath
	getArticles := routes.GetArticles
	isAuthorized := middlewares.IsAuthorized // jwt authentication middleware
	myRouter.HandleFunc("/", homePage)
	http.Handle("/path/", isAuthorized(urlPath))
	myRouter.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	greeting := greetings.Greetings("Jide")
	fmt.Println(greeting)
	handleRequests()
}
