package config

import (
	app "RESTApi/app"
	"log"
	"net/http"
)

//HandleRequests used in main.go
func HandleRequests() {
	//myRouter := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/", app.Homepage)
	http.HandleFunc("/cars", app.ReturnAllArticles)
	http.HandleFunc("/cars/*/return", app.ReturnSingleArticle)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
