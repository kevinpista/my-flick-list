package controllers

import (
    "encoding/json"
    "net/http"
	"database/sql"
	"errors"

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

// GET/user{id} - id passed through body
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var userRequest services.UserService // contains the uuid only
	err := json.NewDecoder(r.Body).Decode(&userRequest.User)
	if err != nil{
		helpers.MessageLogs.ErrorLog.Println(err)
	}

	userData, err := user.GetUserByID(userRequest.User.ID) // pass in uuid only
	if err != nil {
		if err == sql.ErrNoRows {
			helpers.ErrorJSON(w, errors.New("user not found"), http.StatusNotFound)
		} else {
			helpers.ErrorJSON(w, err, http.StatusBadRequest)
		}
		return
	}
	helpers.WriteJSON(w, http.StatusOK, userData)
}