package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Watchlist struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"users_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TODO add member_id, ignore for now as USERS resource has not been implemented yet
func (c *Watchlist) CreateWatchlist(watchlist Watchlist) (*Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		INSERT INTO watchlist (name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4) returning *
	`
	_, err := db.ExecContext(
		ctx,
		query,
		watchlist.Name,
		watchlist.Description,
		time.Now(), // watchlist.CreatedAt
		time.Now(), // watchlist.UpdatedAt
	)

	if err != nil {
		return nil, err
	}
	return &watchlist, nil
}

func (c *Watchlist) GetAllWatchlists() ([]*Watchlist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
	SELECT id, users_id, name, description, created_at, updated_at FROM watchlist
	`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var watchlists []*Watchlist
	for rows.Next() {
		var watchlist Watchlist
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
func (c *Watchlist) GetWatchlistByID(id int) (*Watchlist, error) {
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
		var watchlist Watchlist
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
