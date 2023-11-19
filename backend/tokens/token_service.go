package tokens

import (
    "time"

    "github.com/golang-jwt/jwt/v4"
    "github.com/google/uuid"

)

// Custom claims structure
type Claims struct {
    UserID uuid.UUID `json:"user_id"`
    jwt.RegisteredClaims
}

// TODO - Secret key for signing tokens -- will load from environment variable
var jwtSecret = []byte("custom_secret_key_change_later")

// Generate a new JWT token
func GenerateToken(userID uuid.UUID) (string, error) {
    // Token expires in 24 hours of the jwt.NumericDate type
    expTime := &jwt.NumericDate{
       Time: time.Now().Add(time.Hour * 24),
    }
    
    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: expTime, 
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Parses JWT token string into claims
func ParseAndValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil // Call back function provided to verify signature of the token against my jwtSecret key
    })

    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok { // Attempts to convert the value of err to a type jwt.ValidationError
            // If conversion successful, it will store the value in ve. Set ok variable to true if conversion was successful
            switch ve.Errors {
            case jwt.ValidationErrorSignatureInvalid:
                return nil, jwt.ErrSignatureInvalid
            case jwt.ValidationErrorExpired:
                return nil, jwt.ErrTokenExpired
            case jwt.ValidationErrorMalformed:
                return nil, jwt.ErrTokenMalformed
            case jwt.ValidationErrorAudience:
                return nil, jwt.ErrTokenInvalidAudience
            default:
                // Other JWT validation error
                return nil, err
            }
        } else {
            // Handle any errors not related to validation error
            return nil, err
        }
    }

    claims, ok := token.Claims.(*Claims) // Extract claims from the parsed *jwt.Token object and store into my Claims struct
    if !ok || !token.Valid { // Invalid if either are True
        return nil, jwt.ErrTokenInvalidClaims
    }

    return claims, nil // Success
}

// Verify JWT token
func VerifyToken(tokenString string) (bool, error) {
    claims, err := ParseAndValidateToken(tokenString)
    if err != nil {
        return false, err
    }
    
    now := time.Now() // Checks if expired
    if now.After(claims.ExpiresAt.Time) {
        return false, jwt.ErrTokenExpired
    }

    return true, nil
}