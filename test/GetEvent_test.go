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

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

func TestGetEvent(t *testing.T) {
	t.Run("Returns event", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.Event]("mocks/events/get_event_200.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			assertAuthorizationHeader(t, r, "api_key")

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
		mockResponse := GetMockResponse[fingerprint.Event]("fixtures/get_event_200_extra_fields.json")

		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")

			assertAuthorizationHeader(t, r, "api_key")

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

	t.Run("ErrorResponse", func(t *testing.T) {
		testCases := []struct {
			StatusCode       int
			MockResponsePath string
		}{
			{400, "mocks/errors/400_visitor_id_invalid.json"},
			{403, "mocks/errors/403_subscription_not_active.json"},
			{404, "mocks/errors/404_event_not_found.json"},
			{429, "mocks/errors/429_too_many_requests.json"},
		}

		for _, testCase := range testCases {
			statusCode := testCase.StatusCode
			mockPath := testCase.MockResponsePath

			t.Run(fmt.Sprintf("Returns ErrorResponse on %d error", statusCode), func(t *testing.T) {
				mockResponse := GetMockResponse[fingerprint.ErrorResponse](mockPath)

				ts := httptest.NewServer(http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request,
				) {
					integrationInfo := r.URL.Query().Get("ii")
					assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
					assert.Equal(t, r.URL.Path, "/events/123")

					assertAuthorizationHeader(t, r, "api_key")

					w.Header().Set("Content-Type", "application/json")

					w.WriteHeader(statusCode)

					err := json.NewEncoder(w).Encode(mockResponse)
					if err != nil {
						log.Fatal(err)
					}
				}))
				defer ts.Close()

				client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

				event, res, err := client.GetEvent(context.Background(), "123")

				assert.Nil(t, event)
				assertErrorResponse(t, statusCode, mockResponse, res, err)
			})

		}
	})

	t.Run("Returns response when JSON field type is not matching response schema", func(t *testing.T) {
		mockMalformedResponse, err := os.ReadFile("fixtures/get_event_200_field_type_mismatch.json")
		if err != nil {
			log.Fatal(err)
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			parseErr := r.ParseForm()
			assert.Nil(t, parseErr)

			assertAuthorizationHeader(t, r, "api_key")

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

	t.Run("Returns event with ruleset evaluation", func(t *testing.T) {
		mockResponse := GetMockResponse[fingerprint.Event]("fixtures/get_event_200_with_ruleset.json")
		eventID := "123"
		rulesetID := "some_ruleset_id"

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, r.URL.Path, fmt.Sprintf("/events/%s", eventID))
			assert.Equal(t, r.URL.Query().Get("ruleset_id"), rulesetID)

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(mockResponse)
			if err != nil {
				log.Fatal(err)
			}
		}))

		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
		event, _, err := client.GetEvent(context.Background(), eventID, fingerprint.WithRulesetID(rulesetID))

		assert.Nil(t, err)
		assert.NotNil(t, event)
		assert.Equal(t, *event, mockResponse)
	})

	t.Run("Returns event with ruleset evaluation with unexpected type", func(t *testing.T) {
		rawJSON, err := os.ReadFile("fixtures/get_event_200_with_ruleset_extra_rule_action.json")
		if err != nil {
			log.Fatal(err)
		}

		eventID := "123"
		rulesetID := "some_ruleset_id"

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(rawJSON)
			if err != nil {
				t.Fatalf("failed to write mock response: %v", err)
			}
		}))

		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))
		event, _, err := client.GetEvent(context.Background(), eventID, fingerprint.WithRulesetID(rulesetID))

		assert.Nil(t, err)
		assert.NotNil(t, event)

		var originalEvent fingerprint.Event
		err = json.Unmarshal(rawJSON, &originalEvent)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, &originalEvent, event)
		assert.Nil(t, event.RuleAction.EventRuleActionAllow)
		assert.Nil(t, event.RuleAction.EventRuleActionBlock)
	})
}
