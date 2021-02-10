package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteResponse generic response writer
func WriteResponse(w http.ResponseWriter, response interface{}, status int) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}

// WriteOK 200 response
func WriteOK(w http.ResponseWriter, response interface{}) {
	WriteResponse(w, response, http.StatusOK)
}

// WriteServerError 500 response
func WriteServerError(w http.ResponseWriter, response interface{}) {
	WriteResponse(w, response, http.StatusInternalServerError)
}

// WriteBadRequest 400 response
func WriteBadRequest(w http.ResponseWriter, response interface{}) {
	WriteResponse(w, response, http.StatusBadRequest)
}
