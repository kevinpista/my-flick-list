package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

// GET/movie

var movie services.Movie

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	all, err := movie.GetAllMovies()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"movies": all})
}

// POST/movies/movie
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


