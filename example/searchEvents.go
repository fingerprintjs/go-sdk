package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fingerprintjs/go-sdk"
	"github.com/joho/godotenv"
)

func main() {
	client := fingerprint.New(fingerprint.WithRegion(fingerprint.RegionUS))

	// Load environment variables
	godotenv.Load()

	// Configure authorization, in our case with API Key
	auth := context.WithValue(context.Background(), fingerprint.ContextAccessToken, os.Getenv("FINGERPRINT_API_KEY"))

	req := client.NewSearchEventsRequest(auth).Limit(5).Suspect(false).TotalHits(1000)

	response, httpRes, err := client.SearchEvents(req)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Printf("Total hits: %d\n", *response.TotalHits)

	if response.Events != nil {
		for _, event := range response.Events {
			fmt.Printf("Got response for event: %v \n", event.EventId)
		}
	}
}
