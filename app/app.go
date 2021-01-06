package app

import (
	"RESTApi/go-rest-api/config"
	m "RESTApi/go-rest-api/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Homepage refered in config.go
func Homepage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	header := w.Header()
	header.Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Welcome to the HomePage!!")
}

//ReturnAllCars used in config.go
func ReturnAllCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var carsList []m.Car
	rows, err := config.DB.Query(`select id as ID, name as Name, description as Description, title as Title from cars c`)
	if err != nil {
		log.Println("Database Error")
		log.Println(err.Error())
		InternalServerError(&w, "Database Error")
	}
	defer rows.Close()
	err = rows.Err()
	if err != nil {
		log.Println("Query Error")
		log.Println(err.Error())
		InternalServerError(&w, "Query Error")
	}

	for rows.Next() {
		car := m.Car{}
		err = rows.Scan(
			&car.ID,
			&car.Name,
			&car.Description,
			&car.Title,
		)
		if err != nil {
			log.Println("Row Scan Error")
			log.Println(err.Error())
			InternalServerError(&w, "Row Scan Error")
		}
		carsList = append(carsList, car)
	}

	if len(carsList) == 0 {
		NotFound(&w, "Car is not found")
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carsList)
}

//ReturnSingleCar returns single car
func ReturnSingleCar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//vars := strings.Split(r.URL.Path, "/")
	//key := vars[2]
	key := ps.ByName("id")
	carFound := false
	//fmt.Fprintf(w, "Key: "+key)
	for _, car := range m.Cars {
		if car.ID == key {
			carFound = true
			header := w.Header()
			header.Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			break
		}
	}
	if !carFound {
		NotFound(&w, "Car is not found")
	}
}

//CreateCar creates new car
func CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car m.Car
	err := json.Unmarshal(reqBody, &car)
	if err != nil {
		BadRequest(&w, "Send A Valid Request")
		err = nil
	} else {
		m.Cars = append(m.Cars, car)
		header := w.Header()
		header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(car)
	}
}

//DeleteCar deletes a car
func DeleteCar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	for index, car := range m.Cars {
		if car.ID == key {
			m.Cars = append(m.Cars[:index], m.Cars[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			break
		}
	}
}

//UpdateCar updates a car
func UpdateCar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	for index, car := range m.Cars {
		if car.ID == key {
			reqBody, _ := ioutil.ReadAll(r.Body)
			var car m.Car
			json.Unmarshal(reqBody, &car)
			m.Cars = append(m.Cars[:index], m.Cars[index+1:]...)
			m.Cars = append(m.Cars, car)
			header := w.Header()
			header.Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			break
		}
	}
}
