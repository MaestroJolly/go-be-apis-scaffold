package articles

import (
	"github.com/MaestroJolly/go-be-apis-scaffold/src/utils"
)

var allArticles []utils.Article

func GetArticles() []utils.Article {
	Articles := []utils.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	return Articles
}
