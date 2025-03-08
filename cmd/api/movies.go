package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"greenlight.altamash.dev/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeResponse(w, struct {
		Message string
		Code    int
		Error   bool
		Body    interface{}
	}{
		Message: "Success",
		Code:    200,
		Error:   false,
		Body:    movie,
	})
	if err != nil {
		app.logger.Error(err.Error())
		app.serverErrorResponse(w, r, err)
	}
}
