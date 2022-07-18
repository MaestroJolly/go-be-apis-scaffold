package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaestroJolly/go-be-apis-scaffold/src/routes"
	"github.com/MaestroJolly/go-be-apis-scaffold/src/services/greetings"
)

func handleRequests() {
	homePage := routes.HomePage
	urlPath := routes.UrlPath
	getArticles := routes.GetArticles
	http.HandleFunc("/", homePage)
	http.HandleFunc("/path/", urlPath)
	http.HandleFunc("/articles", getArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	greeting := greetings.Greetings("Jide")
	fmt.Println(greeting)
	handleRequests()
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }
