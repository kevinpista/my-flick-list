package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	User models.User
}

// Add JWT to response
type RegistrationResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Add JWT to response
type LoginResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *UserService) RegisterUser(user models.User) (*RegistrationResponse, error) {
	// Hash the user's password. Database will store hashed version
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set new user with the hashed password
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        INSERT INTO users (name, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5) returning id, name, email, created_at, updated_at
    `
	var registrationResponse RegistrationResponse
	queryErr := db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
		time.Now(),
	).Scan(&registrationResponse.ID, &registrationResponse.Name, &registrationResponse.Email, &registrationResponse.CreatedAt, &registrationResponse.UpdatedAt) // populates these fields; id and time() handled by DB

	if queryErr != nil {
		return nil, queryErr
	}
	return &registrationResponse, nil
}

func (c *UserService) HandleLogin(receivedUserData models.User) (*LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, name, email, password, created_at, updated_at FROM users
		WHERE email = $1
    `
	var fetchedUserData models.User
	queryErr := db.QueryRowContext(
		ctx,
		query,
		receivedUserData.Email,
	).Scan(
		&fetchedUserData.ID,
		&fetchedUserData.Name,
		&fetchedUserData.Email,
		&fetchedUserData.Password,
		&fetchedUserData.CreatedAt,
		&fetchedUserData.UpdatedAt,
	)
	if queryErr == sql.ErrNoRows {
		// User not found
		return nil, queryErr
	} else if queryErr != nil {
		// Other database related error
		return nil, queryErr
	}

	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(fetchedUserData.Password), []byte(receivedUserData.Password))
	if err != nil {
		// Incorrect password
		return nil, err
	}

	// Password matches, return login response
	loginResponse := LoginResponse {
		ID:        fetchedUserData.ID,
		Name:      fetchedUserData.Name,
		Email:     fetchedUserData.Email,
		CreatedAt: fetchedUserData.CreatedAt,
		UpdatedAt: fetchedUserData.UpdatedAt,
	}
	return &loginResponse, nil
}

func (c *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		SELECT id, name, email, created_at, updated_at FROM users
		WHERE id = $1
	`

	row, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	if row.Next() {
		var user models.User
		err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, sql.ErrNoRows // Case where query did not find any matching rows
	}
}
/*
// Get all Users --- testing purposes only
func (c *UserService) GetAllUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT id, name, email, password, created_at, updated_at FROM users
		`
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
*/