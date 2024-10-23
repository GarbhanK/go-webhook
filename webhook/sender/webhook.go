package sender

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type PaymentPayload struct {
	Event   string
	Date    string
	Id      string
	Payment string
}

type SongPayload struct {
	Event     string
	Date      string
	Id        string
	SongTitle string
}

// sends a JSON POST req to the specified URL and updates event status in the db
func SendWebhook(data interface{}, url string, webhookId string) error {
	// marshal data into json
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// prepare the webhook request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the webhook request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// determine status based on response code
	status := "failed"
	if resp.StatusCode == http.StatusOK {
		status = "delivered"
	}

	log.Println(status)

	if status == "failed" {
		return errors.New(status)
	}

	return nil
}
