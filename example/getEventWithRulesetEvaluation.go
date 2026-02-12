package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	// Usually this data will come from your frontend app
	eventID := os.Getenv("EVENT_ID")

	// Usually this data will come from your frontend app
	rulesetID := os.Getenv("RULESET_ID")

	response, httpRes, err := client.GetEvent(context.Background(), eventID, fingerprint.WithRulesetID(rulesetID))

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	if response.RuleAction != nil {
		fmt.Printf("Got response with RuleAction: %v \n", response.RuleAction.AdditionalProperties["type"])
	}
}
