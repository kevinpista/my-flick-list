package tmdb_controllers

import (
	// "github.com/joho/godotenv"
	// "encoding/json"
	"strings"
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
	query := r.URL.Query().Get("query") // user passes it as "harry+potter"
	// this method will transform it into string of "harry potter" with no + sign
	// to preserve the + sign, need the user to pass query url as "harry%2Bpotter"

	// if you want to add the + manually here -- we will for testing purposes
    query = strings.Replace(query, " ", "+", -1)

	// regardless, frontend can manipulate ther URI component to the format we want
	// for now we'll replace the space in query ourselves for the backend

	allResults, err:= searchResults.TMDBSearchMovieByKeywords(query) // will be returned with a formatted JSON of the data we need for search results page

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	// TODO - when implementing frontend, will need to structure json response
	// into a format that the frontend can use to display each search result cleanly on a 
	// search result page based on movie 
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"search_results": allResults})
}

// https://api.themoviedb.org/3/search/movie?query=Harry+Potter 