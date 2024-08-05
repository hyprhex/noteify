package main

import (
	"fmt"
	"net/http"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new article")
}

func (app *application) showArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Check if id exists
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Print out Article
	fmt.Fprintf(w, "show the details of article %d\n", id)
}
