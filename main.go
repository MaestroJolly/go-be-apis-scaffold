package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaestroJolly/go-be-apis-scaffold/src/middlewares"
	"github.com/MaestroJolly/go-be-apis-scaffold/src/routes"
	"github.com/MaestroJolly/go-be-apis-scaffold/src/services/greetings"
)

func handleRequests() {
	homePage := routes.HomePage
	urlPath := routes.UrlPath
	getArticles := routes.GetArticles
	isAuthorized := middlewares.IsAuthorized // jwt authentication middleware
	http.HandleFunc("/", homePage)
	http.Handle("/path/", isAuthorized(urlPath))
	http.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	greeting := greetings.Greetings("Jide")
	fmt.Println(greeting)
	handleRequests()
}
