package fingerprint_test

import (
	"context"
	"fmt"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
)

func ExampleAsErrorResponse() {
	client := fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))

	_, _, err := client.GetEvent(context.TODO(), "invalid_event_id")

	if fpErr, ok := fingerprint.AsErrorResponse(err); ok {
		switch fpErr.Error.Code {
		case fingerprint.ErrorCodeEvent_not_found:
			fmt.Println("event not found")
		}
	}

}
