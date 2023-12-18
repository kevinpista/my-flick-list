package models

type TMDBMovie struct {

	/*
    BelongsToCollection *struct {
        // Define struct for the "belongs_to_collection" field if needed
    } `json:"belongs_to_collection"`
	*/

    // Homepage            string `json:"homepage"`
    // IMDbID              string `json:"imdb_id"`
    // OriginalLanguage    string `json:"original_language"`

    // Popularity          float64 `json:"popularity"`
	/*
    ProductionCompanies []struct {
        ID            int    `json:"id"`
        LogoPath      string `json:"logo_path"`
        Name          string `json:"name"`
        OriginCountry string `json:"origin_country"`
    } `json:"production_companies"`
	*/
    
	ID                  int    `json:"id"`
	OriginalTitle       string `json:"original_title"`
	ReleaseDate string `json:"release_date"`
    Overview            string `json:"overview"`
    Revenue     int    `json:"revenue"`
	Budget              int     `json:"budget"`
    Runtime     int    `json:"runtime"`
    Status      string `json:"status"`
    Tagline     string `json:"tagline"`
    Title       string `json:"title"`
    Video       bool   `json:"video"`
    VoteAverage float64 `json:"vote_average"`
    VoteCount   int    `json:"vote_count"`
	Adult               bool    `json:"adult"`
    BackdropPath        string  `json:"backdrop_path"`
	PosterPath          string  `json:"poster_path"`
    Genres              []struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    } `json:"genres"`
}

// If TMDB returns an error status code, decode its response to this model
type TMDBError struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}