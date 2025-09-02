package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var cache *redis.Client

func Client() *redis.Client {
	return cache
}

func Connect() {
	ctx := context.Background()

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
		cache = redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	} else {
		redisOptions, err := redis.ParseURL(redisURL)
		if err != nil {
			log.Fatalf("failed to parse Redis URL: %v", err)
		}
		cache = redis.NewClient(redisOptions)
	}

	cmd := cache.Ping(ctx)
	if cmd.Err() != nil {
		fmt.Println("Error connecting caching database: ", cmd.Err())
		panic(cmd.Err())
	}

	fmt.Println("Successfully connected to the redis cache")
}
