package services

import (
	"database/sql"
	"time"
)

var db *sql.DB // pointer of the sql db struct
const dbTimeout = time.Second * 3

// make a models that contains all fields we're going to use

type Models struct {
	Movie Movie
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models{ // create a pool of connections. returns the model connections
	db = dbPool // assigning this dbPool to var db pointer above
	return Models{}
}