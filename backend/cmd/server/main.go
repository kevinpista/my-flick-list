package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/kevinpista/my-flick-list/backend/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct{
	Config Config
	// TODO - add models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("Backend server is now listening on port", port)

	srv := &http.Server {
		Addr: fmt.Sprintf(":%s", port),
		// TODO - add router
	}
	return srv.ListenAndServe()
}

func main () {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	cfg := Config {
		Port: os.Getenv("PORT"),
	}

	app := &Application {
		Config: cfg,
		// TODO - add models
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}