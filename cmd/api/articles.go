package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/hyprhex/noteify/internal/data"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new article")
}

func (app *application) showArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Check if id exists
	id, err := app.readIDParam(r)
	if err != nil {
		// Use not found response helper
		app.notFoundResponse(w, r)
		return
	}

	// Instance of the Article struct
	article := data.Article{
		ID:          id,
		Title:       "How to write Go program?",
		Description: "wdjojsda jasndansdi djasndoajsdjaoidjoiasjdosajdojasod",
		Genres:      []string{"Go", "backend", "development"},
		CreatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"article": article}, nil)
	if err != nil {
		// use server error response helper
		app.serverErrorResponse(w, r, err)
	}
}
