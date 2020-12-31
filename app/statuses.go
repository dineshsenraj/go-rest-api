package app

import (
	m "RESTApi/go-rest-api/model"
	"encoding/json"
	"net/http"
)

//InternalServerError Response
func InternalServerError(w *http.ResponseWriter, msg string) {
	error := m.Error{
		Error{code: 500, message: msg},
	}
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}
