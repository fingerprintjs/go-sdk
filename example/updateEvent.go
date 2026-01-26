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

	httpRes, err := client.UpdateEvent(auth, os.Getenv("EVENT_ID"), req)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
