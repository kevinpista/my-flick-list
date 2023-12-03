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
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}
	/* When endpoint took in user_id via url parameter instead of from JWT Token
	userIDStr := chi.URLParam(r, "userID")
	userID, err := helpers.ConvertStringToUUID(userIDStr) // parameter will be a string. convert to int
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	*/
	all, err := watchlist.GetAllWatchlistsByUserID(userID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlists": all})
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