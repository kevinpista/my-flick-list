package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kevinpista/my-flick-list/backend/controllers"
	"github.com/kevinpista/my-flick-list/backend/controllers/tmdb_controllers"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Movie resources from internal database
	router.Post("/api/movie", controllers.CreateMovie)
    router.Get("/api/movie/{id}", controllers.GetMovieByID)
    router.Get("/api/movies", controllers.GetAllMovies)

	// Watchlist resources
	// router.Get("/api/watchlists", controllers.GetAllWatchlists)     // GET ALL watchlists in database for testing purposes only
    router.Get("/api/watchlists-by-user-id", controllers.GetWatchlistsByUserID) // GET all watchlists belonging to specific user; user_id in JWT token
	router.Post("/api/watchlist", controllers.CreateWatchlist)    // POST a watchlist; user_id retrieved from JWT token
	router.Get("/api/watchlists/movie/{movieID}", controllers.GetWatchlistsByUserIDWithMovieIDCheck) // GET all watchlists belong to user + watchlist_item count for each + boolean if queried movieID is in the watchlist
	router.Get("/api/watchlist/{id}", controllers.GetWatchlistByID) // GET a specific watchlist by watchlist ID
	router.Delete("/api/watchlist", controllers.DeleteWatchlistByID) // DELETE watchlist via its id
	// expects "?id={watchlistID}" query param
	router.Patch("/api/watchlist-name", controllers.UpdateWatchlistNameByID) // PATCH a watchlist name
	// expects "?id={watchlistID}" query param + new name in the json body
	router.Patch("/api/watchlist-description", controllers.UpdateWatchlistDescriptionByID) // PATCH watchlist description
	// expects "?id={watchlistID}" query param + new description in the json body"

	// Watchlist-items resources
	router.Post("/api/watchlist-item", controllers.CreateWatchlistItemByWatchlistID) // POST create a watchlist item for a specific watchlist

	router.Get("/api/watchlist-items-with-movies", controllers.GetAllWatchlistItemsWithMoviesByWatchlistID) // GET fetch all watchlist items along with full movie data
	// expects "?watchlistID={watchlistID}" query param
	router.Delete("/api/watchlist-item", controllers.DeleteWatchlistItemByID) // DELETE watchlist item via its id
	// expects "?id={watchlistItemID}" query param
	router.Put("/api/watchlist-item-checkmarked", controllers.UpdateCheckmarkedBooleanByWatchlistItemByID) // PUT watchlist item checkmarked boolean update
	// expects "?id={watchlistItemID}" query parameter + 'checkmarked' field with boolean in the JSON body

	// Watchlist-item-note resources
    router.Post("/api/watchlist-item-note", controllers.CreateWatchlistItemNote) // POST create watchlist item note for a specific watchlist item
    router.Get("/api/watchlist-item-note/{watchlistItemID}", controllers.GetWatchlistItemNoteByWatchlistItemID) // GET fetch the note for a specific watchlist item
    router.Get("/api/watchlist-item-note", controllers.GetNotesTest) // GET fetch all notes in database. Testing purposes only
	router.Patch("/api/watchlist-item-note", controllers.UpdateWatchlistItemNote) // PATCH watchlist_item_note 'item notes'. Passed through JSON body

	// Genre resources
    router.Post("/api/genre", controllers.CreateGenreDataByMovieID) // POST add genre data for a movie
    router.Get("/api/genre/{movieID}", controllers.GetGenreByMovieID) // GET the genre data for a movie
	router.Get("/api/genre", controllers.GetAllGenres) // GET all genre entries TESTING PURPOSES

	// User resources
    router.Post("/api/user-registration", controllers.RegisterUser) // POST register a user
	router.Post("/api/user-login", controllers.HandleLogin) // POST user login
    router.Get("/api/user/{userID}", controllers.GetUserByID) // GET user by their user id


    // TMDB movie search resources
    router.Get("/api/tmdb-search", tmdb_controllers.TMDBSearchMovieByKeyWords)
    // expects "?query={keywords+keywords..}" query paramter

    // TMDB GET movie resources
    router.Get("/api/tmdb-movie", tmdb_controllers.TMDBGetMovieByID)
    // expects "?query={movie_id}" query paramter

	// TMDB GET movie trailer resources
	router.Get("/api/tmdb-trailer", tmdb_controllers.TMDBGetMovieTrailerByID)
	// expects "?query={movie_id}" query paramter

	return router
}
