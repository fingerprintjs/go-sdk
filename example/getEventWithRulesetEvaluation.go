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
		if response.RuleAction.EventRuleActionAllow != nil {
			fmt.Println("action is allow")
			if response.RuleAction.EventRuleActionAllow.RequestHeaderModifications != nil {
				fmt.Printf("request header modifications to set %v, to append %v, to remove %v\n",
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Set,
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Append,
					response.RuleAction.EventRuleActionAllow.RequestHeaderModifications.Remove)
			}
		} else if response.RuleAction.EventRuleActionBlock != nil {
			fmt.Printf("action is block. Body %s statusCode: %d headers: %v\n",
				*response.RuleAction.EventRuleActionBlock.Body,
				*response.RuleAction.EventRuleActionBlock.StatusCode,
				response.RuleAction.EventRuleActionBlock.Headers)
		} else {
			fmt.Println("action is unknown (should not happen)")
		}
	}
}
