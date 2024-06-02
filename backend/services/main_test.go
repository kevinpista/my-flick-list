package services

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
		// Alias imported package to avoid name clash
	dbpkg "github.com/kevinpista/my-flick-list/backend/db"

)

func TestMain(m *testing.M) {
    // Load test environment
    err := godotenv.Load("../.env.test")
    if err != nil {
        log.Fatalf("Error loading .env.test file: %v", err)
    }

    // Setup database connection to testDB
    dsn := os.Getenv("DSN")
    dbConn, err := dbpkg.ConnectPostgres(dsn)
    if err != nil {
        log.Fatalf("Cannot connect to test database: %v", err)
    }
    
    // Will be assigning test database to the global variable 'db' for service functions to access
    // only initiated during automated testing only
    defer dbConn.DB.Close()
    db = dbConn.DB

    // Run the tests
    os.Exit(m.Run())
}