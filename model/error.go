package model

//Error model
type Error struct {
	code     int    `json:"code"`
	message  string `json:"message"`
	moreInfo string `json:"moreInfo"`
}

var error Error
