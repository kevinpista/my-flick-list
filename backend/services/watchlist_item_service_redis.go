package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kevinpista/my-flick-list/backend/models"
)

// SET Single Watchlist for User in Redis
// Key = 'mfl:watchlist:single:{watchlistID}'
func (c *WatchlistItemService) SetSingleWatchlistInCache(watchlistID int, items []*models.WatchlistItemWithMovie, name string, description string) error {
	if cache == nil {
		fmt.Println("Cache set failed, Redis not available. Skipping cache operation.")
		return errors.New("redis not available. set operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	// Create a struct to hold 3 values
	type WatchlistCache struct {
		Items       []*models.WatchlistItemWithMovie `json:"items"`
		Name        string                           `json:"name"`
		Description string                           `json:"description"`
	}

	watchlistData := WatchlistCache{
		Items:       items,
		Name:        name,
		Description: description,
	}

	data, err := json.Marshal(watchlistData)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("mfl:watchlist:single:%d", watchlistID)

	err = cache.Set(ctx, key, data, cacheExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// GET Single Watchlist for User in Redis
// Key = 'mfl:watchlist:single:{watchlistID}'
func (c *WatchlistItemService) GetSingleWatchlistFromCache(watchlistID int) ([]*models.WatchlistItemWithMovie, string, string, error) {
	if cache == nil {
		fmt.Println("Cache set failed, Redis not available. Skipping cache operation.")
		return nil, "", "", errors.New("redis not available. get operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	key := fmt.Sprintf("mfl:watchlist:single:%d", watchlistID)

	// Attempt to fetch the data from Redis
	cachedData, err := cache.Get(ctx, key).Result()
	if err == redis.Nil {
		// Cache miss: Key does not exist
		return nil, "", "", nil
	} else if err != nil {
		return nil, "", "", err
	}

	// Create a struct to hold 3 values
	type WatchlistCache struct {
		Items       []*models.WatchlistItemWithMovie `json:"items"`
		Name        string                           `json:"name"`
		Description string                           `json:"description"`
	}

	// Deserialize JSON into struct
	var watchlistData WatchlistCache
	err = json.Unmarshal([]byte(cachedData), &watchlistData)
	if err != nil {
		return nil, "", "", err
	}

	return watchlistData.Items, watchlistData.Name, watchlistData.Description, nil
}

