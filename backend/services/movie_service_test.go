package services

import (
	"testing"

	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/kevinpista/my-flick-list/backend/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMovieByID(t *testing.T) {
	movieService := MovieService{}

	// Create and insert test movie first
	testMovie := models.Movie{
		ID:            util.RandomInt32(9999991, 99999992),
		OriginalTitle: util.RandomName(),
		Overview:      util.RandomParagraph(),
		Tagline:       "Tag line test movie",
		ReleaseDate:   "2023-10-21",
		PosterPath:    "abcdefg",
		BackdropPath:  "abcdefg",
		Runtime:       100,
		Adult:         false,
		Budget:        99999992,
		Revenue:       99999999,
		Rating:        9,
		Votes:         99999,
	}

	createMovieResponse, err := movieService.CreateMovie(testMovie)

	require.NoError(t, err, "CreateMovie should return no error")
	require.Nil(t, err, "CreateMovie err should be nil")
	require.NotNil(t, createMovieResponse, "CreateMovie response should not be nil")
	require.IsType(t, createMovieResponse, &models.Movie{}, "CreateMovie response should be type: Movie")

	getMovieByIDResponse, err := movieService.GetMovieByID(testMovie.ID)
	assert.NoError(t, err, "GetMovieByID should return no error")
	assert.Nil(t, err, "GetMovieByID err should be nil")
	assert.NotNil(t, getMovieByIDResponse, "GetMovieByID response should not be nil")
	assert.IsType(t, getMovieByIDResponse, &models.Movie{}, "GetMovieByID response should be type: Movie")

	assert.Equal(t, getMovieByIDResponse.ID, testMovie.ID, "GetMovieByID response ID should equal input ID")
	assert.Equal(t, getMovieByIDResponse.OriginalTitle, testMovie.OriginalTitle, "GetMovieByID response OriginalTitle should equal input OriginalTitle")
	assert.Equal(t, getMovieByIDResponse.Overview, testMovie.Overview, "GetMovieByID response Overview should equal input Overview")
	assert.Equal(t, getMovieByIDResponse.Tagline, testMovie.Tagline, "GetMovieByID response Tagline should equal input Tagline")
	assert.Equal(t, getMovieByIDResponse.PosterPath, testMovie.PosterPath, "GetMovieByID response PosterPath should equal input PosterPath")
	assert.Equal(t, getMovieByIDResponse.BackdropPath, testMovie.BackdropPath, "GetMovieByID response BackdropPath should equal input BackdropPath")
	assert.Equal(t, getMovieByIDResponse.Runtime, testMovie.Runtime, "GetMovieByID response Runtime should equal input Runtime")
	assert.Equal(t, getMovieByIDResponse.Adult, testMovie.Adult, "GetMovieByID response Adult should equal input Adult")
	assert.Equal(t, getMovieByIDResponse.Budget, testMovie.Budget, "GetMovieByID response Budget should equal input Budget")
	assert.Equal(t, getMovieByIDResponse.Revenue, testMovie.Revenue, "GetMovieByID response Revenue should equal input Revenue")
	assert.Equal(t, getMovieByIDResponse.Rating, testMovie.Rating, "GetMovieByID response Rating should equal input Rating")
	assert.Equal(t, getMovieByIDResponse.Votes, testMovie.Votes, "GetMovieByID response Votes should equal input Votes")
}

func TestGetAllMovies(t *testing.T) {
	movieService := MovieService{}

	getAllMoviesResponse, err := movieService.GetAllMovies()
	assert.NoError(t, err, "GetAllMovies should return no error")
	assert.Nil(t, err, "GetAllMovies err should be nil")
	assert.NotNil(t, getAllMoviesResponse, "GetAllMovies response should not be nil")
	assert.IsType(t, getAllMoviesResponse, []*models.Movie{}, "GetAllMovies response should be type: []*Movie")
}

func TestCreateMovie(t *testing.T) {
	movieService := MovieService{}

	testMovie := models.Movie{
		ID:            util.RandomInt32(9999991, 99999992),
		OriginalTitle: util.RandomName(),
		Overview:      util.RandomParagraph(),
		Tagline:       "Tag line test movie",
		ReleaseDate:   "2023-10-21",
		PosterPath:    "abcdefg",
		BackdropPath:  "abcdefg",
		Runtime:       100,
		Adult:         false,
		Budget:        99999992,
		Revenue:       99999999,
		Rating:        9,
		Votes:         99999,
	}

	createMovieResponse, err := movieService.CreateMovie(testMovie)

	assert.NoError(t, err, "CreateMovie should return no error")
	assert.Nil(t, err, "CreateMovie err should be nil")
	assert.NotNil(t, createMovieResponse, "CreateMovie response should not be nil")
	assert.IsType(t, createMovieResponse, &models.Movie{}, "CreateMovie response should be type: Movie")

	assert.Equal(t, createMovieResponse.ID, testMovie.ID, "CreateMovie response ID should equal input ID")
	assert.Equal(t, createMovieResponse.OriginalTitle, testMovie.OriginalTitle, "CreateMovie response OriginalTitle should equal input OriginalTitle")
	assert.Equal(t, createMovieResponse.Overview, testMovie.Overview, "CreateMovie response Overview should equal input Overview")
	assert.Equal(t, createMovieResponse.Tagline, testMovie.Tagline, "CreateMovie response Tagline should equal input Tagline")
	assert.Equal(t, createMovieResponse.ReleaseDate, testMovie.ReleaseDate, "CreateMovie response ReleaseDate should equal input ReleaseDate")
	assert.Equal(t, createMovieResponse.PosterPath, testMovie.PosterPath, "CreateMovie response PosterPath should equal input PosterPath")
	assert.Equal(t, createMovieResponse.BackdropPath, testMovie.BackdropPath, "CreateMovie response BackdropPath should equal input BackdropPath")
	assert.Equal(t, createMovieResponse.Runtime, testMovie.Runtime, "CreateMovie response Runtime should equal input Runtime")
	assert.Equal(t, createMovieResponse.Adult, testMovie.Adult, "CreateMovie response Adult should equal input Adult")
	assert.Equal(t, createMovieResponse.Budget, testMovie.Budget, "CreateMovie response Budget should equal input Budget")
	assert.Equal(t, createMovieResponse.Revenue, testMovie.Revenue, "CreateMovie response Revenue should equal input Revenue")
	assert.Equal(t, createMovieResponse.Rating, testMovie.Rating, "CreateMovie response Rating should equal input Rating")
	assert.Equal(t, createMovieResponse.Votes, testMovie.Votes, "CreateMovie response Votes should equal input Votes")
}

/*
type Movie struct {
	ID            int       `json:"id"`
	OriginalTitle string    `json:"original_title"`
	Overview      string    `json:"overview"`
	Tagline       string    `json:"tagline"`
	ReleaseDate   string    `json:"release_date"`
	PosterPath    string    `json:"poster_path"`
	BackdropPath  string    `json:"backdrop_path"`
	Runtime       uint16    `json:"runtime"`
	Adult         bool      `json:"adult"`
	Budget        uint32    `json:"budget"`
	Revenue       uint64    `json:"revenue"`
	Rating        float32   `json:"rating"`
	Votes         uint32    `json:"votes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
*/
