package main

import (
	"context"
	"log"
	"os"

	"github.com/garbhank/go-webhook/queue"
	redisClient "github.com/garbhank/go-webhook/redis"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialise the Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})

	// create a channel to act as the queue
	webhookQueue := make(chan redisClient.WebhookPayload, 100)

	go queue.ProcessWebhooks(ctx, webhookQueue)

	// subscribe to 'transactions' channel
	err := redisClient.Subscribe(ctx, client, webhookQueue)
	if err != nil {
		log.Println("Error:", err)
	}

	// create infinite loop to keep the program running
	select {}
}
