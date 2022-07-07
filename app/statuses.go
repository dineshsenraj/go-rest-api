package app

import (
	m "github.com/dineshsenraj/go-rest-api/model"
	"encoding/json"
	"net/http"
	"os"
)

var errResp m.Error

//InternalServerError Response
func InternalServerError(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 500, MoreInfo: msg, Message: os.Getenv("MSG_INTERNAL_SERVER")}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(*w).Encode(errResp)
}

//BadRequest Response
func BadRequest(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 400, MoreInfo: msg, Message: os.Getenv("MSG_BAD_REQUEST")}
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusBadRequest)
	json.NewEncoder(*w).Encode(errResp)
}

//NotFound Response
func NotFound(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 404, MoreInfo: msg, Message: os.Getenv("MSG_NOT_FOUND")}
	header := (*w).Header()
	header.Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusNotFound)
	json.NewEncoder(*w).Encode(errResp)
}
