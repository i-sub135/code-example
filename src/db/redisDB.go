package db

import (
	"os"

	"github.com/go-redis/redis/v8"
)

// RedisConnect --
func (r *RestDB) RedisConnect() *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")
	clien := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})

	return clien
}
