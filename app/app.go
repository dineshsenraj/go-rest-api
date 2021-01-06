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

//ReturnAllCars used in config.go
func ReturnAllCars(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	if len(m.Cars) == 0 {
		NotFound(&w, "No Cars Found")
	} else {
		json.NewEncoder(w).Encode(m.Cars)
	}
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
