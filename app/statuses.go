package app

import (
	m "RESTApi/go-rest-api/model"
	"encoding/json"
	"net/http"
)

var errResp m.Error

//InternalServerError Response
func InternalServerError(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 500, MoreInfo: msg, Message: "Internal Server Error"}
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(errResp)
}

//BadRequest Response
func BadRequest(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 400, MoreInfo: msg, Message: "Bad Request"}
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(errResp)
}

//NotFound Response
func NotFound(w *http.ResponseWriter, msg string) {
	errResp = m.Error{Code: 404, MoreInfo: msg, Message: "Not Found"}
	header := (*w).Header()
	header.Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusNotFound)
	json.NewEncoder(*w).Encode(errResp)
}
