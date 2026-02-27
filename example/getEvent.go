package main

import (
	"context"
	"encoding/json"
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
	eventID := os.Getenv("EVENT_ID")

	response, httpRes, err := client.GetEvent(context.Background(), eventID)

	fmt.Printf("%+v\n", httpRes)

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	if response.Bot != nil {
		fmt.Printf("Got response with Botd: %v \n", response.Bot)
	}

	if response.Identification != nil {
		stringResponse, _ := json.Marshal(response.Identification)
		fmt.Printf("Got response with Identification: %s \n", string(stringResponse))

	}

	if response.Emulator != nil {
		fmt.Printf("Got response with Emulator: %v \n", response.Emulator)
	}

	if response.IPInfo != nil {
		fmt.Printf("Got response with IPInfo: %v \n", response.IPInfo)
	}

	if response.Incognito != nil {
		fmt.Printf("Got response with Incognito: %v \n", response.Incognito)
	}

	if response.RootApps != nil {
		fmt.Printf("Got response with RootApps: %v \n", response.RootApps)
	}

	if response.ClonedApp != nil {
		fmt.Printf("Got response with ClonedApp: %v \n", response.ClonedApp)
	}

	if response.FactoryResetTimestamp != nil {
		fmt.Printf("Got response with FactoryReset: %v \n", response.FactoryResetTimestamp)
	}

	if response.Jailbroken != nil {
		fmt.Printf("Got response with Jailbroken: %v \n", response.Jailbroken)
	}

	if response.Frida != nil {
		fmt.Printf("Got response with Frida: %v \n", response.Frida)
	}

	if response.IPBlockList != nil {
		fmt.Printf("Got response with IPBlockList: %v \n", response.IPBlockList)
	}

	if response.PrivacySettings != nil {
		fmt.Printf("Got response with PrivacySettings: %v \n", response.PrivacySettings)
	}

	if response.VirtualMachine != nil {
		fmt.Printf("Got response with VirtualMachine: %v \n", response.VirtualMachine)
	}

	if response.VPN != nil {
		fmt.Printf("Got response with VPN: %v \n", response.VPN)
	}

	if response.Proxy != nil {
		fmt.Printf("Got response with Proxy: %v \n", response.Proxy)
	}

	if response.Tampering != nil {
		fmt.Printf("Got response with Tampering: %v \n", response.Tampering)
	}
}
