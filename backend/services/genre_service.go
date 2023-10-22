package services

import (
	"context"
	"database/sql"
)

type Genre struct {
	MovieID int    `json:"movie_id"`
	GenreID int    `json:"genre_id"`
	Genre   string `json:"genre"`
}

// TODO - Genre data is taken from 3rd party data base. Handle case where 3rd party DB does not have genre data
func (c *Genre) GetGenreByMovieID(id int) (*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT movie_id, genre_id, genre FROM genre
		WHERE movie_id = $1
	`
	row, err := db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		var genre Genre
		err = row.Scan(
			&genre.MovieID,
			&genre.GenreID,
			&genre.Genre,
		)
		if err != nil {
			return nil, err
		}
		return &genre, nil
	}
	return nil, sql.ErrNoRows // Case where query did not find any genre data for the movie

}

func (c *Genre) CreateGenreDataByMovieID(genre Genre) (*Genre, error) {
	/// movie id passed via the json body
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO genre (movie_id, genre_id, genre)
		VALUES ($1, $2, $3) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		genre.MovieID,
		genre.GenreID,
		genre.Genre,
	)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}
