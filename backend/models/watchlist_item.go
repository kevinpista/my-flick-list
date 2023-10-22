package models

import (
	"time"
)

type WatchlistItem struct {
	ID          int       `json:"id"`
	WatchlistID int       `json:"watchlist_id"`
	MovieID     int       `json:"movie_id"`
	Checkmarked bool      `json:"checkmarked"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}