package services

import (
	"context"
	"time"

	"github.com/kevinpista/my-flick-list/backend/models"

)

type WatchlistItemService struct {
	WatchlistItem models.WatchlistItem
}

// Watchlist Item must always belong to a Watchlist
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


