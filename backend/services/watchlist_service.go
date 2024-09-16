package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
)

type WatchlistService struct {
	Watchlist models.Watchlist
}

// Create watchlist. Returns the watchlist details of the newly created watchlist - primarily for frontend to access id
func (c *WatchlistService) CreateWatchlist(userID uuid.UUID, watchlist models.Watchlist) (*models.Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO watchlist (users_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) returning id, users_id, name, description, created_at, updated_at
	`
	var createdWatchlist models.Watchlist
	queryErr := db.QueryRowContext(
		ctx,
		query,
		userID,
		watchlist.Name,
		watchlist.Description,
		time.Now(),
		time.Now(),
	).Scan(
		&createdWatchlist.ID,
		&createdWatchlist.UserID,
		&createdWatchlist.Name,
		&createdWatchlist.Description,
		&createdWatchlist.CreatedAt,
		&createdWatchlist.UpdatedAt,
	)
	if queryErr != nil {
		return nil, queryErr
	}

	// Delete User's All Watchlists Cache if any
	err := c.DeleteAllWatchlistsFromCache(userID)
	if err != nil {
		fmt.Println("Warning: Cache DELETE query failed. Continuing with returning data. Error:", err)
	}

	return &createdWatchlist, nil
}

func (c *WatchlistService) GetAllWatchlists() ([]*models.Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	SELECT id, users_id, name, description, created_at, updated_at FROM watchlist
	`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var watchlists []*models.Watchlist
	for rows.Next() {
		var watchlist models.Watchlist
		err := rows.Scan(
			&watchlist.ID,
			&watchlist.UserID,
			&watchlist.Name,
			&watchlist.Description,
			&watchlist.CreatedAt,
			&watchlist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		watchlists = append(watchlists, &watchlist)
	}

	return watchlists, nil
}

func (c *WatchlistService) GetAllWatchlistsByUserID(userID uuid.UUID) ([]*models.Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT id, users_id, name, description, created_at, updated_at FROM watchlist
		WHERE users_id = $1
	`

	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var watchlists []*models.Watchlist
	for rows.Next() {
		var watchlist models.Watchlist
		err := rows.Scan(
			&watchlist.ID,
			&watchlist.UserID,
			&watchlist.Name,
			&watchlist.Description,
			&watchlist.CreatedAt,
			&watchlist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		watchlists = append(watchlists, &watchlist)
	}
	return watchlists, nil
}

// Gets all watchlist data and returns count
// (REDIS CACHED)
func (c *WatchlistService) GetWatchlistsByUserIDWithMovieCount(userID uuid.UUID) ([]*models.WatchlistWithItemCount, error) {
	// Check Redis cache first
	cachedWatchlists, err := c.GetAllWatchlistsFromCache(userID)
	if err != nil {
		fmt.Println("Warning: Cache GET query failed. Continuing with database query. Error:", err)
	}

	if cachedWatchlists != nil {
		return cachedWatchlists, nil
	}
	// Cache miss: fetch from database
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT w.id,
			w.name,
			w.description,
			COUNT(wi.id) AS watchlist_item_count,
			w.created_at,
			w.updated_at
		FROM watchlist w
		LEFT JOIN watchlist_item wi ON w.id = wi.watchlist_id
		WHERE w.users_id = $1
		GROUP BY w.id, w.name, w.description, w.created_at, w.updated_at;
	`

	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var watchlists []*models.WatchlistWithItemCount
	for rows.Next() {
		var watchlist models.WatchlistWithItemCount
		err := rows.Scan(
			&watchlist.ID,
			&watchlist.Name,
			&watchlist.Description,
			&watchlist.WatchlistItemCount,
			&watchlist.CreatedAt,
			&watchlist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		watchlists = append(watchlists, &watchlist)
	}

	// Update Redis Cache
	err = c.SetAllWatchlistsInCache(userID, watchlists)
	if err != nil {
		fmt.Println("Warning: Cache SET query failed. Continuing with returning data. Error:", err)
	}
	return watchlists, nil
}

// Fetches all watchlists belonging to a user. Takes a movieID parameter.
// Returns all watchlists with the count of all watchlist_items belonging to each watchlist
// and also a boolean of "contains_movie" on whether or not there are any watchlists that
// have a watchlist item that contain the movieID parameter
func (c *WatchlistService) GetWatchlistsByUserIDWithMovieIDCheck(userID uuid.UUID, movieID int) ([]*models.WatchlistWithCountAndContainsMovie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT w.id, w.name,
			COUNT(wi.id) AS watchlist_item_count,
			COUNT(wi.movie_id = $2 OR NULL) > 0 AS contains_queried_movie
		FROM watchlist w
		LEFT JOIN watchlist_item wi ON w.id = wi.watchlist_id
		WHERE w.users_id = $1
		GROUP BY w.id
	`

	rows, err := db.QueryContext(ctx, query, userID, movieID)
	if err != nil {
		return nil, err
	}

	var watchlists []*models.WatchlistWithCountAndContainsMovie
	for rows.Next() {
		var watchlist models.WatchlistWithCountAndContainsMovie
		err := rows.Scan(
			&watchlist.ID,
			&watchlist.Name,
			&watchlist.WatchlistItemCount,
			&watchlist.ContainsQueriedMovie,
		)
		if err != nil {
			return nil, err
		}
		watchlists = append(watchlists, &watchlist)
	}
	return watchlists, nil
}

func (c *WatchlistService) GetWatchlistByID(id int) (*models.Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT id, users_id, name, description, created_at, updated_at FROM watchlist
		WHERE id = $1
	`

	row, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		var watchlist models.Watchlist
		err = row.Scan(
			&watchlist.ID,
			&watchlist.UserID,
			&watchlist.Name,
			&watchlist.Description,
			&watchlist.CreatedAt,
			&watchlist.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &watchlist, nil
	} else {
		return nil, sql.ErrNoRows // Case where query did not find any matching rows
	}

}

// Deletes a watchlist with its id - deletes all associated watchlist_items first, then the watchlist
func (c *WatchlistService) DeleteWatchlistByID(watchlistID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Delete associated watchlist_item_notes first
	_, err = tx.ExecContext(ctx, "DELETE FROM watchlist_item_note WHERE watchlist_item_id IN (SELECT id FROM watchlist_item WHERE watchlist_id = $1)", watchlistID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete associated watchlist_items
	_, err = tx.ExecContext(ctx, "DELETE FROM watchlist_item WHERE watchlist_id = $1", watchlistID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Delete watchlist itself
	_, err = tx.ExecContext(ctx, "DELETE FROM watchlist WHERE id = $1", watchlistID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Get's the UUID user_id owner of the watchlist
func (c *WatchlistService) GetWatchlistOwnerUserID(watchlistID int) (uuid.UUID, error) {
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

// Updates name of Watchlist. Returns the watchlist name for the frontend
func (c *WatchlistService) UpdateWatchlistName(watchlistID int, watchlist models.Watchlist) (*models.Watchlist, error) {

	// Update the timestamp for the updated_at field
	watchlist.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE watchlist
		SET name = $1, updated_at = $2
		WHERE id = $3
		RETURNING name
    `
	var updatedWatchlist models.Watchlist
	err := db.QueryRowContext(
		ctx,
		query,
		watchlist.Name,
		watchlist.UpdatedAt,
		watchlistID,
	).Scan(
		&updatedWatchlist.Name,
	)
	// populates model's name only
	if err != nil {
		return nil, err
	}

	userID, err := c.GetWatchlistOwnerUserID(watchlistID)
	if err != nil {
		fmt.Println("Failed to get watchlist owner userID from watchlist ID")
	} else {
		delErr := c.DeleteAllWatchlistsFromCache(userID)
		if delErr != nil {
			fmt.Println("Warning: Cache DELETE query failed. Continuing with returning data. Error:", delErr)
		}
	}

	return &updatedWatchlist, nil
}

// Updates description of Watchlist. Returns the watchlist description for the frontend
func (c *WatchlistService) UpdateWatchlistDescription(watchlistID int, watchlist models.Watchlist) (*models.Watchlist, error) {

	// Update the timestamp for the updated_at field
	watchlist.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE watchlist
		SET description = $1, updated_at = $2
		WHERE id = $3
		RETURNING description
    `
	var updatedWatchlist models.Watchlist
	err := db.QueryRowContext(
		ctx,
		query,
		watchlist.Description,
		watchlist.UpdatedAt,
		watchlistID,
	).Scan(
		&updatedWatchlist.Description,
	)
	// Populates model's description only
	if err != nil {
		return nil, err
	}

	userID, err := c.GetWatchlistOwnerUserID(watchlistID)
	if err != nil {
		fmt.Println("Failed to get watchlist owner userID from watchlist ID")
	} else {
		delErr := c.DeleteAllWatchlistsFromCache(userID)
		if delErr != nil {
			fmt.Println("Warning: Cache DELETE query failed. Continuing with returning data. Error:", delErr)
		}
	}
	
	return &updatedWatchlist, nil
}
