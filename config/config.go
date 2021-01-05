package config

import (
	app "RESTApi/go-rest-api/app"
	"fmt"
	"log"
	"net/http"
	"os"

	m "RESTApi/go-rest-api/model"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

//Configurations used in all files
var Configurations m.Configuration

//Init to intialize the configurations
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Configurations)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	os.Setenv("MSG_BAD_REQUEST", viper.GetString("MSG_BAD_REQUEST"))
	os.Setenv("MSG_NOT_FOUND", viper.GetString("MSG_NOT_FOUND"))
	os.Setenv("MSG_INTERNAL_SERVER", viper.GetString("MSG_INTERNAL_SERVER"))
	// fmt.Println("server port number is ", configuration.Server.Port)
	// fmt.Println("BAD REQUEST ", configuration.MsgBadRequest)
	// fmt.Println("Not Found ", configuration.MsgNotFound)
	// fmt.Println("Internal Server ", configuration.MsgInternalServer)
}

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
	port := fmt.Sprint(":", Configurations.Server.Port)
	// fmt.Println("server port number is ", port)
	// fmt.Println("ENV READ ", os.Getenv("MSG_BAD_REQUEST"))
	log.Fatal(http.ListenAndServe(port, myRouter))
}
