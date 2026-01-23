package main

import (
	"log"

	"github.com/fingerprintjs/fingerprint-server-sdk-go"
)

func main() {
	// Your webhook signing secret.
	secret := "secret"

	// Request data. In real life scenario this will be the body of incoming request
	data := []byte("data")

	// Value of the "fpjs-event-signature" header.
	header := "v1=1b2c16b75bd2a870c114153ccda5bcfca63314bc722fa160d690de133ccbb9db"

	isValid := fingerprint.IsValidWebhookSignature(header, data, secret)

	if !isValid {
		log.Fatal("Invalid signature")
	}
}
