package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

// WebhookPayload defines the structure of data expected
// to be recieved from Redis, including URL, WebhookID, and relevant Data
type WebhookPayload struct {
	Url       string `json:"url"`
	WebhookId string `json:"webhookId"`
	Data      struct {
		Id      string `json:"id"`
		Payment string `json:"payment"`
		Event   string `json:"event"`
		Date    string `json:"created"`
	} `json:"data"`
}

func Subscribe(ctx context.Context, client *redis.Client, webhookQueue chan WebhookPayload) error {
	// subscribe to webhooks channel in redis
	pubSub := client.Subscribe(ctx, "payments")

	// ensure pubsub connection is closed when func exits
	defer func(pubSub *redis.PubSub) {
		if err := pubSub.Close(); err != nil {
			log.Println("Error closing PubSub:", err)
		}
	}(pubSub)

	var payload WebhookPayload

	// infinte loop to continuously recieve messages from 'webhooks' channel
	for {
		msg, err := pubSub.ReceiveMessage(ctx)
		if err != nil {
			return err
		}

		// unmarshal the JSON payload into the WebhookPayload structure
		err = json.Unmarshal([]byte(msg.Payload), &payload)
		if err != nil {
			log.Println("Error unmarshalling payload:", err)
			continue // Continue with next message if error unmarshalling
		}

		webhookQueue <- payload // Sending the payload to the channel
	}
}
