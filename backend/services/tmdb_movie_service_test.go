package services

import (
	"testing"

	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestTMDBGetMovieByID(t *testing.T) {
	tmdbMovieService := TMDBMovieService{}

	movieIDQuery:= "1895"

	tmdbGetMovieByIDResponse, err := tmdbMovieService.TMDBGetMovieByID(movieIDQuery)

	assert.NoError(t, err, "TMDBGetMovieByID should return no error")
	assert.Nil(t, err, "TMDBGetMovieByID err should be nil")
	assert.NotNil(t, tmdbGetMovieByIDResponse, "TMDBGetMovieByID response should not be nil")
	assert.IsType(t, tmdbGetMovieByIDResponse, &models.TMDBMovie{}, "TMDBGetMovieByID response should be of type: TMDBMovie")
	assert.Equal(t, tmdbGetMovieByIDResponse.ID, 1895, "tmdbGetMovieByIDResponse.ID should equal MovieID Input")
}

func TestTMDBGetMovieByIDAddToLocalDatabase(t *testing.T) {
	tmdbMovieService := TMDBMovieService{}

	movieIDQuery:= "1792"

	err := tmdbMovieService.TMDBGetMovieByIDAddToLocalDatabase(movieIDQuery)

	assert.NoError(t, err, "TMDBGetMovieByIDAddToLocalDatabase should return no error")
	assert.Nil(t, err, "TMDBGetMovieByIDAddToLocalDatabase err should be nil")
}

func TestTMDBGetMovieTrailerByID(t *testing.T) {
	tmdbMovieService := TMDBMovieService{}

	movieIDQuery:= "1791"

	tmdbGetMovieTrailerByIDResponse, err := tmdbMovieService.TMDBGetMovieTrailerByID(movieIDQuery)

	assert.NoError(t, err, "TMDBGetMovieTrailerByID should return no error")
	assert.Nil(t, err, "TMDBGetMovieTrailerByID err should be nil")
	assert.NotNil(t, tmdbGetMovieTrailerByIDResponse, "TMDBGetMovieTrailerByID response should not be nil")
}