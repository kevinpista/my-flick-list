package services

import (
	"encoding/json"
	"net/http"
	"errors"

	
	"github.com/kevinpista/my-flick-list/backend/models"
)

type TMDBMovieSearchService struct {
	MovieSearch models.TMDBMovieSearch
}

func (c *TMDBMovieSearchService) TMDBSearchMovieByKeywords(query string, page string) (*models.TMDBSearchResponse, error) {
	apiUrl := baseAPIUrl + query + "&api_key=" + APIKey + "&page=" + page
	// Send GET request to TMDB
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("TMDB API is unavailable at this time")
	}

	if resp.StatusCode != http.StatusOK {
		// Catch all for TMDB API error
		return nil, errors.New("error with TMDB API")
	}

	// Decode the JSON response into a response struct with pagination
	var response models.TMDBSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		// Error related to decoding successful TMDB response
		return nil, err
	}
	return &response, nil
}
