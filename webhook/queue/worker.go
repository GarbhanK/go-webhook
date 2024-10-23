package queue

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	redisClient "github.com/garbhank/go-webhook/redis"
	"github.com/garbhank/go-webhook/sender"
)

// iterate through the webhooks in the channel queue and send on the url
func ProcessWebhooks(ctx context.Context, webhookQueue chan interface{}) {
	for payload := range webhookQueue {

		channelType := reflect.ValueOf(p).Type().Elem()

		switch channelType.Name() {
		case "PaymentPayload":
			p = redisClient.PaymentPayload{}
		case "SongPayload":
			p = redisClient.SongPayload{}
		default:
			errMsg := fmt.Sprintf("unknown type: %s", channelType.Name())
			fmt.Println(errMsg)
		}

		go func(p channelType) {
			backoffTime := time.Second
			maxBackoffTime := time.Hour
			retries := 0
			maxRetries := 5

			for {
				err := sender.SendWebhook(p.Data, p.Url, p.WebhookId)
				if err == nil {
					break
				}
				log.Println("Error sending webhook:", err)

				retries++
				if retries >= maxRetries {
					log.Println("Max retries reached! Giving up on webhook:", p.WebhookId)
					break
				}

				time.Sleep(backoffTime)

				// Double the backoff time for the next iteration, capped at the max
				backoffTime *= 2
				log.Println(backoffTime)
				if backoffTime > maxBackoffTime {
					backoffTime = maxBackoffTime
				}
			}
		}(payload)
	}
}
