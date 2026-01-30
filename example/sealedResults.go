package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
)

func base64Decode(input string) []byte {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[sealedResults error] Error while loading env", err)
		return
	}

	sealedResult := base64Decode(os.Getenv("BASE64_SEALED_RESULT"))
	key := base64Decode(os.Getenv("BASE64_SEALED_RESULT_KEY"))

	keys := []fingerprint.DecryptionKey{
		{
			Key:       key,
			Algorithm: fingerprint.AlgorithmAES256GCM,
		},
	}
	unsealedResponse, err := fingerprint.UnsealEventsResponse(sealedResult, keys)
	if err != nil {
		fmt.Println("Unseal error:", err)
		return
	}

	var response = *unsealedResponse
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

	if response.IpInfo != nil {
		fmt.Printf("Got response with IpInfo: %v \n", response.IpInfo)
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

	if response.IpBlocklist != nil {
		fmt.Printf("Got response with IpBlocklist: %v \n", response.IpBlocklist)
	}

	if response.PrivacySettings != nil {
		fmt.Printf("Got response with PrivacySettings: %v \n", response.PrivacySettings)
	}

	if response.VirtualMachine != nil {
		fmt.Printf("Got response with VirtualMachine: %v \n", response.VirtualMachine)
	}

	if response.Vpn != nil {
		fmt.Printf("Got response with Vpn: %v \n", response.Vpn)
	}

	if response.Proxy != nil {
		fmt.Printf("Got response with Proxy: %v \n", response.Proxy)
	}

	if response.Tampering != nil {
		fmt.Printf("Got response with Tampering: %v \n", response.Tampering)
	}
}
