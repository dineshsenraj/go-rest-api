package main

import (
	"RESTApi/go-rest-api/config"
	m "RESTApi/go-rest-api/model"
)

func main() {
	m.Articles = []m.Article{
		m.Article{Title: "Car", Description: "Red Hatch Back", Name: "Honda Jazz", ID: "1"},
		m.Article{Title: "Car", Description: "Black Sedan", Name: "Hyundai Verna", ID: "2"},
		m.Article{Title: "Car", Description: "Green Hatch Back", Name: "Chevrolet Beat", ID: "3"},
		m.Article{Title: "Car", Description: "White SUV", Name: "Kia Seltos", ID: "4"},
	}
	config.HandleRequests()
}
