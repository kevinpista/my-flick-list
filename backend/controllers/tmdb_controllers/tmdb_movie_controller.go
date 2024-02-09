package tmdb_controllers

import (
	"errors"
	"net/http"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)


var movieResult services.TMDBMovieService

// GET/tmdb-movie?query={movie_id}
func TMDBGetMovieByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // user passes movie ID

	result, err:= movieResult.TMDBGetMovieByID(query)
	// The error return will be in errors.New('error message') already
	if err != nil {
		if err.Error() == "TMDB API is unavailable at this time" {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New("error with TMDB API"), http.StatusBadRequest)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"movie": result})
}

// GET/tmdb-trailer?query={movie_id}
func TMDBGetMovieTrailerByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // user passes movie ID

	result, err:= movieResult.TMDBGetMovieTrailerByID(query)
	// The error return will be in errors.New('error message') already
	if err != nil {
		if err.Error() == "TMDB API is unavailable at this time" {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusInternalServerError)
			return
		}
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New("error with TMDB API"), http.StatusBadRequest)
		return
	}

	if result == nil {
		helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"youtube_video_id": nil})
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"youtube_video_id": result})
}
// Service function makes request to https://api.themoviedb.org/3/movie/{movie_id}/videos