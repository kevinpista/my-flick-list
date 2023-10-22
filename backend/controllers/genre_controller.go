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

var genre services.Genre

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
	var genreData services.Genre

	err := json.NewDecoder(r.Body).Decode(&genreData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	genreDataCreated, err := genreData.CreateGenreDataByMovieID(genreData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, genreDataCreated)

}
