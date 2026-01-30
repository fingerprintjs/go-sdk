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

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	// Usually this data will come from your frontend app
	tags := map[string]interface{}{
		"key": "value",
	}
	suspect := false
	linkedId := "new_linked_id"
	req := fingerprint.EventUpdate{
		Suspect:  &suspect,
		LinkedId: &linkedId,
		Tags:     tags,
	}

	httpRes, err := client.UpdateEvent(context.Background(), os.Getenv("EVENT_ID"), req)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
