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

type WatchlistItemWithMovie struct {
	ID             int       `json:"id"`
	WatchlistID    int       `json:"watchlist_id"`
	MovieID        int       `json:"movie_id"`
	Checkmarked    bool      `json:"checkmarked"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	OriginalTitle  string    `json:"original_title"`
	Overview       string    `json:"overview"`
	Tagline        string    `json:"tagline"`
	ReleaseDate    string    `json:"release_date"`
	PosterPath     string    `json:"poster_path"`
	BackdropPath   string    `json:"backdrop_path"`
	Runtime        int       `json:"runtime"`
	Adult          bool      `json:"adult"`
	Budget         int       `json:"budget"`
	Revenue        int       `json:"revenue"`
	Rating         float64   `json:"rating"`
	Votes          int       `json:"votes"`
	MovieCreatedAt time.Time `json:"movie_created_at"`
	MovieUpdatedAt time.Time `json:"movie_updated_at"`
	WatchlistItemNoteFetch // Embedded struct. Its fields can be accessed directly 
}

// Same as 'WatchlistItemNote' from watchlist_item_note.go model file
// except for the field name so it can be used with the 'WatchlistItemWithMovie' model
// in order to return all watchlist items and their watchlist item notes in 1 query
// The "*" signifies these fields are 'nullable' types so database will COALESCE values as null if watchlist_item_note not found
type WatchlistItemNoteFetch struct {
	ItemNotes     *string    `json:"item_notes"`
	NoteCreatedAt *time.Time `json:"note_created_at"`
	NoteUpdatedAt *time.Time `json:"note_updated_at"`
}
