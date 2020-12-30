package model

//Article model
type Article struct {
	Title       string `json:"Title"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	ID          string `json:"ID"`
}

//Articles array defined
var Articles []Article
