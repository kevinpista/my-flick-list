package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

    "github.com/kevinpista/my-flick-list/backend/db"
	"github.com/kevinpista/my-flick-list/backend/services"
	"github.com/kevinpista/my-flick-list/backend/router"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct{
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	// LOCAL DEV
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	// LOCAL DEV
	port := os.Getenv("PORT")
	fmt.Println("Backend server is now listening on port", port)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", port),
        Handler: router.Routes(), 
	}
	return srv.ListenAndServe()
}

func main () {
	// LOCAL DEV
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	// LOCAL DEV

	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN") ///
    dbConn, err := db.ConnectPostgres(dsn)
    if err != nil {
        log.Fatal("Cannot connect to database")
    }

    defer dbConn.DB.Close()

	app := &Application {
		Config: cfg,
        Models: services.New(dbConn.DB), // creates a connection with DB to get the models whenever we want to perform CRUD
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}