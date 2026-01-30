package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
)

func exampleWrongAPIKey() {
	fmt.Println("Running 'exampleWrongAPIKey'...")

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.RegionUS), fingerprint.WithAPIKey("apiKey"))
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

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.RegionUS), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))
	eventId := "wrongEventId"

	_, _, err := client.GetEvent(context.Background(), eventId)
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
