package services

import (
	"context"
	"database/sql"
	"time"
)

type WatchlistItemNote struct {
	WatchlistItemID int       `json:"watchlist_item_id"`
	ItemNotes       string    `json:"item_notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// TODO - handle the case of this, do not want to actually return a real error because having an
// watchlist item note is optional. But probably deal with it in a batch request for watchlist page itself
func (c *WatchlistItemNote) GetWatchlistItemNoteByWatchlistItemID(id int) (*WatchlistItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT watchlist_item_id, item_notes, created_at, updated_at FROM watchlist_item_note
		WHERE watchlist_item_id = $1
	`
	row, err := db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		var watchlistItemNote WatchlistItemNote
		err = row.Scan(
			&watchlistItemNote.WatchlistItemID,
			&watchlistItemNote.ItemNotes,
			&watchlistItemNote.CreatedAt,
			&watchlistItemNote.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &watchlistItemNote, nil
	}
	return nil, sql.ErrNoRows // Case where query did not find any match itemnotes 

}

func (c *WatchlistItemNote) CreateWatchlistItemNote(watchlistItemNote WatchlistItemNote) (*WatchlistItemNote, error) {
	/// watchlist item id passed via the json body
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO watchlist_item_note (watchlist_item_id, item_notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4) returning *
	`

	_, err := db.ExecContext(
		ctx,
		query,
		watchlistItemNote.WatchlistItemID,
		watchlistItemNote.ItemNotes,
		time.Now(), // watchlistItemNote.CreatedAt
		time.Now(), // watchlistItemNote.CreatedAt
	)
	if err != nil{
		return nil, err
	}

	return &watchlistItemNote, nil
}