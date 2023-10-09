package api

import (
	"time"
	"github.com/google/uuid"
)

type Movie struct{
	ID uuid.UUID `json:"id"`
	OriginalTitle string `json:"original_title"`
	Overview string `json:"overview"`
	Tagline string `json:"tagline"`
	ReleaseDate time.Time `json:"release_date"`
	PosterPath string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
	Runtime uint16 `json:"r untime"`
	Adult bool `json:"adult"`
	Budget uint32 `json:"budget"`
	Revenue uint64 `json:"revenue"`
	Rating float32 `json:"rating"`
	Votes uint32 `json:"votes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}