package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"


	"github.com/go-chi/chi/v5"
	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/helpers/error_constants"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/kevinpista/my-flick-list/backend/tokens"
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
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest) // external frontend
        return
    }
	// Trim leading and trailing white spaces from the email address if any
	userData.User.Email = strings.TrimSpace(userData.User.Email)

    if !isValidEmail(userData.User.Email) {
		helpers.MessageLogs.ErrorLog.Println("user entered an invalid email format")
		helpers.ErrorJSON(w, errors.New(error_constants.InvalidEmail), http.StatusBadRequest)
        return
    }
    // Check for empty name field or whitespace-only name
    if len(strings.TrimSpace(userData.User.Name)) == 0 {
        helpers.MessageLogs.ErrorLog.Println("empty or whitespace-only name")
        helpers.ErrorJSON(w, errors.New(error_constants.InvalidName), http.StatusBadRequest)
        return
    }
    if userData.User.Password == "" {
        helpers.MessageLogs.ErrorLog.Println("empty password field")
        helpers.ErrorJSON(w, errors.New(error_constants.PasswordEmpty), http.StatusBadRequest)
        return
    }
    // Check for whitespace in the password
    if strings.Contains(userData.User.Password, " ") {
        helpers.MessageLogs.ErrorLog.Println("password contains whitespace")
        helpers.ErrorJSON(w, errors.New(error_constants.PasswordWhitespace), http.StatusBadRequest)
        return
    }
    userCreated, err := user.RegisterUser(userData.User)
	if err != nil {
        // Check the error message for "duplicate key value violates unique constraint" and SQLSTATE 23505
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") &&
            strings.Contains(err.Error(), "SQLSTATE 23505") {
            
			helpers.MessageLogs.ErrorLog.Println(err)
			helpers.ErrorJSON(w, errors.New(error_constants.EmailExists), http.StatusConflict)
			return
        }
		// Other issues related to services/database during account creation
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest)
    }
	// If userCreated successful, Generate JWT Token
	jwtToken, err := tokens.GenerateToken(userCreated.ID) // Pass in UUID generated by Postgresql
	if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusInternalServerError) // Error generating JWT token
		return
	}
    // Respond with the newly created user with new info including ID,  but excluding the password
	helpers.WriteJSONWithToken(w, http.StatusCreated, userCreated, jwtToken)
}

// POST/user-login
func HandleLogin(w http.ResponseWriter, r *http.Request) {
    var receivedUserData services.UserService
    err := json.NewDecoder(r.Body).Decode(&receivedUserData.User)
    if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) // internal log
		helpers.ErrorJSON(w, errors.New(error_constants.BadRequest), http.StatusBadRequest) // external frontend
        return
    }
	// Trim leading and trailing white spaces from the email address + passsword if any
	// Note password field during registration does not allow whitespaces
	receivedUserData.User.Email = strings.TrimSpace(receivedUserData.User.Email)
	receivedUserData.User.Password = strings.TrimSpace(receivedUserData.User.Password)

    if !isValidEmail(receivedUserData.User.Email) {
		helpers.MessageLogs.ErrorLog.Println("user entered an invalid email format")
		helpers.ErrorJSON(w, errors.New(error_constants.InvalidEmail), http.StatusBadRequest)
        return
    }

	userLogin, err := user.HandleLogin(receivedUserData.User)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err) // if user not found, exact error string from postgresql
		helpers.ErrorJSON(w, errors.New(error_constants.InvalidLogin), http.StatusUnauthorized)
		return
	}

	// If userLogin successful, Generate JWT Token
	jwtToken, err := tokens.GenerateToken(userLogin.ID) // Pass in UUID generated by Postgresql
	if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
		helpers.ErrorJSON(w, err, http.StatusInternalServerError) // Error generating JWT token
		return
	}
    // Respond with the newly created user with new info including ID,  but excluding the password
	helpers.WriteJSONWithToken(w, http.StatusCreated, userLogin, jwtToken)
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

/*
// GET/users -- testing purposes only -- with JWT header check
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	tokenErr := tokens.VerifyUserJWT(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}
	// User's JWT token is valid
	all, err := user.GetAllUsers()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"users": all})
}
*/