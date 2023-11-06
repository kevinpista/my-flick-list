package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/services"
)

var user services.UserService

func isValidEmail(email string) bool {
    // Define a regular expression for a basic email format validation.
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
    return regexp.MustCompile(emailRegex).MatchString(email)
}

// POST/user-registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var userData services.UserService
    err := json.NewDecoder(r.Body).Decode(&userData.User)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) // internal log
		helpers.ErrorJSON(w, errors.New("invalid request data"), 400) // external frontend
        return
    }
    if !isValidEmail(userData.User.Email) {
		helpers.MessageLogs.ErrorLog.Println("email format invalid interal") // internal log
		helpers.ErrorJSON(w, errors.New("email format INVALID"), 406) // external
        return
    }
	// TODO AT HERE
    userCreated, err := user.RegisterUser(userData.User)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New("database insertion failed"), 406) // external
        return
    }
    // Respond with the newly created user excluding the password if successful
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


// GET/users -- testing purposes only
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	all, err := user.GetAllUsers()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"users": all})
}