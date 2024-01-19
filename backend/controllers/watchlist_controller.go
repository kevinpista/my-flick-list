package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

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

// GET/watchlists-by-user-id - user_id fetched from JWT token - returns list of watchlists.
// Used for watchlist page
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

// GET/watchlists/movie/:movieID - userID fetched from JWT token
// Fetches all watchlists belonging to user. Includes count of watchlist_items + boolean if any watchlist_items
// of a particular watchlist references the movieID that is queried
// Use for movie page and the dialog dropdown menu for user to pick which watchlist to add to
func GetWatchlistsByUserIDWithMovieIDCheck(w http.ResponseWriter, r *http.Request) {
	// Verify JWT token sent in by user and fetch their UserID
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	// Extract movie ID from URL path parameter
	movieIDStr := chi.URLParam(r, "movieID")

	if movieIDStr == "" {
		helpers.MessageLogs.ErrorLog.Println("User did not include movieID in URL")
		helpers.ErrorJSON(w, errors.New("movieID not provided"), http.StatusBadRequest)
		return
	}

	movieID, err := strconv.Atoi(movieIDStr)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println("Error converting string to number")
		helpers.ErrorJSON(w, errors.New("error converting string to number"), http.StatusBadRequest)
	}

	// Service call
	results, err := watchlist.GetWatchlistsByUserIDWithMovieIDCheck(userID, movieID)
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

// POST/watchlists - user_id fetched from JWT token - returns the new watchlist details
func CreateWatchlist(w http.ResponseWriter, r *http.Request) {
	var watchlistData services.WatchlistService

	err := json.NewDecoder(r.Body).Decode(&watchlistData.Watchlist)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest)
		return
	}

	// Checks for empty or invalid names + descriptions & trim trailing space

	watchlistData.Watchlist.Name = strings.TrimSpace(watchlistData.Watchlist.Name)
	if watchlistData.Watchlist.Name == "" {
		helpers.MessageLogs.ErrorLog.Println("empty watchlist name field")
		helpers.ErrorJSON(w, errors.New("watchlist name cannot be empty"), http.StatusBadRequest)
		return
	}

	watchlistData.Watchlist.Description = strings.TrimSpace(watchlistData.Watchlist.Description)
	if watchlistData.Watchlist.Description == "" {
		helpers.MessageLogs.ErrorLog.Println("empty watchlist description field")
		helpers.ErrorJSON(w, errors.New("watchlist description cannot be empty"), http.StatusBadRequest)
		return
	}

	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

	createdWatchlist, queryErr := watchlist.CreateWatchlist(userID, watchlistData.Watchlist)
	if queryErr != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, createdWatchlist)
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

// PATCH/watchlist-name - user_id fetched from JWT token
func UpdateWatchlistNameByID(w http.ResponseWriter, r *http.Request) {
	// New watchlist name passed through JSON body. Decode
	var watchlistData services.WatchlistService

	err := json.NewDecoder(r.Body).Decode(&watchlistData.Watchlist)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) // internal log
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest) // external frontend
        return
    }
	
	// Trim any trailing white space
	watchlistData.Watchlist.Name = strings.TrimSpace(watchlistData.Watchlist.Name)

	if watchlistData.Watchlist.Name == "" {
        helpers.MessageLogs.ErrorLog.Println("empty name field")
        helpers.ErrorJSON(w, errors.New(error_constants.InvalidName), http.StatusBadRequest)
        return
    }

	watchlistID := r.URL.Query().Get("id")
	// Check if 'id' is provided in the URL param
	if watchlistID == "" {
		helpers.MessageLogs.ErrorLog.Println("id not provided for PATCH request")
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

	// Service function call to update watchlist name
	watchlist, err := watchlist.UpdateWatchlistName(watchlistIDInt, watchlistData.Watchlist)
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"name": watchlist.Name}) // return name only
}

// PATCH/watchlist-description - user_id fetched from JWT token
func UpdateWatchlistDescriptionByID(w http.ResponseWriter, r *http.Request) {
	// New watchlist description passed through JSON body. Decode
	var watchlistData services.WatchlistService

	err := json.NewDecoder(r.Body).Decode(&watchlistData.Watchlist)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) // internal log
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest) // external frontend
        return
    }
	
	// Trim any trailing white space
	watchlistData.Watchlist.Description = strings.TrimSpace(watchlistData.Watchlist.Description)

	if watchlistData.Watchlist.Description == "" {
        helpers.MessageLogs.ErrorLog.Println("empty description field")
        helpers.ErrorJSON(w, errors.New(error_constants.InvalidName), http.StatusBadRequest)
        return
    }

	watchlistID := r.URL.Query().Get("id")
	// Check if 'id' is provided in the URL param
	if watchlistID == "" {
		helpers.MessageLogs.ErrorLog.Println("id not provided for PATCH request")
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

	// Service function call to update watchlist description
	watchlist, err := watchlist.UpdateWatchlistDescription(watchlistIDInt, watchlistData.Watchlist)
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

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"description": watchlist.Description}) // return description only
}