package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/kevinpista/my-flick-list/backend/models"
)

/*
type CacheService struct {
	RedisClient *redis.Client
}
*/

const cacheExpiration = time.Minute * 10
const redisTimeout = time.Second * 4

// SET All Watchlists for User in Redis
func (c *WatchlistService) SetAllWatchlistsInCache(userID uuid.UUID, watchlists []*models.WatchlistWithItemCount) error {
	if cache == nil {
		fmt.Println("Get fail, Redis not available. Skipping cache operation")
		return errors.New("redis not available. fetch operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	// Serialize the watchlist data to JSON format for storage
	data, err := json.Marshal(watchlists)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("mfl:watchlist:all:%s", userID.String())

	err = cache.Set(ctx, key, data, cacheExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// GET ALL Watchlists for User in Redis
func (c *WatchlistService) GetAllWatchlistsFromCache(userID uuid.UUID) ([]*models.WatchlistWithItemCount, error) {
	if cache == nil {
		fmt.Println("Get fail, Redis not available. Skipping fetch operation")
		return nil, errors.New("redis not available. fetch operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	key := fmt.Sprintf("mfl:watchlist:all:%s", userID.String())
	data, err := cache.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var watchlists []*models.WatchlistWithItemCount
	err = json.Unmarshal([]byte(data), &watchlists)
	if err != nil {
		return nil, err
	}
	return watchlists, nil
}
