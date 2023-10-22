package models

import (
	"time"
)

type WatchlistItemNote struct {
	WatchlistItemID int       `json:"watchlist_item_id"`
	ItemNotes       string    `json:"item_notes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}