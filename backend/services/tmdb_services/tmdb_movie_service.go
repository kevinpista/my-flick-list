package tmdb_services

import (
	"encoding/json"
	"net/http"
	"errors"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/models"
)

type TMDBMovieService struct {
	Movie models.TMDBMovie
}
// GET request to TMDB API. Query is the {movie_id}
func (c *TMDBMovieService) TMDBGetMovieByID(query string) (*models.TMDBMovie, error) {
	apiUrl := baseMovieAPIUrl + query + "?api_key=" + APIKey
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
		var errorResponse models.TMDBError
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)

		if err != nil {
			// This is a JSON decoding issue related to decoding to TMDBError model
			return nil, errors.New("error decoding TMDB error response")
		}

		// TMDB API returns a 'success : false' response if any errors
		if !errorResponse.Success {
			return nil, errors.New(errorResponse.StatusMessage)
		}

		// Catch all for TMDB API error for non StatusOK
		return nil, errors.New("error with TMDB API")
	}

	var response models.TMDBMovie
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("error related to decoding successful TMDB response")
		return nil, err
	}

	return &response, nil
}
