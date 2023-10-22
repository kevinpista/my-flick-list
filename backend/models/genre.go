package models

type Genre struct {
	MovieID int    `json:"movie_id"`
	GenreID int    `json:"genre_id"`
	Genre   string `json:"genre"`
}