package main

import (
	"github.com/dineshsenraj/go-rest-api/app"
	"github.com/dineshsenraj/go-rest-api/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.Init()
	config.InitDB()
	myRouter := httprouter.New()
	//myRouter.HandleMethodNotAllowed = true
	myRouter.GET("/", app.Homepage)
	myRouter.GET("/cars", app.ReturnAllCars)
	myRouter.GET("/cars/", app.ReturnAllCars)
	myRouter.GET("/cars/:id/", app.ReturnSingleCar)
	myRouter.GET("/cars/:id", app.ReturnSingleCar)
	myRouter.POST("/cars", app.CreateCar)
	myRouter.POST("/cars/", app.CreateCar)
	myRouter.DELETE("/cars/:id/", app.DeleteCar)
	myRouter.DELETE("/cars/:id", app.DeleteCar)
	myRouter.PUT("/cars/:id/", app.UpdateCar)
	myRouter.PUT("/cars/:id", app.UpdateCar)
	port := fmt.Sprint(":", os.Getenv("GO_SERVER_PORT"))
	log.Fatal(http.ListenAndServe(port, myRouter))
}
