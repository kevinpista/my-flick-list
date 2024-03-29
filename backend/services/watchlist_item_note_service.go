package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
)

type WatchlistItemNoteService struct {
	WatchlistItemNote models.WatchlistItemNote
}

func (c *WatchlistItemNoteService) GetWatchlistItemNoteByWatchlistItemID(id int) (*models.WatchlistItemNote, error) {
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
		var watchlistItemNote models.WatchlistItemNote
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

func (c *WatchlistItemNoteService) CreateWatchlistItemNote(watchlistItemNote models.WatchlistItemNote) (*models.WatchlistItemNote, error) {
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
		time.Now(), // watchlistItemNote.UpdatedAt
	)
	if err != nil{
		return nil, err
	}
	return &watchlistItemNote, nil
}

// Updates watchlist_item_note text
func (c *WatchlistItemNoteService) UpdateWatchlistItemNote(watchlistItemNote models.WatchlistItemNote) (*models.WatchlistItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		UPDATE watchlist_item_note
		SET item_notes = $1, updated_at = $2
		WHERE watchlist_item_id = $3
		RETURNING watchlist_item_id, item_notes, created_at, updated_at
	`
	var updatedWatchlistItemNote models.WatchlistItemNote

	err := db.QueryRowContext(
		ctx,
		query,
		watchlistItemNote.ItemNotes,
		time.Now(),// watchlistItemNote.UpdatedAt
		watchlistItemNote.WatchlistItemID,
	).Scan(
		&updatedWatchlistItemNote.WatchlistItemID,
		&updatedWatchlistItemNote.ItemNotes,
		&updatedWatchlistItemNote.CreatedAt,
		&updatedWatchlistItemNote.UpdatedAt,
	)

	if err != nil{
		return nil, err
	}

	return &updatedWatchlistItemNote, nil
}

// Fetches all notes in database. Testing purposes only
func (c *WatchlistItemNoteService) GetNotesTest() ([]*models.WatchlistItemNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT watchlist_item_id, item_notes, created_at, updated_at FROM watchlist_item_note
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*models.WatchlistItemNote // holds multiple movie pointers. a slice called 'movies' holding pointers of type Movie struct
	for rows.Next() {   // for every row we get from our db query
		var note models.WatchlistItemNote // we create a var called movie with type Movie struct and append it to our movies slice
		// order should follow the order of your query
		// scan each row from our query and assigns the column field data from our query to each movie Movie struct field
		err := rows.Scan(
			&note.WatchlistItemID,
			&note.ItemNotes,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, &note)
	}

	return notes, nil // Case where query did not find any match itemnotes 
}


/*
HELPER FUNCTIONS
*/

// Checks if watchlist_item exists
func (c *WatchlistItemNoteService) CheckIfWatchlistItemExists(watchListItemID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM watchlist_item WHERE id = $1)
		`
	var exists bool
	err := db.QueryRowContext(ctx, query, watchListItemID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Checks if a watchlist_item_note exists
func (c *WatchlistItemNoteService) CheckIfWatchlistItemNoteExists(watchListItemID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM watchlist_item_note WHERE watchlist_item_id = $1)
		`
	var exists bool
	err := db.QueryRowContext(ctx, query, watchListItemID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Checks if a user_id is the owner of the watchlist_item_note > watchlist_item > watchlist
func (c *WatchlistItemNoteService) CheckIfUserOwnsWatchlistItem(userID uuid.UUID, watchListItemID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM watchlist_item
			JOIN watchlist ON watchlist_item.watchlist_id = watchlist.id
			WHERE watchlist_item.id = $1
			AND watchlist.users_id = $2
		)
	`
	var exists bool
	err := db.QueryRowContext(ctx, query, watchListItemID, userID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}