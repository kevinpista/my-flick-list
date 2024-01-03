package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
)

type WatchlistItemService struct {
	WatchlistItem          models.WatchlistItem
	WatchlistItemWithMovie models.WatchlistItemWithMovie
}

// Fetches all watchlist items that belongs to a specific watchlist via its watchlistID
func (c *WatchlistItemService) GetAllWatchlistItemsByWatchlistID(watchlistID int) ([]*models.WatchlistItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT id, watchlist_id, movie_id, checkmarked, created_at, updated_at FROM watchlist_item 
		WHERE watchlist_id = $1
	`
	rows, err := db.QueryContext(ctx, query, watchlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var watchlistItems []*models.WatchlistItem
	for rows.Next() {
		var watchlistItem models.WatchlistItem
		err := rows.Scan(
			&watchlistItem.ID,
			&watchlistItem.WatchlistID,
			&watchlistItem.MovieID,
			&watchlistItem.Checkmarked,
			&watchlistItem.CreatedAt,
			&watchlistItem.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		watchlistItems = append(watchlistItems, &watchlistItem)
	}
	return watchlistItems, nil
}

// Fetches watchlist name & description + all watchlist items and its movie data
func (c *WatchlistItemService) GetWatchListWithWatchlistItemsByWatchListID(watchlistID int) ([]*models.WatchlistItemWithMovie, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

    // Fetch watchlist name and description
    var watchlistName, watchlistDescription string
    err := db.QueryRowContext(ctx, "SELECT name, description FROM watchlist WHERE id = $1", watchlistID).Scan(&watchlistName, &watchlistDescription)
    if err != nil {
        return nil, "", "", err
    }

	// Fetch watchlist items belonging to the watchlist
	query := `
		SELECT 
			wi.id, wi.watchlist_id, wi.movie_id, wi.checkmarked, wi.created_at, wi.updated_at,
			m.original_title, m.overview, m.tagline, m.release_date, m.poster_path, m.backdrop_path, m.runtime, m.adult,
			m.budget, m.revenue, m.rating, m.votes, m.created_at AS movie_created_at, m.updated_at AS movie_updated_at
		FROM watchlist_item wi
		JOIN movie m ON wi.movie_id = m.id
		JOIN watchlist w ON wi.watchlist_id = w.id
		WHERE wi.watchlist_id = $1
	`

    rows, err := db.QueryContext(ctx, query, watchlistID)
    if err != nil {
        return nil, "", "", err
    }
    defer rows.Close()

	var watchlistItemsWithMovies []*models.WatchlistItemWithMovie
	for rows.Next() {
		var watchlistItemWithMovie models.WatchlistItemWithMovie
		err := rows.Scan(
			&watchlistItemWithMovie.ID,
			&watchlistItemWithMovie.WatchlistID,
			&watchlistItemWithMovie.MovieID,
			&watchlistItemWithMovie.Checkmarked,
			&watchlistItemWithMovie.CreatedAt,
			&watchlistItemWithMovie.UpdatedAt,
			&watchlistItemWithMovie.OriginalTitle,
			&watchlistItemWithMovie.Overview,
			&watchlistItemWithMovie.Tagline,
			&watchlistItemWithMovie.ReleaseDate,
			&watchlistItemWithMovie.PosterPath,
			&watchlistItemWithMovie.BackdropPath,
			&watchlistItemWithMovie.Runtime,
			&watchlistItemWithMovie.Adult,
			&watchlistItemWithMovie.Budget,
			&watchlistItemWithMovie.Revenue,
			&watchlistItemWithMovie.Rating,
			&watchlistItemWithMovie.Votes,
			&watchlistItemWithMovie.MovieCreatedAt,
			&watchlistItemWithMovie.MovieUpdatedAt,
		)

        if err != nil {
            return nil, "", "", err
        }
		watchlistItemsWithMovies = append(watchlistItemsWithMovies, &watchlistItemWithMovie)
	}
    return watchlistItemsWithMovies, watchlistName, watchlistDescription, nil
}

/*
// Fetches all watchlist items + the associated movie data only
func (c *WatchlistItemService) GetAllWatchlistItemsWithMoviesByWatchListID(watchlistID int) ([]*models.WatchlistItemWithMovie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT wi.id, wi.watchlist_id, wi.movie_id, wi.checkmarked, wi.created_at, wi.updated_at,
			m.original_title, m.overview, m.tagline, m.release_date, m.poster_path, m.backdrop_path, m.runtime, m.adult,
			m.budget, m.revenue, m.rating, m.votes, m.created_at AS movie_created_at, m.updated_at AS movie_updated_at
		FROM watchlist_item wi
		JOIN movie m ON wi.movie_id = m.id
		WHERE wi.watchlist_id = $1
	`

	rows, err := db.QueryContext(ctx, query, watchlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var watchlistItemsWithMovies []*models.WatchlistItemWithMovie
	for rows.Next() {
		var watchlistItemWithMovie models.WatchlistItemWithMovie
		err := rows.Scan(
			&watchlistItemWithMovie.ID,
			&watchlistItemWithMovie.WatchlistID,
			&watchlistItemWithMovie.MovieID,
			&watchlistItemWithMovie.Checkmarked,
			&watchlistItemWithMovie.CreatedAt,
			&watchlistItemWithMovie.UpdatedAt,
			&watchlistItemWithMovie.OriginalTitle,
			&watchlistItemWithMovie.Overview,
			&watchlistItemWithMovie.Tagline,
			&watchlistItemWithMovie.ReleaseDate,
			&watchlistItemWithMovie.PosterPath,
			&watchlistItemWithMovie.BackdropPath,
			&watchlistItemWithMovie.Runtime,
			&watchlistItemWithMovie.Adult,
			&watchlistItemWithMovie.Budget,
			&watchlistItemWithMovie.Revenue,
			&watchlistItemWithMovie.Rating,
			&watchlistItemWithMovie.Votes,
			&watchlistItemWithMovie.MovieCreatedAt,
			&watchlistItemWithMovie.MovieUpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		watchlistItemsWithMovies = append(watchlistItemsWithMovies, &watchlistItemWithMovie)
	}
	return watchlistItemsWithMovies, nil
}
*/

// Creates a watchlist_item with a movie_id with the watchlist ID it belongs to
func (c *WatchlistItemService) CreateWatchlistItemByWatchlistID(watchlistItem models.WatchlistItem) (*models.WatchlistItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO watchlist_item (watchlist_id, movie_id, checkmarked, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) returning *
		`
	_, err := db.ExecContext(
		ctx,
		query,
		watchlistItem.WatchlistID,
		watchlistItem.MovieID,
		watchlistItem.Checkmarked,
		time.Now(), // watchlistItem.CreatedAt
		time.Now(), // watchlistItem.UpdatedAt
	)
	if err != nil {
		return nil, err
	}
	return &watchlistItem, nil
}

// Deletes a watchlist_item with its id
func (c *WatchlistItemService) DeleteWatchlistItemByID(watchlistItemID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		DELETE FROM watchlist_item WHERE id = $1
	`
	
	_, err := db.ExecContext(ctx, query, watchlistItemID)
	if err != nil {
		return err
	}

	return nil
}


// Updates the "checkmarked" boolean of the particular watchlist_item
func (c *WatchlistItemService) UpdateCheckmarkedBooleanByWatchlistItemByID(watchlistItemID int, watchlistItem models.WatchlistItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		UPDATE watchlist_item
		SET checkmarked = $1
		WHERE id = $2
		`
	_, err := db.ExecContext(
		ctx,
		query,
		watchlistItem.Checkmarked,
		watchlistItemID,
	)
	if err != nil {
		return err
	}
	return nil
}

/*
HELPER FUNCTIONS
*/

// Check if exisiting watchlist_items within a watchlist contain a movie_id as the one being passed by user
func (c *WatchlistItemService) CheckIfMovieInWatchlistExists(watchlistID, movieID int) (bool, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
		SELECT EXISTS (SELECT 1 FROM watchlist_item WHERE watchlist_id = $1 AND movie_id = $2)
    `
    var exists bool
    err := db.QueryRowContext(ctx, query, watchlistID, movieID).Scan(&exists)
    if err != nil {
        return false, err
    }

    return exists, nil
}

// Checks if a watchlist with the given ID exists
func (c *WatchlistItemService) CheckIfWatchlistExists(watchlistID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM watchlist WHERE id = $1)
		`
	var exists bool
	err := db.QueryRowContext(ctx, query, watchlistID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Get's the user_id owner of the watchlist
func (c *WatchlistItemService) GetWatchlistOwnerUserID(watchlistID int) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT users_id FROM watchlist WHERE id = $1"
	var watchlistOwnerID uuid.UUID
	err := db.QueryRowContext(ctx, query, watchlistID).Scan(&watchlistOwnerID)
	if err != nil {
		return uuid.Nil, err
	}

	return watchlistOwnerID, nil
}

// Get's the watchlist_id of the watchlist_item
func (c *WatchlistItemService) GetWatchlistItemWatchlistId(watchlistItemID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT watchlist_id FROM watchlist_item WHERE id = $1"
	var watchlist_id int
	err := db.QueryRowContext(ctx, query, watchlistItemID).Scan(&watchlist_id)
	if err != nil {
		return 0, err
	}

	return watchlist_id, nil
}