package tmdb_controllers

import (
	"github.com/joho/godotenv"
	"encoding/json"
	"net/http"

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

var searchResults tmdb_services.MovieSearchService

// GET/search?query={keyword+keyword..}
func SearchMovieByKeyWords(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query") // holds literal search string of "Harry+Potter"

	allResults, err:= searchResults.TMDBSearchMovieByKeywords(query) // will be returned with a formatted JSON of the data we need for search results page

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"search_results": allResults})
}

// https://api.themoviedb.org/3/search/movie?query=Harry+Potter 