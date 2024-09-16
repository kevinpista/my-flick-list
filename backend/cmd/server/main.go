package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/kevinpista/my-flick-list/backend/cache"
	"github.com/kevinpista/my-flick-list/backend/db"
	"github.com/kevinpista/my-flick-list/backend/router"
	"github.com/kevinpista/my-flick-list/backend/services"
)

type Config struct {
	Port      string
	RedisHost string
	RedisPort string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	// LOCAL DEV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// LOCAL DEV
	port := os.Getenv("PORT")
	fmt.Println("Backend server is now listening on port", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}
	return srv.ListenAndServe()
}

func main() {
	// LOCAL DEV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// LOCAL DEV

	cfg := Config{
		Port:      os.Getenv("PORT"),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}

	// Initialize DB
	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database. Error:", err)
	}

	defer dbConn.DB.Close()

	// Initialize Redis Cache
	cacheConn, err := cache.ConnectRedis(cfg.RedisHost, cfg.RedisPort)
	if err != nil {
		log.Fatal("Cannot connect to Redis cache. Error:", err)
	}

	// Initialize Models with both DB and Redis
	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB, cacheConn.Client),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
