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
	header := w.Header()
	header.Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Welcome to the HomePage!!")
}

//ReturnAllArticles used in config.go
func ReturnAllArticles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if len(m.Articles) == 0 {
		NotFound(&w, "No Cars Found")
	} else {
		json.NewEncoder(w).Encode(m.Articles)
	}
}

//ReturnSingleArticle returns single car
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//vars := strings.Split(r.URL.Path, "/")
	//key := vars[2]
	key := ps.ByName("id")
	carFound := false
	//fmt.Fprintf(w, "Key: "+key)
	for _, article := range m.Articles {
		if article.ID == key {
			carFound = true
			header := w.Header()
			header.Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
			break
		}
	}
	if !carFound {
		NotFound(&w, "Car is not found")
	}
}

//CreateArticle creates new car
func CreateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car m.Article
	err := json.Unmarshal(reqBody, &car)
	if err != nil {
		BadRequest(&w, "Send A Valid Request")
		err = nil
	} else {
		m.Articles = append(m.Articles, car)
		header := w.Header()
		header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(car)
	}
}

//DeleteArticle deletes a car
func DeleteArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	for index, article := range m.Articles {
		if article.ID == key {
			m.Articles = append(m.Articles[:index], m.Articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			break
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
			header := w.Header()
			header.Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			break
		}
	}
}
