package repository

import (
	"app/backend/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rdb *redis.Client
	ctx context.Context
	ttl time.Duration
}

// NewRedisDb creates a new Redis client and checks if it can connect to the Redis server.
// It returns a RedisClient struct and an error if there was any issue with connecting.
//
// cfg: A RedisConfig struct containing the Redis server's host and port.
// Returns: A RedisClient struct and an error.
func NewRedisDb(cfg config.RedisConfig) (*RedisClient, error) {
	// Create a new Redis client.
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})

	// Create a context for the Redis client.
	ctx := context.Background()

	// Check if the Redis client can connect to the server.
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	// Return a new Redis client.
	return &RedisClient{
		rdb: rdb,
		ctx: ctx,
		ttl: time.Duration(cfg.TTL) * time.Minute,
	}, nil

}
