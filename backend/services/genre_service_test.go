package services

import (
	"testing"

	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/stretchr/testify/assert"
)

func TestGetGenreByMovieID(t *testing.T) {
	genreService := GenreService{}

	getGenreByMovieIDResponse, err := genreService.GetGenreByMovieID(1895)
	assert.NoError(t, err, "GetGenreByMovieID should return no error")
	assert.Nil(t, err, "GetGenreByMovieID err should be nil")
	assert.NotNil(t, getGenreByMovieIDResponse, "GetGenreByMovieID response should not be nil")
	assert.IsType(t, getGenreByMovieIDResponse, &models.Genre{}, "GetGenreByMovieID response should be type: Genre")
}

func TestCreateGenreDataByMovieID(t *testing.T) {
	genreService := GenreService{}

	testGenreData := models.Genre{
		MovieID: 1895,
		GenreID: 28,
		Genre: "Action",
	}

	createGenreDataByMovieIDResponse, err := genreService.CreateGenreDataByMovieID(testGenreData)
	assert.NoError(t, err, "CreateGenreDataByMovieID should return no error")
	assert.Nil(t, err, "CreateGenreDataByMovieID err should be nil")
	assert.NotNil(t, createGenreDataByMovieIDResponse, "CreateGenreDataByMovieID response should not be nil")
	assert.IsType(t, createGenreDataByMovieIDResponse, &models.Genre{}, "CreateGenreDataByMovieID response should be type: Genre")
	assert.Equal(t, createGenreDataByMovieIDResponse.MovieID, testGenreData.MovieID, "createGenreDataByMovieIDResponse MovieID should equal inpurt MovieID")
	assert.Equal(t, createGenreDataByMovieIDResponse.GenreID, testGenreData.GenreID, "createGenreDataByMovieIDResponse GenreID should equal inpurt GenreID")
	assert.Equal(t, createGenreDataByMovieIDResponse.Genre, testGenreData.Genre, "createGenreDataByMovieIDResponse Genre should equal inpurt Genre")
}
