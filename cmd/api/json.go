package main

import (
	"encoding/json"
	"net/http"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func writeJson(w http.ResponseWriter, status int, data any) error{
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	 
	return json.NewEncoder(w).Encode(data)
}

func readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := int64(1024 * 1024) // 1MB max
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	
	return dec.Decode(data)
}

func writeJsonError(w http.ResponseWriter, status int, message string) error {
	// errorResponse := map[string]string{"error": message}
	type envelope struct {
		Error string `json:"error"`
	}

	return writeJson(w, status, &envelope{Error: message})
}

func (app *application) jsonResponse (w http.ResponseWriter, status int, data any) error{
	type envelope struct {
		Data any `json:"data"`
	}

	return writeJson(w, status, &envelope{Data: data})
}