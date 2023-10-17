package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/go-chi/chi/v5"
)


var watchlist services.Watchlist

// GET/watchlists - this will get all watchlists for testing purposes. will make one specific to the user
func GetAllWatchlists(w http.ResponseWriter, r *http.Request) {
	all, err := watchlist.GetAllWatchlists()
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
	if watchlistErr != nil{
		helpers.MessageLogs.ErrorLog.Println(watchlistErr)
	}

	helpers.WriteJSON(w, http.StatusOK, watchlistData)

}

// POST/watchlists -- making some without user ID first -- this makes 1 watch list only TODO// make user id required
func CreateWatchlists(w http.ResponseWriter, r *http.Request){
	var watchlistData services.Watchlist
	err := json.NewDecoder(r.Body).Decode(&watchlistData)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	watchlistCreated, err := watchlist.CreateWatchlist(watchlistData)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, watchlistCreated)
}