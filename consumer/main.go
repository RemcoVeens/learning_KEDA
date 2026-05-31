package main

import (
	"crypto/rand"
	"fmt"
	"keda/internal/redis_client"
	"math/big"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis_client.GetRDB()
	defer rdb.Close()
	for {
		messages, err := rdb.BRPop(redis_client.Ctx, 1*time.Second, redis_client.QueueName).Result()

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

			completion_message := fmt.Sprintf("Completed processing: %s", messageContent)
			rdb.RPush(redis_client.Ctx, redis_client.CompletionQueueName, completion_message)
			randomNumber, err := rand.Int(rand.Reader, big.NewInt(5))
			if err != nil {
				fmt.Printf("Error generating random number: %v\n", err)
				continue
			}
			time.Sleep(time.Duration(randomNumber.Int64()) * time.Second)
		} else {
			// This case should ideally not happen for a successful BLPop result,
			// but it's a good practice to handle unexpected formats.
			fmt.Printf("Unexpected BLPop result format: %v\n", messages)
		}
	}
}
