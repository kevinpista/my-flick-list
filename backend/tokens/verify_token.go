package tokens

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
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

// Function to help CONTROLLERS verify JWT token passed by User's request headers; make use of VerifyToken call which
// simply returns a boolean
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

// Function to help CONTROLLERS verify JWT token passed by User's request headers + returns the user_id from claims
// Directly calls ParseAndValidateToken fct in order to validate and access the claims data to extract user_id
func VerifyUserJWTAndFetchUserId (r *http.Request) (uuid.UUID, error) {

	// Extract the token from the Authorization header
	tokenString := helpers.ExtractTokenFromHeader(r)
	if tokenString == "" {
		// No token provided err
		helpers.MessageLogs.ErrorLog.Println("no JWT found") // internal log
		tokenStringErr := errors.New(error_constants.UnauthorizedRequest)
		return uuid.Nil, tokenStringErr
	}
	// Calls the ParseToken to validate directly
	claims, err := ParseAndValidateToken(tokenString)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
			return uuid.Nil, tokenErr
		}

		if err == jwt.ErrTokenExpired {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.TokenExpired)
			return uuid.Nil, tokenErr
		}

		if err == jwt.ErrTokenMalformed {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
			return uuid.Nil, tokenErr
		}

		if err == jwt.ErrTokenInvalidAudience {
			helpers.MessageLogs.ErrorLog.Println(err)
			tokenErr := errors.New(error_constants.UnauthorizedRequest)
			return uuid.Nil, tokenErr
		}

		// Other error
		helpers.MessageLogs.ErrorLog.Println("other JWT or claims error:", err)
		tokenErr := errors.New(error_constants.UnauthorizedRequest)
		return uuid.Nil, tokenErr
	}

	// Valid token, extract the user_id from claims
	helpers.MessageLogs.ErrorLog.Println("TESTING EXTRACT USER_ID", claims.UserID)
	return claims.UserID, nil
}