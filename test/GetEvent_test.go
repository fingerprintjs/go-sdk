package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestGetEvent(t *testing.T) {
	t.Run("Returns event", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.Event]("mocks/get_event_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			apiKey := r.Header.Get("Authorization")
			assert.Equal(t, apiKey, "Bearer api_key")

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "123")

		assert.Nil(t, err)
		assert.NotNil(t, event)
		assert.Equal(t, *event, mockResponse)
	})

	t.Run("Returns event with unexpected fields", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.Event]("mocks/get_event_200_extra_fields.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			apiKey := r.Header.Get("Authorization")
			assert.Equal(t, apiKey, "Bearer api_key")

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "123")

		assert.Nil(t, err)
		assert.NotNil(t, event)
		assert.Equal(t, *event, mockResponse)
	})

	t.Run("Returns ErrorResponse on 403 error", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.ErrorResponse]("mocks/errors/403_subscription_not_active.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			apiKey := r.Header.Get("Authorization")
			assert.Equal(t, apiKey, "Bearer api_key")

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(403)

			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "123")

		assert.Error(t, err)
		assert.IsType(t, &fingerprint.GenericOpenAPIError{}, err)
		assert.Nil(t, event)

		errorModel := err.(*fingerprint.GenericOpenAPIError).Model().(fingerprint.ErrorResponse)

		assert.IsType(t, errorModel, fingerprint.ErrorResponse{})
		assert.Equal(t, mockResponse, errorModel)
	})

	t.Run("Returns response when JSON field type is not matching response schema", func(t *testing.T) {
		mockMalformedResponse, err := os.ReadFile("mocks/get_event_200_field_type_mismatch.json")
		if err != nil {
			log.Fatal(err)
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

			apiKey := r.Header.Get("Authorization")
			assert.Equal(t, apiKey, "Bearer api_key")

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(mockMalformedResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		event, _, err := client.GetEvent(context.Background(), "request_id")
		assert.Error(t, err)
		assert.NotNil(t, event)
		assert.Nil(t, event.Incognito)
	})
}
