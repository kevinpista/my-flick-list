package models

import (
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

// Model for the dialog drop down menu for "Add to Watchlist" button on Movie page
type WatchlistWithCountAndContainsMovie struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	WatchlistItemCount   int    `json:"watchlist_item_count"`
	ContainsQueriedMovie bool   `json:"contains_queried_movie"`
}

type WatchlistWithItemCount struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	WatchlistItemCount int       `json:"watchlist_item_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}