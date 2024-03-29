package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"database/sql"
	"errors"


	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/tokens"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/kevinpista/my-flick-list/backend/helpers/error_constants"
)

var watchlistItemNote services.WatchlistItemNoteService

// GET/watchlist-item-note/{watchlistItemID}
func GetWatchlistItemNoteByWatchlistItemID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "watchlistItemID")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	watchlistItemNoteData, watchlistItemNoteErr := watchlistItemNote.GetWatchlistItemNoteByWatchlistItemID(id)
    if watchlistItemNoteErr != nil {
        if watchlistItemNoteErr == sql.ErrNoRows {
            helpers.ErrorJSON(w, errors.New("watchlist item note not found"), http.StatusNotFound)
        } else {
            helpers.ErrorJSON(w, watchlistItemNoteErr, http.StatusBadRequest)
        }
        return
    }
	helpers.WriteJSON(w, http.StatusOK, watchlistItemNoteData)
}

// POST/watchlist-item-note
func CreateWatchlistItemNote(w http.ResponseWriter, r *http.Request) {
	var watchlistItemNoteData services.WatchlistItemNoteService

	err := json.NewDecoder(r.Body).Decode(&watchlistItemNoteData.WatchlistItemNote)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Trim any trailing white space
	watchlistItemNoteData.WatchlistItemNote.ItemNotes = strings.TrimSpace(watchlistItemNoteData.WatchlistItemNote.ItemNotes)

	// Verify user's JWT Token
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	// Check if the particular watchlist_item even exists in the watchlist_item DB table
	watchlistItemExists, err := watchlistItemNote.CheckIfWatchlistItemExists(watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// If does not exists, no watchlist_item for the note to reference
	if !watchlistItemExists {
		helpers.MessageLogs.ErrorLog.Println("Watchlist_item does not exist. Note cannot be created")
		helpers.ErrorJSON(w, errors.New("watchlist_item not found"), http.StatusBadRequest)
		return
	}

	// Check if watchlist_item_note already exists 
	watchlistItemNoteExists, err := watchlistItemNote.CheckIfWatchlistItemNoteExists(watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// If exists, cannot create a new one. Return error.
	if watchlistItemNoteExists {
		helpers.MessageLogs.ErrorLog.Println("Watchlist_item_note already created. Must use PATCH endpoint to update")
		helpers.ErrorJSON(w, errors.New("watchlist_item_note already exists. cannot create new note object"), http.StatusBadRequest)
		return
	}

	// Check if user is the owner of the watchlist_item_note > watchlist_item > watchlist
	userIsOwner, err := watchlistItemNote.CheckIfUserOwnsWatchlistItem(userID, watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !userIsOwner {
		helpers.MessageLogs.ErrorLog.Println("User is not the owner if this watchlist_item. Unable to create note")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
		return
	}

	// At this point, user is verified, watchlist_item exists, and note does not exist.
	watchlistItemNoteCreated, err := watchlistItemNoteData.CreateWatchlistItemNote(watchlistItemNoteData.WatchlistItemNote)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistItemNoteCreated)
}

// PATCH/watchlist-item-note
func UpdateWatchlistItemNote(w http.ResponseWriter, r *http.Request) {
	var watchlistItemNoteData services.WatchlistItemNoteService

	err := json.NewDecoder(r.Body).Decode(&watchlistItemNoteData.WatchlistItemNote)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Trim any trailing white space
	watchlistItemNoteData.WatchlistItemNote.ItemNotes = strings.TrimSpace(watchlistItemNoteData.WatchlistItemNote.ItemNotes)

	// Verify user's JWT Token
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	// Check if the particular watchlist_item even exists in the watchlist_item DB table
	watchlistItemExists, err := watchlistItemNote.CheckIfWatchlistItemExists(watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// If does not exists, no watchlist_item for the note to reference
	if !watchlistItemExists {
		helpers.MessageLogs.ErrorLog.Println("Watchlist_item does not exist. Note cannot be created")
		helpers.ErrorJSON(w, errors.New("watchlist_item not found"), http.StatusBadRequest)
		return
	}

	// Check if watchlist_item_note already exists 
	watchlistItemNoteExists, err := watchlistItemNote.CheckIfWatchlistItemNoteExists(watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// If does not exist, cannot update note. Must create one first
	if !watchlistItemNoteExists {
		helpers.MessageLogs.ErrorLog.Println("Watchlist_item_note not created yet. Must use POST endpoint to create first")
		helpers.ErrorJSON(w, errors.New("watchlist_item_note does not exist - cannot update note"), http.StatusBadRequest)
		return
	}

	// Check if user is the owner of the watchlist_item_note > watchlist_item > watchlist
	userIsOwner, err := watchlistItemNote.CheckIfUserOwnsWatchlistItem(userID, watchlistItemNoteData.WatchlistItemNote.WatchlistItemID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !userIsOwner {
		helpers.MessageLogs.ErrorLog.Println("User is not the owner if this watchlist_item. Unable to update note")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
		return
	}

	// At this point, user is verified, watchlist_item exists, and note exists.
	updatedWatchlistItemNote, err := watchlistItemNoteData.UpdateWatchlistItemNote(watchlistItemNoteData.WatchlistItemNote)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, updatedWatchlistItemNote)
}

// Fetches all notes in database. Testing purposes only
// GET/watchlist-item-note
func GetNotesTest(w http.ResponseWriter, r *http.Request) {

	watchlistItemNoteData, watchlistItemNoteErr := watchlistItemNote.GetNotesTest()
    if watchlistItemNoteErr != nil {
            helpers.ErrorJSON(w, watchlistItemNoteErr, http.StatusBadRequest)
        
        return
    }
	helpers.WriteJSON(w, http.StatusOK, watchlistItemNoteData)
}