package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	// "github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

var watchlistItem services.WatchlistItem

// GET/watchlist-items?watchlistID={watchlistID}
func GetAllWatchlistItemsByWatchListID(w http.ResponseWriter, r *http.Request) {
    watchlistID := r.URL.Query().Get("watchlistID")

	// Check if watchlistID is empty or not provided in the URL
	if watchlistID == "" {
		http.Error(w, "watchlistID parameter is missing", http.StatusBadRequest)
		return
	}

	watchlistIDInt, err := strconv.Atoi(watchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	all, err := watchlistItem.GetAllWatchlistItemsByWatchlistID(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlist-items": all})
}

// POST/watchlist-item
func CreateWatchlistItemByWatchlistID(w http.ResponseWriter, r *http.Request) {
	var watchlistItemData services.WatchlistItem

	err := json.NewDecoder(r.Body).Decode(&watchlistItemData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	watchlistItemCreated, err := watchlistItemData.CreateWatchlistItemByWatchlistID(watchlistItemData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	helpers.WriteJSON(w, http.StatusOK, watchlistItemCreated)
}