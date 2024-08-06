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
		http.NotFound(w, r)
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
		app.logger.Print(err)
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}
}
