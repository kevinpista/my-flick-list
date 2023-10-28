package tmdb_controllers

import (
	"net/http"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services/tmdb_services"
)


var movieResult tmdb_services.TMDBMovieService

// GET/search?query={movie_id}
func TMDBGetMovieByKeywords(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // user passes movie ID


	result, err:= movieResult.TMDBGetMovieByKeywords(query) // will be returned with a formatted JSON of the data we need for search results page

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	// TODO - when implementing frontend, will need to structure json response
	// into a format that the frontend can use to display each search result cleanly on a dedicated movie page

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"movie": result})
}

// https://api.themoviedb.org/3/movie/{movie_id}