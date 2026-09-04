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
	_ = godotenv.Load()

	client := fingerprint.New(fingerprint.WithRegion(fingerprint.Region(os.Getenv("REGION"))), fingerprint.WithAPIKey(os.Getenv("FINGERPRINT_API_KEY")))

	ipv4 := "34.162.244.71"
	req := fingerprint.EdgeRequest{
		Method:      "GET",
		URL:         "https://example.com/login",
		Ipv4Address: &ipv4,
		Headers: []fingerprint.EdgeRequestHeadersInner{
			{Name: "Host", Value: "example.com"},
			{Name: "User-Agent", Value: "Mozilla/5.0"},
			{Name: "Authorization", Value: ""},
		},
	}

	event, httpRes, err := client.AnalyzeRequestForAutomationIntelligence(context.Background(), req)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	fmt.Println(event.EventID)
}
