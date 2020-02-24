package utils

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

// RespondWithError adds content type header and writes eror message to w
func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage{message})
}

//RespondWithJSON adds content type header and writes data  to w
func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)

}
