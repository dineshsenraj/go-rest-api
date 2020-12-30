package app

import (
	m "RESTApi/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Homepage refered in config.go
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Endpoint Hit : HomePage")
}

//ReturnAllArticles used in config.go
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(m.Articles)
}

//ReturnSingleArticle returns single car
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(r.URL.Path, "/")
	key := vars[2]

	fmt.Fprintf(w, "Key: "+key)
	for _, article := range m.Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
