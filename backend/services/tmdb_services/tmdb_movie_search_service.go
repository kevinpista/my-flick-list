package tmdb_services

import (
	"encoding/json"
	"net/http"
	"errors"

	
	"github.com/kevinpista/my-flick-list/backend/helpers"
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

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New("TMDB API is unavailable at this time")
	}

	if resp.StatusCode != http.StatusOK {
		// Catch all for TMDB API error
		helpers.MessageLogs.ErrorLog.Println("eEror with TMDB API")
		return nil, errors.New("error with TMDB API")
	}

	// Decode the JSON response into a response struct
	var response struct {
		Results []models.TMDBMovieSearch
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error related to decoding successful TMDB response")
		return nil, err
	}

	return &response.Results, nil
}
