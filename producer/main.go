package main

import (
	"fmt"
	"keda/internal/redis_client"
	"time"
)

var numMessages = 261

func CreateMessages() {
	rdb := redis_client.GetRDB()
	redis_client.Flush(rdb)
	defer rdb.Close()
	for i := range numMessages {
		msg := fmt.Sprintf("job-%d-%x", i+1, time.Now().UnixNano())
		listLen, err := rdb.LPush(redis_client.Ctx, redis_client.QueueName, msg).Result() // Using LPUSH as specified (LPUSH or RPUSH)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("pushed: %s, listLen: %d\n", msg, listLen)
		time.Sleep(10 * time.Millisecond) // Maintain original sleep to simulate a steady rate
	}
	list, err := rdb.LRange(redis_client.Ctx, redis_client.QueueName, 0, -1).Result()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("found %d tasks in queue: %v\n", len(list), redis_client.QueueName)
}

func ResponseProcessor() {
	rdb := redis_client.GetRDB()
	defer rdb.Close()
	recieved_resposed := 0
	for {
		messages, err := rdb.BLPop(redis_client.Ctx, 0, redis_client.CompletionQueueName).Result()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		message := messages[1]
		fmt.Printf("message: %s\n", message)
		recieved_resposed++
		if recieved_resposed == numMessages {
			break
		}
	}
}
func main() {
	go CreateMessages()
	ResponseProcessor()
}
