package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/kevinpista/my-flick-list/backend/tokens"
	"github.com/kevinpista/my-flick-list/backend/helpers/error_constants"
)

var watchlist services.WatchlistService

// GET/watchlists - this will get all watchlists for testing purposes
func GetAllWatchlists(w http.ResponseWriter, r *http.Request) {
	all, err := watchlist.GetAllWatchlists()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlists": all})
}

// GET/watchlists-by-user-id - user_id fetched from JWT token
func GetWatchlistsByUserID(w http.ResponseWriter, r *http.Request) {
	// Verify JWT token sent in by user and fetch their UserID
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	results, err := watchlist.GetAllWatchlistsByUserID(userID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	// User may not have any watchlists yet. Services wil return "null"
	if results == nil {
		helpers.MessageLogs.ErrorLog.Println("User has no watchlists created yet")
		helpers.WriteJSON(w, http.StatusNoContent, helpers.Envelope{})
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlists": results})
}

// GET/watchlist/{id}
func GetWatchlistByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr) // parameter will be a string. convert to int
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	watchlistData, watchlistErr := watchlist.GetWatchlistByID(id)
	if watchlistErr != nil {
		if watchlistErr == sql.ErrNoRows {
			helpers.ErrorJSON(w, errors.New("watchlist not found"), http.StatusNotFound)
		} else {
			helpers.ErrorJSON(w, watchlistErr, http.StatusBadRequest)
		}
		return
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistData)
}

// POST/watchlists - user_id fetched from JWT token
func CreateWatchlist(w http.ResponseWriter, r *http.Request) {
	var watchlistData services.WatchlistService

	err := json.NewDecoder(r.Body).Decode(&watchlistData.Watchlist)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest)
		return
	}

	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	watchlistCreated, err := watchlist.CreateWatchlist(userID, watchlistData.Watchlist)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistCreated)
}

// DELETE /watchlist?id={watchlistID}
func DeleteWatchlistByID(w http.ResponseWriter, r *http.Request) {
	watchlistID := r.URL.Query().Get("id")

	// Check if 'id' is provided in the URL param
	if watchlistID == "" {
		helpers.MessageLogs.ErrorLog.Println("id not provided for DELETE request")
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	// Check if URL param is an integer
	watchlistIDInt, err := strconv.Atoi(watchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New("id parameter must be an integer"), http.StatusBadRequest)
		return
	}

	// Get userID from JWT token
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized)
		return
	}

	// Get user_id via the watchlist_id
	watchlistOwnerID, err := watchlist.GetWatchlistOwnerUserID(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Compare watchlistOwnerID vs jwt user_id
    if userID != watchlistOwnerID {
        helpers.MessageLogs.ErrorLog.Println("User is not the owner of the watchlist")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
        return
	}
	// Service function call to delete watchlist
	err = watchlist.DeleteWatchlistByID(watchlistIDInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.MessageLogs.ErrorLog.Println("Watchlist does not exist")
			helpers.ErrorJSON(w, errors.New("watchlist does not exist"), http.StatusBadRequest)
		} else {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusBadRequest)
		}
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "Watchlist deleted successfully"})
}