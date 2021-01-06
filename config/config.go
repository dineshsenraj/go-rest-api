package config

import (
	app "RESTApi/go-rest-api/app"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	m "RESTApi/go-rest-api/model"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq" //postgresql
	"github.com/spf13/viper"
)

//Configurations used in all files
var Configurations m.Configuration

//DB connection
var DB *sql.DB

const (
	dbHost = "DB_HOST"
	dbPort = "DB_PORT"
	dbUser = "DB_USER"
	dbPass = "DB_PASS"
	dbName = "DB_NAME"
)

func configDB() map[string]string {
	conf := make(map[string]string)
	conf[dbHost] = Configurations.Database.DBServer
	conf[dbPort] = Configurations.Database.DBPort
	conf[dbUser] = Configurations.Database.DBUser
	conf[dbPass] = Configurations.Database.DBPassword
	conf[dbName] = Configurations.Database.DBName
	return conf
}

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
}

//InitDB connect to database
func InitDB() {
	config := configDB()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbHost], config[dbPort],
		config[dbUser], config[dbPass], config[dbName])

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database Connection Error")
		log.Fatalln(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		log.Println("Database Ping Error")
		log.Fatalln(err.Error())
	}
	log.Println("Database connection established.")
}

//HandleRequests used in main.go
func HandleRequests() {
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
	port := fmt.Sprint(":", Configurations.Server.Port)
	log.Fatal(http.ListenAndServe(port, myRouter))
}
