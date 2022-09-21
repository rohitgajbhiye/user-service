package service

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

type NotifyService interface {
	Notify(context.Context, string, string) error
}

type notifyService struct {
}

func NewNotifyService() NotifyService {
	return &notifyService{}
}

func (n *notifyService) Notify(ctx context.Context, accountID string, EventType string) error {
	log.Println("starting notification")
	jsonBody := []byte(`{"accountId":"123456789","eventType":"ACCOUNT_CREATED"}`)
	body := bytes.NewReader(jsonBody)
	url := os.Getenv("notify_service")
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Printf("could not create request: %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	_, doErr := client.Do(req)
	if doErr != nil {
		log.Printf("client: error making http request: %s\n", doErr)
		os.Exit(1)
	}
	return nil
}
