package models

import(
	"github.com/google/uuid"
	"time"
)

type Watchlist struct {
	ID          int       `json:"id"`
	UserID      uuid.UUID `json:"users_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
