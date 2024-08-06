package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Get ID from URL
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Retrive id from url
	params := httprouter.ParamsFromContext(r.Context())

	// Convert id to int64 and check if it's nil or less than 1
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

// Write JSON response
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encoding Object to JSON
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Newline for better view in Terminal
	js = append(js, '\n')

	// If header found add it to response
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Application response type
	w.Header().Set("Content-Type", "application/json")

	// Write status
	w.WriteHeader(status)

	// Write JSON to response body
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}
