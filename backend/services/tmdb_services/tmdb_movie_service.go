package tmdb_services

import (
	"encoding/json"
	"net/http"

	"github.com/kevinpista/my-flick-list/backend/models"
)

type TMDBMovieService struct {
	
}

func (c *TMDBMovieService) TMDBGetMovieByID(query string) (*models.TMDBMovie, error) {
	apiUrl := baseMovieAPIUrl + query + "?api_key=" + APIKey
	// Send GET request to TMDB
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var response models.TMDBMovie
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
