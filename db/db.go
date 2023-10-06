package db

import (
	"fmt"
	"database/sql"
	"time"

    _  "github.com/jackc/pgconn"
    _  "github.com/jackc/pgx/v4"
    _  "github.com/jackc/pgx/stdlib"
    _  "github.com/lib/pq"
	
)

type DB struct {
	DB *sql.DB
} 

var dbConn = &DB {}

const maxOpenDBConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// DSN contains all info about our db such as port, pw etc.
// Here we open our database, check for any errors, if not
// we return our DB connection which points to sql.DB
func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil{
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	err = testDB(d) // Pings our database
	if err != nil {
		return nil, err
	}
	dbConn.DB = d
	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}
	fmt.Println("*** Pinged database successfully ***")
	return nil
}