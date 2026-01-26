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
	visitorId := os.Getenv("VISITOR_ID")

	// Delete visitor data. If you are interested in using this API, please contact our support team (https://fingerprint.com/support/) to activate it for you
	httpRes, err := client.DeleteVisitorData(auth, visitorId)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

}
