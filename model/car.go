package model

//Car model
type Car struct {
	Title       string `json:"Title"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	ID          string `json:"ID"`
}

//Cars array defined
var Cars []Car
