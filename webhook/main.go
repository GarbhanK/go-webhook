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
	paymentQueue := make(chan redisClient.PaymentPayload, 100)
	songQueue := make(chan redisClient.SongPayload, 100)

	go queue.ProcessWebhooks(ctx, paymentQueue)
	go queue.ProcessWebhooks(ctx, songQueue)

	// subscribe to 'transactions' channel
	err = redisClient.Subscribe(ctx, client, paymentQueue)
	if err != nil {
		log.Println("Error subscribing:", err)
	}

	err = redisClient.Subscribe(ctx, client, songQueue, reflect.TypeOf(redisClient.PaymentPayload{}))

	// create infinite loop to keep the program running
	select {}
}
