package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/go-chi/chi/v5"
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

// POST/movie
func CreateMovie(w http.ResponseWriter, r *http.Request){
	var movieData services.Movie
	err := json.NewDecoder(r.Body).Decode(&movieData)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	movieCreated, err := movie.CreateMovie(movieData)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, movieCreated)
}

// POST/movie/{id}
func CreateMovieById(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr) // parameter will be a string. convert to an int for our Movie struct field
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	var movieData services.Movie
	err = json.NewDecoder(r.Body).Decode(&movieData) // decodes r.Body and stores it in &movieData by populating the fields of the Movie Data struct we created in services/myflicklist.go
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	movieCreated, err := movie.CreateMovieById(movieData, id)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, movieCreated)
}
