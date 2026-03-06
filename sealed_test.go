package fingerprint_test

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
)

func ExampleUnsealEventsResponse() {
	// Utility function to decode base64 string
	base64Decode := func(input string) []byte {
		output, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			log.Fatal(err)
		}
		return output
	}

	// Sealed result from the frontend.
	sealedResult := base64Decode(os.Getenv("BASE64_SEALED_RESULT"))
	// Base64 encoded key generated in the dashboard.
	key := base64Decode(os.Getenv("BASE64_SEALED_RESULT_KEY"))

	keys := []fingerprint.DecryptionKey{
		// You can provide more than one key to support key rotation. The SDK will try to decrypt the result with each key.
		{
			Key:       key,
			Algorithm: fingerprint.AlgorithmAES256GCM,
		},
	}
	unsealedResponse, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

	if err != nil {
		panic(err)
	}

	// Do something with unsealed response, e.g: send it back to the frontend.
	fmt.Println(unsealedResponse)
}
