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
	router.Post("/api/watchlists", controllers.CreateWatchlist)    // POST a watchlist; user_id retrieved from JWT token
	router.Get("/api/watchlist/{id}", controllers.GetWatchlistByID) // GET a specific watchlist by watchlist ID
	router.Delete("/api/watchlist", controllers.DeleteWatchlistByID) // DELETE watchlist via its id
	// expects "?id={watchlistID}" query param

	// Watchlist-items resources
	router.Post("/api/watchlist-item", controllers.CreateWatchlistItemByWatchlistID) // POST create a watchlist item for a specific watchlist
	router.Get("/api/watchlist-items", controllers.GetAllWatchlistItemsByWatchListID) // GET fetch all watchlist items from a specific watchlist.
	// expects "?watchlistID={watchlistID}" query param
	router.Get("/api/watchlist-items-with-movies", controllers.GetAllWatchlistItemsWithMoviesByWatchListID) // GET fetch all watchlist items along with full movie data
	// expects "?watchlistID={watchlistID}" query param
	router.Delete("/api/watchlist-item", controllers.DeleteWatchlistItemByID) // DELETE watchlist item via its id
	// expects "?id={watchlistItemID}" query param
	router.Put("/api/watchlist-item-checkmarked", controllers.UpdateCheckmarkedBooleanByWatchlistItemByID) // PUT watchlist item checkmarked boolean update
	// expects "?id={watchlistItemID}" query parameter + 'checkmarked' field with boolean in the JSON body

	// Watchlist-item-note resources
    router.Post("/api/watchlist-item-note", controllers.CreateWatchlistItemNote) // POST create watchlist item note for a specific watchlist item
    router.Get("/api/watchlist-item-note/{watchlistItemID}", controllers.GetWatchlistItemNoteByWatchlistItemID) // GET fetch the note for a specific watchlist item

	// Genre resources
    router.Post("/api/genre", controllers.CreateGenreDataByMovieID) // POST add genre data for a movie
    router.Get("/api/genre/{movieID}", controllers.GetGenreByMovieID) // GET the genre data for a movie

	// User resources
    router.Post("/api/user-registration", controllers.RegisterUser) // POST register a user
	router.Post("/api/user-login", controllers.HandleLogin) // POST user login
    router.Get("/api/user/{userID}", controllers.GetUserByID) // GET user by their user id
	router.Get("/api/users", controllers.GetAllUsers) // GET all users --- testing purposes only


    // TMDB movie search resources
    router.Get("/api/tmdb-search", tmdb_controllers.TMDBSearchMovieByKeyWords)
    // expects "?query={keywords+keywords..}" query paramter

    // TMDB GET movie resources
    router.Get("/api/tmdb-movie", tmdb_controllers.TMDBGetMovieByID)
    // expects "?query={movie_id}" query paramter

	return router
}
