package tokens

import (
	"errors"
	"net/http"


	"github.com/kevinpista/my-flick-list/backend/helpers"
	"github.com/kevinpista/my-flick-list/backend/helpers/error_constants"
	"github.com/golang-jwt/jwt/v4"

)

/*
// Add this code below to all controllers that need to check for a user JWT in their request headers

	tokenErr := tokens.VerifyUserJWT(r)
	if tokenErr != nil {
		helpers.ErrorJSON(w, tokenErr, http.StatusUnauthorized) // tokenErr will be a errors.New(error_constants) object
		return
	}

*/

// Function to help controllers verify JWT token passed by User's request headers
func VerifyUserJWT (r *http.Request) error {
	// Extract the token from the Authorization header
	tokenString := helpers.ExtractTokenFromHeader(r) // Send request header in 
	if tokenString == "" {
		// No token provided, handle accordingly (e.g., unauthorized access)
		tokenErr := errors.New(error_constants.UnauthorizedRequest)
		helpers.MessageLogs.ErrorLog.Println("no JWT found") // internal log
		return tokenErr
	}

	verifyToken, err := VerifyToken(tokenString)
	if err != nil && !verifyToken {

		if err == jwt.ErrSignatureInvalid {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
		return tokenErr
		}

		if err == jwt.ErrTokenExpired {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.TokenExpired)
			return tokenErr
		}

		if err == jwt.ErrTokenMalformed {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
			return tokenErr
		}

		if err == jwt.ErrTokenInvalidAudience {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
			return tokenErr
		}

		// Other error
		helpers.MessageLogs.ErrorLog.Println("other JWT error:", err)
		tokenErr := errors.New(error_constants.UnauthorizedRequest)
		return tokenErr
	}

	// Else token is valid
	return nil

}