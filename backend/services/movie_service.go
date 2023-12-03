package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/kevinpista/my-flick-list/backend/models"
)

type MovieService struct {
	Movie models.Movie
}

func (c *MovieService) GetMovieByID(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, original_title, overview, tagline, release_date, poster_path, backdrop_path, runtime, adult, 
		budget, revenue, rating, votes, created_at, updated_at FROM movie
		WHERE id = $1
	`

	row, err := db.QueryContext(ctx, query, id) // ctx is the state of the db, pass in our query
	if err != nil{
		return nil, err
	}
	
	defer row.Close()

	if row.Next() {
		var movie models.Movie
		err = row.Scan(
			&movie.ID,
			&movie.OriginalTitle,
			&movie.Overview,
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
		return &movie, nil
	}
	return nil, sql.ErrNoRows // Case where query did not find any matching rows
}

func (c *MovieService) GetAllMovies() ([]*models.Movie, error) {
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

	var movies []*models.Movie // holds multiple movie pointers. a slice called 'movies' holding pointers of type Movie struct
	for rows.Next() {   // for every row we get from our db query
		var movie models.Movie // we create a var called movie with type Movie struct and append it to our movies slice
		// order should follow the order of your query
		// scan each row from our query and assigns the column field data from our query to each movie Movie struct field
		err := rows.Scan(
			&movie.ID,
			&movie.OriginalTitle,
			&movie.Overview,
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

func (c *MovieService) CreateMovie(movie models.Movie) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO movie (id, original_title, overview, tagline, release_date, poster_path, backdrop_path, runtime, adult, budget, revenue, rating, votes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) returning *
	`
	_, err := db.ExecContext(
		ctx,
		query,
		movie.ID,
		movie.OriginalTitle,
		movie.Overview,
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
	return &movie, nil
}

// Note currently the CreatedAt & UpdatedAt is default 0:00 time. Did not scan database after it was generated to return in JSON body
