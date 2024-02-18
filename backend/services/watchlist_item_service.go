package services

import (
	"context"
	// "errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
)

type WatchlistItemService struct {
	WatchlistItem          models.WatchlistItem
	WatchlistItemWithMovie models.WatchlistItemWithMovie
}

var movieQuery TMDBMovieService
var watchlistItemHelper WatchlistItemService

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

// Fetches watchlist name & description + all watchlist items and its movie data + watchlist_item_note if exists
func (c *WatchlistItemService) GetWatchlistWithWatchlistItemsByWatchlistID(watchlistID int) ([]*models.WatchlistItemWithMovie, string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Fetch watchlist name and description
	var watchlistName, watchlistDescription string
	err := db.QueryRowContext(ctx, "SELECT name, description FROM watchlist WHERE id = $1", watchlistID).Scan(&watchlistName, &watchlistDescription)
	if err != nil {
		return nil, "", "", err
	}

	// Fetch watchlist items belonging to the watchlist
	// If a watchlist_item does not have a watchlist_note belonging to it, the COALESCE function will return
	// the values of NULL to the item_notes, notes_created_at field etc.  
	query := `
		SELECT 
			wi.id, wi.watchlist_id, wi.movie_id, wi.checkmarked, wi.created_at, wi.updated_at,
			m.original_title, m.overview, m.tagline, m.release_date, m.poster_path, m.backdrop_path, m.runtime, m.adult,
			m.budget, m.revenue, m.rating, m.votes, m.created_at AS movie_created_at, m.updated_at AS movie_updated_at,
			COALESCE(win.item_notes, NULL) AS item_notes,
			COALESCE(win.created_at, NULL) AS note_created_at,
			COALESCE(win.updated_at, NULL) AS note_updated_at
		FROM watchlist_item wi
		JOIN movie m ON wi.movie_id = m.id
		JOIN watchlist w ON wi.watchlist_id = w.id
		LEFT JOIN watchlist_item_note win ON wi.id = win.watchlist_item_id
		WHERE wi.watchlist_id = $1
	`

	rows, err := db.QueryContext(ctx, query, watchlistID)
	if err != nil {
		return nil, "", "", err
	}
	defer rows.Close()

	var watchlistItemsWithMovies []*models.WatchlistItemWithMovie
	for rows.Next() {
		var watchlistItemWithMovie models.WatchlistItemWithMovie
		err := rows.Scan(
			&watchlistItemWithMovie.ID,
			&watchlistItemWithMovie.WatchlistID,
			&watchlistItemWithMovie.MovieID,
			&watchlistItemWithMovie.Checkmarked,
			&watchlistItemWithMovie.CreatedAt,
			&watchlistItemWithMovie.UpdatedAt,
			&watchlistItemWithMovie.OriginalTitle,
			&watchlistItemWithMovie.Overview,
			&watchlistItemWithMovie.Tagline,
			&watchlistItemWithMovie.ReleaseDate,
			&watchlistItemWithMovie.PosterPath,
			&watchlistItemWithMovie.BackdropPath,
			&watchlistItemWithMovie.Runtime,
			&watchlistItemWithMovie.Adult,
			&watchlistItemWithMovie.Budget,
			&watchlistItemWithMovie.Revenue,
			&watchlistItemWithMovie.Rating,
			&watchlistItemWithMovie.Votes,
			&watchlistItemWithMovie.MovieCreatedAt,
			&watchlistItemWithMovie.MovieUpdatedAt,
			&watchlistItemWithMovie.ItemNotes,        // WatchlistItemNote
			&watchlistItemWithMovie.NoteCreatedAt,    // WatchlistItemNote
			&watchlistItemWithMovie.NoteUpdatedAt,    // WatchlistItemNote
		)

		if err != nil {
			return nil, "", "", err
		}
		watchlistItemsWithMovies = append(watchlistItemsWithMovies, &watchlistItemWithMovie)
	}
	return watchlistItemsWithMovies, watchlistName, watchlistDescription, nil
}


// Important TWO-PART service function
// Creates a watchlist_item with a movie_id with the watchlist ID it belongs to.
// If the movie data is not in the local database yet, query TMDB API to fetch movie data, add to database, then connect
// watchlist_item to the newly added movie via its movie_id
func (c *WatchlistItemService) CreateWatchlistItemByWatchlistID(watchlistItem models.WatchlistItem) (*models.WatchlistItem, error) {
	// Initial attempt to connect watchlist_item to movie_id. Success if movie_id is already in database.
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Use 'tx' to ensure atomicity of operations since the Watchlist updated_at field needs to be updated
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Check if movie with movie_id exists in the database
	movieExists, err := watchlistItemHelper.CheckIfMovieExists(watchlistItem.MovieID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Query TMDB API to add to database if movie does not exist
	if !movieExists {
		// Convert movieID integer to a string to correctly send query to TMDB API service function
		movieIDString := strconv.Itoa(watchlistItem.MovieID)

		// Service function call to query TMDB API and add to local database
		addToDatabaseErr := movieQuery.TMDBGetMovieByIDAddToLocalDatabase(movieIDString)
		if addToDatabaseErr != nil {
			// Error related to TMDBGetMovieByIDAddToLocalDatabase
			tx.Rollback()
			return nil, addToDatabaseErr
		}
	}

	// At this point, the movie should exists in database, create watchlist_item in database
	query := `
		INSERT INTO watchlist_item (watchlist_id, movie_id, checkmarked, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) returning *
		`
	err = tx.QueryRowContext(
		ctx,
		query,
		watchlistItem.WatchlistID,
		watchlistItem.MovieID,
		watchlistItem.Checkmarked,
		time.Now(), // watchlistItem.CreatedAt
		time.Now(), // watchlistItem.UpdatedAt
	).Scan(
		&watchlistItem.ID,
		&watchlistItem.WatchlistID,
		&watchlistItem.MovieID,
		&watchlistItem.Checkmarked,
		&watchlistItem.CreatedAt,
		&watchlistItem.UpdatedAt,
	)

	if err != nil {
		// Check if err is related to the movie not existing
		movieNotInDataBase := strings.Contains(strings.ToLower(err.Error()), "watchlist_item_movie_id_fkey") // Returns true if error message contains exact string error from DB
		if movieNotInDataBase {
			tx.Rollback()
			return nil, err // Movie not in DB
		} else {
			// Error unrelated to 'movie not in database'
			tx.Rollback()
			return nil, err
		} 
	}

	// Update Watchlist's updated_at time field
	updateWatchlistQuery := `
		UPDATE watchlist
		SET updated_at = $1
		WHERE id = $2
	`

	_, err = tx.ExecContext(
		ctx,
		updateWatchlistQuery,
		time.Now(),
		watchlistItem.WatchlistID,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Send in transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &watchlistItem, nil
}

// Deletes a watchlist_item with its id and updates the updated_at time of the watchlist it belongs in
func (c *WatchlistItemService) DeleteWatchlistItemByID(watchlistItemID int, watchlistID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

    // Delete watchlist_item_note (if any) that references the watchlist_item
    deleteWatchlistItemNoteQuery := `
        DELETE FROM watchlist_item_note WHERE watchlist_item_id = $1
    `
    _, err = tx.ExecContext(ctx, deleteWatchlistItemNoteQuery, watchlistItemID)
    if err != nil {
        tx.Rollback()
        return err
    }

	deleteWatchlistItemQuery := `
		DELETE FROM watchlist_item WHERE id = $1
	`
	_, err = tx.ExecContext(ctx, deleteWatchlistItemQuery, watchlistItemID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update Watchlist's updated_at time field
	updateWatchlistQuery := `
		UPDATE watchlist
		SET updated_at = $1
		WHERE id = $2
	`
	_, err = tx.ExecContext(
		ctx,
		updateWatchlistQuery,
		time.Now(),
		watchlistID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Send in transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Updates the "checkmarked" boolean of the particular watchlist_item. Updates the updated_at field of watchlist item belongs to
func (c *WatchlistItemService) UpdateCheckmarkedBooleanByWatchlistItemByID(watchlistItemID int, watchlistItem models.WatchlistItem, watchlistID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	booleanUpdateQuery := `
		UPDATE watchlist_item
		SET checkmarked = $1
		WHERE id = $2
		`
	_, err = tx.ExecContext(
		ctx,
		booleanUpdateQuery,
		watchlistItem.Checkmarked,
		watchlistItemID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	// Update Watchlist's updated_at time field
	updateWatchlistQuery := `
		UPDATE watchlist
		SET updated_at = $1
		WHERE id = $2
	`
	_, err = tx.ExecContext(
		ctx,
		updateWatchlistQuery,
		time.Now(),
		watchlistID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Send in transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}


/*
HELPER FUNCTIONS
*/

// Check if exisiting watchlist_items within a watchlist contain a movie_id as the one being passed by user
func (c *WatchlistItemService) CheckIfMovieInWatchlistExists(watchlistID, movieID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM watchlist_item WHERE watchlist_id = $1 AND movie_id = $2)
    `
	var exists bool
	err := db.QueryRowContext(ctx, query, watchlistID, movieID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Checks if a watchlist with the given ID exists
func (c *WatchlistItemService) CheckIfWatchlistExists(watchlistID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM watchlist WHERE id = $1)
		`
	var exists bool
	err := db.QueryRowContext(ctx, query, watchlistID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Get's the user_id owner of the watchlist
func (c *WatchlistItemService) GetWatchlistOwnerUserID(watchlistID int) (uuid.UUID, error) {
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

// Get's the watchlist_id of the watchlist_item
func (c *WatchlistItemService) GetWatchlistItemWatchlistId(watchlistItemID int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT watchlist_id FROM watchlist_item WHERE id = $1"
	var watchlist_id int
	err := db.QueryRowContext(ctx, query, watchlistItemID).Scan(&watchlist_id)
	if err != nil {
		return 0, err
	}

	return watchlist_id, nil
}

// Checks if a movie with the movieID exists in the local database yet
func (c *WatchlistItemService) CheckIfMovieExists(movieID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT EXISTS (SELECT 1 FROM movie WHERE id = $1)
		`
	var exists bool
	err := db.QueryRowContext(ctx, query, movieID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}