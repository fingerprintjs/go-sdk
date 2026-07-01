package test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

func TestUnsupportedEnumValueDeserialization(t *testing.T) {
	runTest := func(
		t *testing.T,
		responseMutator func(responseBody map[string]interface{}),
	) *fingerprint.Event {
		rawJSON, err := os.ReadFile("mocks/events/get_event_200.json")
		if err != nil {
			log.Fatal(err)
		}

		var responseBody map[string]interface{}
		err = json.Unmarshal(rawJSON, &responseBody)
		if err != nil {
			log.Fatal(err)
		}

		responseMutator(responseBody)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.URL.Path, "/events/123")

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(responseBody)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "123")

		assert.Nil(t, err)
		assert.NotNil(t, event)

		return event
	}

	t.Run("Returns event with unknown proxy_type", func(t *testing.T) {
		setProxyType := func(responseBody map[string]interface{}) {
			responseBody["proxy_details"].(map[string]interface{})["proxy_type"] = "unknown-value"
		}

		event := runTest(t, setProxyType)
		assert.NotNil(t, event.ProxyDetails)
		assert.Equal(t, "unknown-value", event.ProxyDetails.ProxyType)
	})

	t.Run("Returns event with unknown bot result", func(t *testing.T) {
		setBotResult := func(responseBody map[string]interface{}) {
			responseBody["bot"] = "unknown-value"
		}

		event := runTest(t, setBotResult)
		assert.NotNil(t, event.Bot)
		assert.Equal(t, "unknown-value", string(*event.Bot))
	})
}
