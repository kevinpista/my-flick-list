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
func ParseToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil // Call back function provided to verify signature of the token against my jwtSecret key
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims) // Extract claims from the parsed *jwt.Token object and store into my Claims struct
    if !ok || !token.Valid { // Invalid if either are True
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil // Success
}

// Verify JWT token
func VerifyToken(tokenString string) (bool, error) {
    claims, err := ParseToken(tokenString)
    if err != nil {
        return false, err
    }

    now := time.Now() // Checks if expired
    if now.After(claims.ExpiresAt.Time) {
        return false, jwt.ErrTokenExpired
    }

    return true, nil
}