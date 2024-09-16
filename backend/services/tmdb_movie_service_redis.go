package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kevinpista/my-flick-list/backend/models"
)

// SET Movie in Redis
// Key = 'mfl:movie:{movieID}'
func (c *TMDBMovieService) SetMovieInCache(movieID string, movie *models.TMDBMovie) error {
	if cache == nil {
		fmt.Println("Cache set failed, Redis not available. Skipping cache operation.")
		return errors.New("redis not available. set operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	// Serialize data into JSON
	data, err := json.Marshal(movie)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("mfl:movie:%s", movieID)

	err = cache.Set(ctx, key, data, cacheExpiration).Err()
	if err != nil {
		return err
	}

	return nil
}

// GET Movie from Redis
// Key = 'mfl:movie:{movieID}'
func (c *TMDBMovieService) GetMovieFromCache(movieID string) (*models.TMDBMovie, error) {
	if cache == nil {
		fmt.Println("Get fail, Redis not available. Skipping cache operation.")
		return nil, errors.New("redis not available. get operation skipped")
	}

	ctx, cancel := context.WithTimeout(context.Background(), redisTimeout)
	defer cancel()

	key := fmt.Sprintf("mfl:movie:%s", movieID)

	data, err := cache.Get(ctx, key).Result()
	if err == redis.Nil {
		// Cache miss, return nil without error
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// Deserialize JSON into struct
	var movie models.TMDBMovie
	err = json.Unmarshal([]byte(data), &movie)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
