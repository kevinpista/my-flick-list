package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/kevinpista/my-flick-list/backend/models"
	"github.com/google/uuid"
)

type WatchlistService struct {
	Watchlist models.Watchlist
}

func (c *WatchlistService) CreateWatchlist(userID uuid.UUID, watchlist models.Watchlist) (*models.Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO watchlist (users_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) returning id, users_id, name, description, created_at, updated_at
	`
	var insertedWatchlist models.Watchlist
    queryErr := db.QueryRowContext(
        ctx,
        query,
        userID,
        watchlist.Name,
        watchlist.Description,
        time.Now(),
        time.Now(),
    ).Scan(
		&insertedWatchlist.ID,
        &insertedWatchlist.UserID,
        &insertedWatchlist.Name,
        &insertedWatchlist.Description,
        &insertedWatchlist.CreatedAt,
        &insertedWatchlist.UpdatedAt,
	)
    if queryErr != nil {
        return nil, queryErr
    }
	
	return &insertedWatchlist, nil
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

// TODO possibly remove id from query as we don't need that
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
