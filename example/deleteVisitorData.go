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
	// Load environment variables
	godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	// Usually this data will come from your frontend app
	visitorID := os.Getenv("VISITOR_ID")

	// Delete visitor data. If you are interested in using this API, please contact our support team (https://fingerprint.com/support/) to activate it for you
	httpRes, err := client.DeleteVisitorData(context.Background(), visitorID)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

}
