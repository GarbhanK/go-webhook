package main

import (
	"context"
	"log"
	"os"
	"reflect"

	"github.com/garbhank/go-webhook/queue"
	redisClient "github.com/garbhank/go-webhook/redis"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}

	log.Println("using redis address:", redisAddress)

	// Initialise the Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	// Ping Redis to check the connection
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Failed to connect to Redis:", err)
	}
	log.Println("Connected to Redis:", pong)

	// create channels to act as the queue
	payloadQueue := make(chan interface{}, 100)

	go queue.ProcessWebhooks(ctx, payloadQueue)

	// subscribe to 'transactions' channel
	err = redisClient.Subscribe(ctx, client, payloadQueue)
	if err != nil {
		log.Println("Error subscribing:", err)
	}

	err = redisClient.Subscribe(ctx, client, songQueue)

	// create infinite loop to keep the program running
	select {}
}

