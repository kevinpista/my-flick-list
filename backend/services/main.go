package services

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
)

var db *sql.DB          // Pointer of the sql db struct; globally accessed by service functions to interact with DB
var cache *redis.Client // Pointer of the Redis client struct; globally accessed by service functions to interact with Redis
const dbTimeout = time.Second * 3

// make a model that contains all fields we're going to use

type Models struct {
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB, cacheClient *redis.Client) Models { // create a pool of connections. returns the model connections
	db = dbPool // assigning this dbPool to var db pointer above

	// Check if Redis is nil
	if cacheClient != nil {
		cache = cacheClient
	} else {
		cache = nil
	}
	return Models{}
}
