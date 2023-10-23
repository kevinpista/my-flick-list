package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

var user services.UserService

// POST/user-registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var userData services.UserService
    err := json.NewDecoder(r.Body).Decode(&userData.User)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
        return
    }

    userCreated, err := user.RegisterUser(userData.User)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    // Respond with the newly created user excluding the password
    helpers.WriteJSON(w, http.StatusCreated, userCreated)
}

// GET/user/{userID}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")

	userID, err := helpers.ConvertStringToUUID(userIDStr)
	if err !=nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	userData, userErr := user.GetUserByID(userID)
	if userErr != nil {
		if userErr == sql.ErrNoRows {
			helpers.ErrorJSON(w, errors.New("user not found"), http.StatusNotFound)
		} else {
			helpers.ErrorJSON(w, userErr, http.StatusBadRequest)
		}
		return
	}
	helpers.WriteJSON(w, http.StatusOK, userData)
}