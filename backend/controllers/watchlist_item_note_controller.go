package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"database/sql"
	"errors"


	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

var watchlistItemNote services.WatchlistItemNoteService

// GET/watchlist-item-note/{watchlistItemID}
func GetWatchlistItemNoteByWatchlistItemID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "watchlistItemID")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
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
	}

	watchlistItemNoteCreated, err := watchlistItemNoteData.CreateWatchlistItemNote(watchlistItemNoteData.WatchlistItemNote)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistItemNoteCreated)
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