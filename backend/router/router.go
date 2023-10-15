package router

import (
	"net/http"


	"github.com/kevinpista/my-flick-list/backend/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
    router.Use(middleware.Recoverer)
    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string{"http://*", "https://*"},
        AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))

    // movie resources
    router.Get("/api/movies", controllers.GetAllMovies)
    router.Post("/api/movie/", controllers.CreateMovie)
    router.Post("/api/movie/{id}", controllers.CreateMovieById)

    // watchlist resources
    // router.Get("/api/watchlists", controllers.GetAllWatchlists) // GET all watchlists belong to specific user
    // router.Get("/api/watchlist/{id}", controllers.GetWatchByID) // GET a specific watch 
    // router.Post("/api/watchlists/", controllers.CreateWatchlist) // POST a watchlist
    

	return router
}

