package app

import (
	m "RESTApi/go-rest-api/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Homepage refered in config.go
func Homepage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to the HomePage!!")
	fmt.Println("Endpoint Hit : HomePage")
}

//ReturnAllArticles used in config.go
func ReturnAllArticles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(m.Articles)
}

//ReturnSingleArticle returns single car
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//vars := strings.Split(r.URL.Path, "/")
	//key := vars[2]
	key := ps.ByName("id")
	//fmt.Fprintf(w, "Key: "+key)
	for _, article := range m.Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

//CreateArticle creates new car
func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car m.Article
	json.Unmarshal(reqBody, &car)
	m.Articles = append(m.Articles, car)
	json.NewEncoder(w).Encode(car)
}

//DeleteArticle deletes a car
func DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	for index, article := range m.Articles {
		if article.ID == key {
			m.Articles = append(m.Articles[:index], m.Articles[index+1:]...)
		}
	}
}

//UpdateArticle updates a car
func UpdateArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	for index, article := range m.Articles {
		if article.ID == key {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var car m.Article
			json.Unmarshal(reqBody, &car)
			m.Articles = append(m.Articles[:index], m.Articles[index+1:]...)
			m.Articles = append(m.Articles, car)
			json.NewEncoder(w).Encode(car)
			break
		}
	}
}
