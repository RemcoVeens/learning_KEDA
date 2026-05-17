package main

import (
	"fmt"
	"keda/internal/redis_client"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis_client.GetRDB()
	defer rdb.Close()
	for {
		messages, err := rdb.BLPop(1*time.Second, redis_client.QueueName).Result()

		if err != nil {
			if err == redis.Nil {
				// No message, continue to the next iteration to try again
				continue
			}
			fmt.Printf("Error during BLPop: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// A message was received. messages will be like ["my-queue", "the message"]
		if len(messages) == 2 {
			queueName := messages[0]
			messageContent := messages[1]
			fmt.Printf("Received message from %s: %s\n", queueName, messageContent)

			// Simulate processing work by sleeping for 1 second
			time.Sleep(1 * time.Second)
		} else {
			// This case should ideally not happen for a successful BLPop result,
			// but it's a good practice to handle unexpected formats.
			fmt.Printf("Unexpected BLPop result format: %v\n", messages)
		}
	}
}
