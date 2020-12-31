package config

import (
	app "RESTApi/go-rest-api/app"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleRequests used in main.go
func HandleRequests() {
	myRouter := httprouter.New()
	//myRouter.HandleMethodNotAllowed = true
	myRouter.GET("/", app.Homepage)
	myRouter.GET("/cars", app.ReturnAllArticles)
	myRouter.GET("/cars/", app.ReturnAllArticles)
	myRouter.GET("/cars/:id/", app.ReturnSingleArticle)
	myRouter.GET("/cars/:id", app.ReturnSingleArticle)
	myRouter.POST("/cars", app.CreateArticle)
	myRouter.POST("/cars/", app.CreateArticle)
	myRouter.DELETE("/cars/:id/", app.DeleteArticle)
	myRouter.DELETE("/cars/:id", app.DeleteArticle)
	myRouter.PUT("/cars/:id/", app.UpdateArticle)
	myRouter.PUT("/cars/:id", app.UpdateArticle)
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}
