package routes

import (
	"encoding/json"
	"net/http"

	"github.com/MaestroJolly/go-be-apis-scaffold/src/services/articles"
)

func GetArticles(w http.ResponseWriter, r *http.Request) {
	allArticles := articles.GetArticles()
	json.NewEncoder(w).Encode(allArticles)
}
