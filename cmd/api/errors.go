package main

import (
	"fmt"
	"net/http"
)

// Generic helper for logging an error message.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Print(err)
}

// Generic method for sending JSON-formatted error
func (app *application) errResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	// Write the response then log it
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// If application encounter a problem at runtime
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := " the server encountered a problem and could not process your request"
	app.errResponse(w, r, http.StatusInternalServerError, message)
}

// Notfound method error response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errResponse(w, r, http.StatusNotFound, message)
}

// Not allowed method error response
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errResponse(w, r, http.StatusMethodNotAllowed, message)
}
