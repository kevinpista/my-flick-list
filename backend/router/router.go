package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kevinpista/my-flick-list/backend/controllers"
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

	// movie resources
	router.Post("/api/movie", controllers.CreateMovieByID)
    router.Get("/api/movie/{id}", controllers.GetMovieByID)
    router.Get("/api/movies", controllers.GetAllMovies)

	// watchlist resources
	router.Get("/api/watchlists", controllers.GetAllWatchlists)     // GET all watchlists belonging to specific user
	router.Post("/api/watchlists", controllers.CreateWatchlists)    // POST a watchlist
	router.Get("/api/watchlist/{id}", controllers.GetWatchlistByID) // GET a specific watch by ID

	// watchlist-items resources
	router.Post("/api/watchlist-item", controllers.CreateWatchlistItemByWatchlistID) // POST create a watchlist item for a specific watchlist
	router.Get("/api/watchlist-items", controllers.GetAllWatchlistItemsByWatchListID) // GET fetch all watchlist items from a specific watchlist.
	// expects "?watchlistID={watchlistID}" query paramter

	// watchlist-item-note resources
    router.Post("/api/watchlist-item-note", controllers.CreateWatchlistItemNote) // POST create watchlist item note for a specific watchlist item
    router.Get("/api/watchlist-item-note/{watchlistItemID}", controllers.GetWatchlistItemNoteByWatchlistItemID) // GET fetch the note for a specific watchlist item

	// genre resources
    router.Post("/api/genre", controllers.CreateGenreDataByMovieID) // POST add genre data for a movie
    router.Get("/api/genre/{movieID}", controllers.GetGenreByMovieID) // GET the genre data for a movie
	return router
}
