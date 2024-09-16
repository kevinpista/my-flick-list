package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type Cache struct {
	Client *redis.Client
}

var cacheConn = &Cache{}

const redisTimeout = time.Second * 4

// ConnectRedis sets up and returns a Redis client connection
func ConnectRedis(redisHost, redisPort string) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	// Ping Redis to ensure connection is established
	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("*** Pinged Redis cache successfully ***")
	cacheConn.Client = client
	return cacheConn, nil
}
