package main

import (
	"fmt"
	"log"
	"net/http"

	"pkgcentral.dev/go-be-apis-scaffold/src/services/greetings"
)

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	greeting := greetings.Greetings("Jide")
	fmt.Println(greeting)
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
