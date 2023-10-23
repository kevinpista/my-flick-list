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

// GET/watchlists/{userID}
func GetAllWatchlistsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := helpers.ConvertStringToUUID(userIDStr) // parameter will be a string. convert to int
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	all, err := watchlist.GetAllWatchlistsByUserID(userID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlists": all})
}

// GET/watchlist{id}
func GetWatchlistByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr) // parameter will be a string. convert to int
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
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

// POST/watchlists -- making some without user ID first -- this makes 1 watch list only TODO// make user id required
func CreateWatchlists(w http.ResponseWriter, r *http.Request) {
	var watchlistData services.WatchlistService

	err := json.NewDecoder(r.Body).Decode(&watchlistData.Watchlist)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	watchlistCreated, err := watchlist.CreateWatchlist(watchlistData.Watchlist)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistCreated)
}