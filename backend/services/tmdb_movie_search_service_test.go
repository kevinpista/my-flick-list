package services

import (
	"testing"

	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestTMDBSearchMovieByKeywords(t *testing.T) {
	tmdbMovieSearchService := TMDBMovieSearchService{}

	query := "potter"
	page := "1"

	tmdbSearchMovieByKeywordsResponse, err := tmdbMovieSearchService.TMDBSearchMovieByKeywords(query, page)

	assert.NoError(t, err, "TMDBSearchMovieByKeywords should return no error")
	assert.Nil(t, err, "TMDBSearchMovieByKeywords err should be nil")
	assert.NotNil(t, tmdbSearchMovieByKeywordsResponse, "TMDBSearchMovieByKeywords response should not be nil")
	assert.IsType(t, tmdbSearchMovieByKeywordsResponse, &models.TMDBSearchResponse{}, "TMDBSearchMovieByKeywords response should be of type: TMDBSearchResponse")
	assert.NotEmpty(t, tmdbSearchMovieByKeywordsResponse.Results, "TMDBSearchResponse.Results should not be empty")
}

/*
type TMDBMovieSearch struct {
	Adult          bool     `json:"adult"`
	BackdropPath   string   `json:"backdrop_path"`
	GenreIds       []int    `json:"genre_ids"`
	ID             int      `json:"id"`
	// OriginalLanguage string `json:"original_language"`
	OriginalTitle  string   `json:"original_title"`
	Overview       string   `json:"overview"`
	Popularity     float64  `json:"popularity"`
	PosterPath     string   `json:"poster_path"`
	ReleaseDate    string   `json:"release_date"`
	// Title          string   `json:"title"`
	// Video          bool     `json:"video"`
	// VoteAverage    float64  `json:"vote_average"`
	// VoteCount      int      `json:"vote_count"`
}

// Include pagination
type TMDBSearchResponse struct {
	Page         int               `json:"page"`
	Results      []TMDBMovieSearch `json:"results"`
	TotalPages   int               `json:"total_pages"`
	TotalResults int               `json:"total_results"`
}
*/
