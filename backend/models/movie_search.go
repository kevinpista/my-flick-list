package models

// TODO finalize fields we want to fetch from search API to display on our front end as a search
// results list
type MovieSearch struct {
	Adult          bool     `json:"adult"`
	BackdropPath   string   `json:"backdrop_path"`
	GenreIds       []int    `json:"genre_ids"`
	ID             int      `json:"id"`
	// OriginalLanguage string `json:"original_language"`
	OriginalTitle  string   `json:"original_title"`
	Overview       string   `json:"overview"`
	Popularity     float64  `json:"popularity"`
	PosterPath     string   `json:"poster_path"`
	ReleaseDate    string   `json:"release_date"`
	// Title          string   `json:"title"`
	// Video          bool     `json:"video"`
	// VoteAverage    float64  `json:"vote_average"`
	// VoteCount      int      `json:"vote_count"`
}

// Test for now