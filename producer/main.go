package main

import (
	"fmt"
	"keda/internal/redis_client"
	"time"
)

func main() {
	rdb := redis_client.GetRDB()
	defer rdb.Close()
	for i := range 200 { // Pushing 200 messages, as specified (100 or 200)
		msg := fmt.Sprintf("job-%d-%x", i+1, time.Now().UnixNano())
		listLen, err := rdb.LPush(redis_client.QueueName, msg).Result() // Using LPUSH as specified (LPUSH or RPUSH)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("pushed: %s, listLen: %d\n", msg, listLen)
		time.Sleep(10 * time.Millisecond) // Maintain original sleep to simulate a steady rate
	}
	list, err := rdb.LRange(redis_client.QueueName, 0, -1).Result()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("list: %v\n", list)
}
