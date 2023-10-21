package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"errors"
	"database/sql"


	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)


var movie services.Movie

// GET/movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	all, err := movie.GetAllMovies()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"movies": all})
}

// GET/movie/{id}
func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	movieData, movieErr := movie.GetMovieByID(id)
	
    if movieErr != nil {
        if movieErr == sql.ErrNoRows {
            helpers.ErrorJSON(w, errors.New("movie not found"), http.StatusNotFound)
        } else {
            helpers.ErrorJSON(w, movieErr, http.StatusBadRequest)
        }
        return
    }

	helpers.WriteJSON(w, http.StatusOK, movieData)
}

// POST/movie - id passed through JSON body
func CreateMovieByID(w http.ResponseWriter, r *http.Request){
	var movieData services.Movie
	err := json.NewDecoder(r.Body).Decode(&movieData) // decodes r.Body and stores it in &movieData by populating the fields of the Movie Data struct we created in services/myflicklist.go
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	movieCreated, err := movie.CreateMovieByID(movieData)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, movieCreated)
}
