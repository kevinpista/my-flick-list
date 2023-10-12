package services

import (
	"time"
	"context"
	// "github.com/google/uuid"
)

type Movie struct{
	ID uint32 `json:"id"`
	OriginalTitle string `json:"original_title"`
	Overview string `json:"overview"`
	Tagline string `json:"tagline"`
	ReleaseDate time.Time `json:"release_date"`
	PosterPath string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	Runtime uint16 `json:"runtime"`
	Adult bool `json:"adult"`
	Budget uint32 `json:"budget"`
	Revenue uint64 `json:"revenue"`
	Rating float32 `json:"rating"`
	Votes uint32 `json:"votes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c * Movie) GetAllMovies() ([]*Movie, error) {
// point to our movie struct, returning a slice of our movie struct (slice of pointers) and also an error.
// ctx is an instance of the context.Context type. provides a way to carry deadlinesm, cancellations, and other request-scoped values
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout) // we set a timeout of val dbTimeout. if query doesn't complete in time it cancels the query
	defer cancel()

	query := `
		SELECT id, original_title, overview, tagline, release_date, poster_path, backdrop_path, runtime, adult, budget, revenue, rating, votes, created_at, updated_at FROM movie
		`
	rows, err := db.QueryContext(ctx, query) // ctx is the state of the db, pass in our query

	if err != nil {
		return nil, err
	}

	var movies []*Movie // holds multiple movie pointers. a slice called 'movies' holding pointers of type Movie struct
	for rows.Next() { // for every row we get from our db query
		var movie Movie	// we create a var called movie with type Movie struct and append it to our movies slice
		// order should follow the order of your query
		err := rows.Scan(
			&movie.ID,
			&movie.OriginalTitle,
			&movie.Tagline,
			&movie.ReleaseDate,
			&movie.PosterPath,
			&movie.BackdropPath,
			&movie.Runtime,
			&movie.Adult,
			&movie.Budget,
			&movie.Revenue,
			&movie.Rating,
			&movie.Votes,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (c *Movie) CreateMovie(movie Movie) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

    // use $ placeholders values to safely pass data from code to SQL queries, preventing SQL injection attacks
	query := `
		INSERT INTO movie (id, original_title, overview, tagline, release_date, poster_path, backdrop_path, runtime, adult, budget, revenue, rating, votes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) return *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		movie.ID,
		movie.OriginalTitle,
		movie.Tagline,
		movie.ReleaseDate,
		movie.PosterPath,
		movie.BackdropPath,
		movie.Runtime,
		movie.Adult,
		movie.Budget,
		movie.Revenue,
		movie.Rating,
		movie.Votes,
		time.Now(), // movie.CreatedAt
		time.Now(), // movie.UpdatedAt
	)

	if err != nil {
		return nil, err
	}
	
	return &movie , nil // not the movie we'll store db, just the info we'll use to create the movie
}