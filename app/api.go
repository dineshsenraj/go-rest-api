package app

import (
	"github.com/dineshsenraj/go-rest-api/config"
	m "github.com/dineshsenraj/go-rest-api/model"
	"database/sql"
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
	rows, err := config.DB.Query(`select id as ID, car_name as Name, description as Description, 
	title as Title from cars c order by c.id asc`)
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
		NotFound(&w, "No cars found")
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
	row := config.DB.QueryRow(`select id as ID, car_name as Name, description as Description, 
	title as Title from cars c where id=$1`, key)
	car := m.Car{}
	switch err := row.Scan(&car.ID, &car.Name, &car.Description, &car.Title); err {
	case sql.ErrNoRows:
		NotFound(&w, "Car not found")
	case nil:
		header := w.Header()
		header.Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(car)
	default:
		log.Println("Row Scan Error")
		log.Println(err.Error())
		InternalServerError(&w, "Row Scan Error")
	}
}

//CreateCar creates new car
func CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car m.Car
	err := json.Unmarshal(reqBody, &car)
	if err != nil {
		BadRequest(&w, "Invalid Request")
	}
	result, err := config.DB.Exec(`INSERT INTO public.cars
	(id, car_name, description, title)
	VALUES($1, $2, $3, $4)`, car.ID, car.Name, car.Description, car.Title)
	if err != nil {
		InternalServerError(&w, "Query Error")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		InternalServerError(&w, "DB Error")
	}
	if rows != 1 {
		InternalServerError(&w, fmt.Sprintf("Expected to insert 1 row, but inserted %d rows", rows))
	}
	header := w.Header()
	header.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

//DeleteCar deletes a car
func DeleteCar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	result, err := config.DB.Exec(`DELETE FROM public.cars
	WHERE id=$1`, key)
	if err != nil {
		InternalServerError(&w, "Query Error")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		InternalServerError(&w, "DB Error")
	}
	if rows > 1 {
		InternalServerError(&w, fmt.Sprintf("Expected to delete 1 row, but deleted %d rows", rows))
	}
	if rows == 0 {
		NotFound(&w, "No Records Found to Delete")
	}
	w.WriteHeader(http.StatusNoContent)
}

//UpdateCar updates a car
func UpdateCar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("id")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var car m.Car
	json.Unmarshal(reqBody, &car)
	result, err := config.DB.Exec(`UPDATE public.cars
	SET id=$1, car_name=$2, description=$3, title=$4 
	where id=$5`, car.ID, car.Name, car.Description, car.Title, key)
	if err != nil {
		InternalServerError(&w, "Query Error")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		InternalServerError(&w, "DB Error")
	}
	if rows > 1 {
		InternalServerError(&w, fmt.Sprintf("Expected to update 1 row, but updated %d rows", rows))
	}
	if rows == 0 {
		NotFound(&w, "No Records Found to Update")
	}
	header := w.Header()
	header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)

}
