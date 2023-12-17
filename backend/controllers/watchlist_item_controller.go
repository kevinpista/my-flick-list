package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"errors"
	"database/sql"

	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/kevinpista/my-flick-list/backend/tokens"
	"github.com/kevinpista/my-flick-list/backend/helpers/error_constants"
)

var watchlistItem services.WatchlistItemService

// GET/watchlist-items?watchlistID={watchlistID} -- only returns the movie_id within each watchlist-item -- testing purposes only
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
		helpers.ErrorJSON(w, errors.New("watchlistID parameter must be an integer"), http.StatusBadRequest)
		return
	}
	// Check if the particular watchlist even exists in the watchlist DB table
	exists, err := watchlistItem.CheckIfWatchlistExists(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !exists {
		helpers.MessageLogs.ErrorLog.Println("Queried watchlist not found in DB")
		helpers.ErrorJSON(w, errors.New("watchlist not found"), http.StatusBadRequest)
		return
	}

	// Does watchlist does exist, but possible it can have no watchlist items in it so services will return " 'watchlist-items': null "
	all, err := watchlistItem.GetAllWatchlistItemsByWatchlistID(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlist-items": all})
}

// GET/watchlist-items-with-movies?watchlistID={watchlistID} -- returns full movie_data for each watchlist-item
func GetAllWatchlistItemsWithMoviesByWatchListID(w http.ResponseWriter, r *http.Request) {
    watchlistID := r.URL.Query().Get("watchlistID")

	// Check if watchlistID is empty or not provided in the URL
	if watchlistID == "" {
		helpers.MessageLogs.ErrorLog.Println("watchlistItemID not provided for GET request")
		http.Error(w, "watchlistID parameter is missing", http.StatusBadRequest)
		return
	}
	watchlistIDInt, err := strconv.Atoi(watchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New("watchlistID parameter must be an integer"), http.StatusBadRequest)
		return
	}
	// Check if the particular watchlist even exists in the watchlist DB table
	exists, err := watchlistItem.CheckIfWatchlistExists(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !exists {
		helpers.MessageLogs.ErrorLog.Println("Queried watchlist not found in DB")
		helpers.ErrorJSON(w, errors.New("watchlist not found"), http.StatusBadRequest)
		return
	}

	// Check if the user is the correct owner of the watchlist being queried based on JWT token
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

    // Query watchlist DB table to fetch owner user_id to check if it matches JWT token user_id
    watchlistOwnerID, err := watchlistItem.GetWatchlistOwnerUserID(watchlistIDInt)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        helpers.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }

    if userID != watchlistOwnerID {
        helpers.MessageLogs.ErrorLog.Println("User is not the owner of the queried watchlist")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
        return
    }

	// Watchlist exists, but possible it can have no watchlist items in it so services will return " 'watchlist-items': null "
	all, err := watchlistItem.GetAllWatchlistItemsWithMoviesByWatchListID(watchlistIDInt)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"watchlist-items": all})
}

// POST/watchlist-item
func CreateWatchlistItemByWatchlistID(w http.ResponseWriter, r *http.Request) {
	var watchlistItemData services.WatchlistItemService

	err := json.NewDecoder(r.Body).Decode(&watchlistItemData.WatchlistItem)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Check if "watchlist_id" was passed in by user. It would unmarshalled as 0 if it did not exist
	if watchlistItemData.WatchlistItem.WatchlistID == 0 {
		helpers.MessageLogs.ErrorLog.Println("User did not pass in a watchlist_id")
		helpers.ErrorJSON(w, errors.New("missing watchlist_id in request"), http.StatusBadRequest)
		return
	}

	// Check if the particular watchlist even exists in the watchlist DB table
	watchlistExists, err := watchlistItem.CheckIfWatchlistExists(watchlistItemData.WatchlistItem.WatchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if !watchlistExists {
		helpers.MessageLogs.ErrorLog.Println("Queried watchlist not found in DB")
		helpers.ErrorJSON(w, errors.New("watchlist not found"), http.StatusBadRequest)
		return
	}

	// Check if the user is the correct owner of the watchlist being queried based on JWT token
	userID, tokenErr := tokens.VerifyUserJWTAndFetchUserId(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

    // Query watchlist DB table to fetch owner user_id to check if it matches JWT token user_id
    watchlistOwnerID, err := watchlistItem.GetWatchlistOwnerUserID(watchlistItemData.WatchlistItem.WatchlistID)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        helpers.ErrorJSON(w, err, http.StatusBadRequest)
        return
    }

    if userID != watchlistOwnerID {
        helpers.MessageLogs.ErrorLog.Println("User is not the owner of the queried watchlist")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
        return
	}

	// Check if "movie_id" was passed in by user. It would unmarshalled as 0 if it did not exist
	if watchlistItemData.WatchlistItem.MovieID == 0 {
		helpers.MessageLogs.ErrorLog.Println("User did not pass in a movie_id")
		helpers.ErrorJSON(w, errors.New("missing movie_id in request"), http.StatusBadRequest)
		return
	}
	
	// Check if an existing watchlist_item has a movie_id equal to the movie_id user wants to create
	movieInWatchlistExists, err := watchlistItem.CheckIfMovieInWatchlistExists(watchlistItemData.WatchlistItem.WatchlistID, watchlistItemData.WatchlistItem.MovieID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
		}

	if movieInWatchlistExists {
		helpers.MessageLogs.ErrorLog.Println("User is attempting to add a movie they already added to their watchlist")
		helpers.ErrorJSON(w, errors.New("movie is already in watchlist"), http.StatusBadRequest)
		return
	}
	
	watchlistItemCreated, err := watchlistItemData.CreateWatchlistItemByWatchlistID(watchlistItemData.WatchlistItem)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		movieNotInDataBase := strings.Contains(strings.ToLower(err.Error()), "watchlist_item_movie_id_fkey")
		if movieNotInDataBase {
			helpers.MessageLogs.ErrorLog.Println("Movie data with this ID not yet added to DB")
			helpers.ErrorJSON(w, errors.New("movie not yet added to database "), http.StatusBadRequest)
			return
		} 
	// TO-DO will consider if want to take care of adding the missing movie data here in order to have it done in 1 go
	// Versus returning an error
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return

	}
	helpers.WriteJSON(w, http.StatusOK, watchlistItemCreated)
}

// DELETE /watchlist-items?={id}
func DeleteWatchlistItemByID(w http.ResponseWriter, r *http.Request) {
	watchlistItemID := r.URL.Query().Get("id")

	// Check if 'id' is provided in the URL param
	if watchlistItemID == "" {
		helpers.MessageLogs.ErrorLog.Println("id not provided for DELETE request")
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}
	// Check if URL param is an integer
	watchlistItemIDInt, err := strconv.Atoi(watchlistItemID)
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
	// Get watchlist_id via watchlist_item_id
	watchlistID, err := watchlistItem.GetWatchlistItemWatchlistId(watchlistItemIDInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.MessageLogs.ErrorLog.Println("Watchlist item does not exist")
			helpers.ErrorJSON(w, errors.New("watchlist item does not exist"), http.StatusBadRequest)
		} else {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusBadRequest)
		}
		return
	}
	// Get user_id via the watchlist_id
	watchlistOwnerID, err := watchlistItem.GetWatchlistOwnerUserID(watchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// Compare watchlistOwnerID vs jwt user_id
    if userID != watchlistOwnerID {
        helpers.MessageLogs.ErrorLog.Println("User is not the owner of the watchlist item")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
        return
	}
	// Service function call to delete watchlist_item
	err = watchlistItem.DeleteWatchlistItemByID(watchlistItemIDInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.MessageLogs.ErrorLog.Println("Watchlist item does not exist")
			helpers.ErrorJSON(w, errors.New("watchlist item does not exist"), http.StatusBadRequest)
		} else {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusBadRequest)
		}
		return
	}

	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "Watchlist item deleted successfully"})
}


// PUT /watchlist-item-checkmarked?id={watchlistItemID} -- expects the 'checkmarked' field to be passed through the body with boolean value
func UpdateCheckmarkedBooleanByWatchlistItemByID(w http.ResponseWriter, r *http.Request) {
	watchlistItemID := r.URL.Query().Get("id")

	// Check if 'id' is provided in the URL param
	if watchlistItemID == "" {
		helpers.MessageLogs.ErrorLog.Println("id not provided for PUT request")
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}
	// Check if URL param is an integer
	watchlistItemIDInt, err := strconv.Atoi(watchlistItemID)
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
	// Get watchlist_id via watchlist_item_id
	watchlistID, err := watchlistItem.GetWatchlistItemWatchlistId(watchlistItemIDInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.MessageLogs.ErrorLog.Println("Watchlist item does not exist")
			helpers.ErrorJSON(w, errors.New("watchlist item does not exist"), http.StatusBadRequest)
		} else {
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, err, http.StatusBadRequest)
		}
		return
	}

	// Get user_id via the watchlist_id
	watchlistOwnerID, err := watchlistItem.GetWatchlistOwnerUserID(watchlistID)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Compare watchlistOwnerID vs jwt user_id
    if userID != watchlistOwnerID {
        helpers.MessageLogs.ErrorLog.Println("User is not the owner of the watchlist item")
        helpers.ErrorJSON(w, errors.New(error_constants.UnauthorizedRequest), http.StatusUnauthorized)
        return
	}

	// Unmarshall JSON body which holds only "checkmarked" field
	var watchlistItemData services.WatchlistItemService

	errDecode := json.NewDecoder(r.Body).Decode(&watchlistItemData.WatchlistItem)
	if errDecode != nil {
		helpers.MessageLogs.ErrorLog.Println(errDecode)
		helpers.ErrorJSON(w, errDecode, http.StatusBadRequest)
		return
	}

	err = watchlistItem.UpdateCheckmarkedBooleanByWatchlistItemByID(watchlistItemIDInt, watchlistItemData.WatchlistItem)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "Checkmarked boolean updated successfully", "new_boolean": watchlistItemData.WatchlistItem.Checkmarked})
}
