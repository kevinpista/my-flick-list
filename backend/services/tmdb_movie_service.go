package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kevinpista/my-flick-list/backend/models"
)

type TMDBMovieService struct {
	Movie   models.TMDBMovie // Struct used to display info on frontend's individual movie page
	Trailer models.TMDBMovie
}

// GET request to TMDB API. Query is the {movie_id}
// (REDIS CACHED)
func (c *TMDBMovieService) TMDBGetMovieByID(query string) (*models.TMDBMovie, error) {
	// Check Redis Cache
	cachedMovie, cachedErr := c.GetMovieFromCache(query)
	if cachedErr != nil {
		fmt.Println("Warning: Cache GET query failed. Continuing with database query. Error:", cachedErr)
	}

	if cachedMovie != nil {
		return cachedMovie, nil
	}

	tmdbAPIKey := os.Getenv("APIKey")
	apiUrl := baseMovieAPIUrl + query + "?api_key=" + tmdbAPIKey
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
		// Error related to decoding successful TMDB response
		return nil, err
	}

	// Update Redis Cache
	updateErr := c.SetMovieInCache(query, &response)
	if updateErr != nil {
		fmt.Println("Warning: Cache SET query failed. Continuing with returning data. Error:", updateErr)
	}
	return &response, nil
}

// ** IMPORTANT SERVICE FUNC BELOW ** -- Only way for a movie to ever get added into our Postgresql database.
// This function only gets called internally by WatchlistItemServices func (CreateWatchlistItemByWatchlistID) when a user wants to add a movie to their watchlist.
// Responsible for querying TMDB API by the movie_id and then taking the decoding the received movie data and adding it to the database.
// Returns nil to watchlist_item services which gives it the go ahead to attempt to connect a user's watchlist_item to the new locally added movie.
// Received movie data from TMDB API is decoded into a models.TMDBMovieDatabaseEntry struct which was required to accomodate for TMDB's json name
// field difference of 'Rating' and 'Votes'. Also adds genre into 'genre' table.
func (c *TMDBMovieService) TMDBGetMovieByIDAddToLocalDatabase(movieID string) error {
	tmdbAPIKey := os.Getenv("APIKey")
	apiUrl := baseMovieAPIUrl + movieID + "?api_key=" + tmdbAPIKey
	// Send GET request to TMDB
	resp, err := http.Get(apiUrl)
	if err != nil {
		// Error with TMDB individual get request
		return err
	}
	defer resp.Body.Close()

	// Handle any TMDB API request errors
	if resp.StatusCode == http.StatusInternalServerError {
		// Error with TMDB individual get request
		return errors.New("unable to add movie locally. TMDB API offline")
	}

	if resp.StatusCode != http.StatusOK {
		var errorResponse models.TMDBError
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)

		if err != nil {
			// This is a JSON decoding issue related to decoding to TMDBError model
			return errors.New("error decoding TMDB error response")
		}

		// If TMDB API returns a 'success : false' response if any errors
		if !errorResponse.Success {
			return errors.New(errorResponse.StatusMessage)
		}

		// Catch all for TMDB API error for non StatusOK
		return errors.New("error with TMDB API")
	}

	// Decode TMDB API movie data to model. Note it is slightly different from a models.Movie
	var movieDataToBeStored models.TMDBMovieForDatabaseEntry

	err = json.NewDecoder(resp.Body).Decode(&movieDataToBeStored)
	if err != nil {
		// Error related to decoding a successful TMDB response
		return err
	}

	// Database service
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Use 'tx' to ensure atomicity of operations since 2 tables are being inserted (Genre + Movie)
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	movieInsertQuery := `
		INSERT INTO movie (id, original_title, overview, tagline, release_date, poster_path, backdrop_path, runtime, adult, budget, revenue, rating, votes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	`
	_, err = tx.ExecContext(
		ctx,
		movieInsertQuery,
		movieDataToBeStored.ID,
		movieDataToBeStored.OriginalTitle,
		movieDataToBeStored.Overview,
		movieDataToBeStored.Tagline,
		movieDataToBeStored.ReleaseDate,
		movieDataToBeStored.PosterPath,
		movieDataToBeStored.BackdropPath,
		movieDataToBeStored.Runtime,
		movieDataToBeStored.Adult,
		movieDataToBeStored.Budget,
		movieDataToBeStored.Revenue,
		movieDataToBeStored.Rating,
		movieDataToBeStored.Votes,
		time.Now(), // movie.CreatedAt
		time.Now(), // movie.UpdatedAt
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	// Update genre table with every genre item in the Genres struct.
	// Insert the genres into the 'genre' table
	genreInsertQuery := `
        INSERT INTO genre (movie_id, genre_id, genre)
        VALUES ($1, $2, $3)
    `
	// Loop through the Genre array provided by TMDB API. Create a row in 'genre' table for every genre
	// the particular movie is categorized under
	for _, genre := range movieDataToBeStored.Genres {
		_, err := tx.ExecContext(
			ctx,
			genreInsertQuery,
			movieDataToBeStored.ID,
			genre.ID,
			genre.Name,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GET request to TMDB API. Query is the {movie_id} - returns the YouTube video ID which is stored as "key" in the model
func (c *TMDBMovieService) TMDBGetMovieTrailerByID(query string) (*string, error) {
	tmdbAPIKey := os.Getenv("APIKey")
	apiUrl := baseMovieAPIUrl + query + "/videos" + "?api_key=" + tmdbAPIKey

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

	dec := json.NewDecoder(resp.Body)
	var trailerData struct { // Response struct to store results into a array as TMDB json response is set up the same way
		Results []models.TMDBMovieTrailer `json:"results"`
	}

	// Decode the response
	if err := dec.Decode(&trailerData); err != nil {
		return nil, errors.New("error decoding TMDB trailer response")
	}

	// Filter for older trailer as a movie may have more than 1
	var oldestMatchingTrailer *models.TMDBMovieTrailer
	oldestPublishedAt := time.Now() // Initialize a max time value (almost like infinity)

	// Filter trailers based on criteria and find oldest match
	for i := range trailerData.Results {
		trailer := &trailerData.Results[i]
		if trailer.Site == "YouTube" && trailer.Type == "Trailer" && trailer.PublishedAt.Before(oldestPublishedAt) {
			oldestPublishedAt = trailer.PublishedAt
			oldestMatchingTrailer = trailer
		}
	}

	// Check if a matching trailer was found, else return nil for frontend to handle
	if oldestMatchingTrailer == nil {
		return nil, nil
	}

	// Return desired data from the oldest matching trailer
	return &oldestMatchingTrailer.Key, nil
}
