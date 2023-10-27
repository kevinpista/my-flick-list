package tmdb_services

import (
	"net/http"
	"encoding/json"
	"github.com/kevinpista/my-flick-list/backend/models"
)

var baseAPIUrl = "https://api.themoviedb.org/3/search/movie?query="
var APIKey = "&api_key" // implement API key via .env load

type MovieSearchService struct {
	MovieSearch models.MovieSearch
}


func (c *MovieSearchService) TMDBSearchMovieByKeywords(query string) ([]models.MovieSearch, error) {
	apiUrl := baseAPIUrl + query + APIKey

	// Send GET request to TMDB
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into a response struct
	var response struct {
		Results []models.MovieSearch
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}