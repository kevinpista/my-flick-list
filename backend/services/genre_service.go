package services

import (
	"context"
	"database/sql"

	"github.com/kevinpista/my-flick-list/backend/models"

)
// custom type that embeds the model of Genre Struct
type GenreService struct {
    Genre models.Genre
}

// TODO - Genre data is taken from 3rd party data base. Handle case where 3rd party DB does not have genre data
func (c *GenreService) GetGenreByMovieID(id int) (*models.Genre, error) {
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
		var genre models.Genre
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

func (c *GenreService) CreateGenreDataByMovieID(genre models.Genre) (*models.Genre, error) {
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
