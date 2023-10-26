package tmdb_services

import (
	"github.com/kevinpista/my-flick-list/backend/models"
	
)

var baseAPIUrl = "https://api.themoviedb.org/3/search/movie?query="

type MovieSearchService struct {
	MovieSearch models.MovieSearch
}

// Return a string for now since TMDB response will be a JSON response of lot of results
func (c *MovieSearchService) TMDBSearchMovieByKeywords(query string) (string, error) {
	apiUrl := baseAPIUrl + query
	return apiUrl, nil
} 