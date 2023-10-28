package tmdb_services

import (
	"encoding/json"
	"net/http"

	"github.com/kevinpista/my-flick-list/backend/models"
)

type TMDBMovieSearchService struct {
	MovieSearch models.TMDBMovieSearch
}

func (c *TMDBMovieSearchService) TMDBSearchMovieByKeywords(query string) (*[]models.TMDBMovieSearch, error) {
	apiUrl := baseAPIUrl + query + "&api_key=" + APIKey
	// Send GET request to TMDB
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the JSON response into a response struct
	var response struct {
		Results []models.TMDBMovieSearch
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response.Results, nil
}
