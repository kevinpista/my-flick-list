package services

import (
	"context"
	"time"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	User models.User
}

type RegistrationResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (us *UserService) RegisterUser(user models.User) (*RegistrationResponse, error) {
	// Hash the user's password
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
        VALUES ($1, $2, $3, $4, $5) returning id, name, created_at, updated_at
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
	).Scan(&registrationResponse.ID, &registrationResponse.Name, &registrationResponse.CreatedAt, &registrationResponse.UpdatedAt) // populates these fields; id and time() handled by DB

	if queryErr != nil {
		return nil, queryErr
	}
	return &registrationResponse, nil
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
	} 	else {
			return nil, sql.ErrNoRows // Case where query did not find any matching rows
	}
}

