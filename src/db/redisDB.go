package db

import (
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisConnect --
func (r *RestDB) RedisConnect(store int) *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass,
		DB:       store,
	})

	return client
}
