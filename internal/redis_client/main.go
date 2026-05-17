package redis_client

import (
	"os"

	"github.com/go-redis/redis"
)

var QueueName = "job-queue"

func GetRDB() *redis.Client {
	addr := os.Getenv("REDIS_HOST")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	return rdb
}
