package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(&data)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	PanicIfError(err)
}