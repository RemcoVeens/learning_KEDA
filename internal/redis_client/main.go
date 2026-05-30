package redis_client

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var QueueName = "job-queue"
var CompletionQueueName = "completion-queue"
var (
	Ctx  = context.Background()
	rdb  *redis.Client
	once sync.Once
)

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

func Flush(rdb *redis.Client) {
	iter := rdb.Scan(Ctx, 0, "*", 0).Iterator()

	deletedCount := 0
	for iter.Next(Ctx) {
		key := iter.Val()
		// 2. Delete them one by one (or batch them in a pipeline for speed)
		rdb.Del(Ctx, key)
		deletedCount++
	}

	if err := iter.Err(); err != nil {
		panic(fmt.Sprintf("Failed to scan: %v", err))
	}

	fmt.Printf("Successfully deleted %d keys.\n", deletedCount)
}
