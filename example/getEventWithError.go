package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fingerprintjs/go-sdk"
	"github.com/joho/godotenv"
)

func exampleWrongAPIKey() {
	fmt.Println("Running 'exampleWrongAPIKey'...")

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey("apiKey"))
	_, _, err := client.GetEvent(context.Background(), "eventId")
	if err != nil {
		if errResp, ok := fingerprint.AsErrorResponse(err); ok {
			switch errResp.Error.Code {
			case fingerprint.ErrorCodeEvent_not_found:
				fmt.Println("event not found")
			case fingerprint.ErrorCodeSecret_api_key_not_found:
				fmt.Println("secret api key not found")
			default:
				fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
			}
		} else {
			fmt.Printf("other error occured: %v\n", err.Error())
		}
	}
}

func exampleWrongEventId() {
	fmt.Println("Running 'exampleWrongEventId'...")
	// Load environment variables
	godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))
	eventID := "wrongEventID"

	_, _, err := client.GetEvent(context.Background(), eventID)
	if err != nil {
		if errResp, ok := fingerprint.AsErrorResponse(err); ok {
			switch errResp.Error.Code {
			case fingerprint.ErrorCodeEvent_not_found:
				fmt.Println("event not found")
			case fingerprint.ErrorCodeSecret_api_key_not_found:
				fmt.Println("secret api key not found")
			default:
				fmt.Printf("unexpected error: %s\n", errResp.Error.Message)
			}
		} else {
			fmt.Printf("other error occured: %v\n", err.Error())
		}
	}
}

func main() {
	exampleWrongAPIKey()
	exampleWrongEventId()
}
