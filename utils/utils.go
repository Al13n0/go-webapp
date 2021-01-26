package utils

import (
	"books-list/models"
	"encoding/json"
	"net/http"
)

//SendError function used to send erro message
func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

//SendSuccess send success status code in the server reply
func SendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
