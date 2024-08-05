package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new article")
}

func (app *application) showArticleHandler(w http.ResponseWriter, r *http.Request) {
	// Retrive ID from the url
	params := httprouter.ParamsFromContext(r.Context())

	// Convert id to int and check is big than 0
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	// Print out Article
	fmt.Fprintf(w, "show the details of article %d\n", id)
}
