package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	req := fingerprint.NewSearchEventsRequest().
		Limit(5).
		Suspect(false).
		TotalHits(1000)
	response, httpRes, err := client.SearchEvents(context.Background(), req)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	if response.TotalHits != nil {
		fmt.Printf("Total hits: %d\n", *response.TotalHits)
	}

	if response.Events != nil {
		for _, event := range response.Events {
			fmt.Printf("Got response for event: %v \n", event.EventID)
		}
	}

	// equivalent to 2026-05-21T02:25:42.063Z
	timestamp := 1779330342063
	searchWithTimestamps := fingerprint.NewSearchEventsRequest().
		Limit(5).
		Start(int64(timestamp))
	response, _, err = client.SearchEvents(context.Background(), searchWithTimestamps)
	if err != nil {
		log.Fatalf("Error with searchWithTimestamps: %s", err.Error())
	}

	fmt.Printf("found %d events after the timestamp %d\n", len(response.Events), timestamp)

	// the same timestamp (2026-05-21T02:25:42.063Z) as a time.Time value
	datetime := time.Date(2026, 5, 21, 2, 25, 42, 63_000_000, time.UTC)
	searchWithDatetime := fingerprint.NewSearchEventsRequest().
		Limit(5).
		StartDateTime(datetime)
	response, _, err = client.SearchEvents(context.Background(), searchWithDatetime)
	if err != nil {
		log.Fatalf("Error with searchWithDatetime: %s", err.Error())
	}

	fmt.Printf("found %d events after the datetime %v\n", len(response.Events), datetime)
}
