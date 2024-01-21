package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"database/sql"
	"errors"


	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

var genre services.GenreService

// GET/genre/{movieID}
func GetGenreByMovieID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "movieID")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	genreData, genreErr := genre.GetGenreByMovieID(id)
	
    if genreErr != nil {
        if genreErr == sql.ErrNoRows {
            helpers.ErrorJSON(w, errors.New("genre data not found"), http.StatusNotFound)
        } else {
            helpers.ErrorJSON(w, genreErr, http.StatusBadRequest)
        }
        return
    }

	helpers.WriteJSON(w, http.StatusOK, genreData)
}

// POST/genre
func CreateGenreDataByMovieID(w http.ResponseWriter, r *http.Request) {
	var genreData services.GenreService

	err := json.NewDecoder(r.Body).Decode(&genreData.Genre)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	genreDataCreated, err := genreData.CreateGenreDataByMovieID(genreData.Genre)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, genreDataCreated)

}

// GET/genre -- TESTING PURPOSES ONLY
func GetAllGenres(w http.ResponseWriter, r *http.Request) {
	genreData, err := genre.GetAllGenres()
	
    if err != nil {
        if err == sql.ErrNoRows {
            helpers.ErrorJSON(w, errors.New("genre data not found"), http.StatusNotFound)
        } else {
            helpers.ErrorJSON(w, err, http.StatusBadRequest)
        }
        return
    }
	helpers.WriteJSON(w, http.StatusOK, genreData)
}