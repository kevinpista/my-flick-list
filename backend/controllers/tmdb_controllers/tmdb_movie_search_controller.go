package tmdb_controllers

import (
	// "github.com/joho/godotenv"
	// "encoding/json"
	// "strings"
	"net/http"
	"errors"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services/tmdb_services"
)

/*
func main(){
err := godotenv.Load(".env")
if err != nil {
	// Handle error, e.g., log it and/or terminate the program
	log.Fatal("Error loading .env file")
}

apiKey := os.Getenv("API_KEY")

}
*/

var searchResults tmdb_services.TMDBMovieSearchService

// GET/search?query={keyword+keyword..}&page={pageNumber}
func TMDBSearchMovieByKeyWords(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // If the user passes it as "harry+potter"
	// this method will transform it into string of "harry potter" with no + sign.
	// FRONTENDS will pass query url as "harry%2Bpotter" with "%2B" between spaces to preserve + sign
    page := r.URL.Query().Get("page")

	if query == "" {
		helpers.MessageLogs.ErrorLog.Println("User sent in an empty query")
		helpers.ErrorJSON(w, errors.New("search query cannot be empty"), http.StatusBadRequest)
		return
	}

    // query = strings.Replace(query, " ", "+", -1)
	// helpers.MessageLogs.ErrorLog.Println(query)

	allResults, err:= searchResults.TMDBSearchMovieByKeywords(query, page) // will be returned with a formatted JSON of the data we need for search results page

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

	// If TMDB API returned no results, return a 204 No Content code along with empty JSON
	if allResults.TotalResults == 0 {
		helpers.MessageLogs.ErrorLog.Println("No movies found with search query")
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{}) // returning JSON object optional
		return
	}

	helpers.WriteJSON(w, http.StatusOK, allResults)
}

// https://api.themoviedb.org/3/search/movie?query=Harry+Potter&pages=1