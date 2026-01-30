package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.RegionUS), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	req := client.NewSearchEventsRequest(context.Background()).Limit(5).Suspect(false).TotalHits(1000)

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
