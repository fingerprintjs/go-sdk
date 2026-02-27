package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

var (
	trueValue     = true
	linkedIDValue = "linked_id"
)

func TestUpdateEvent(t *testing.T) {
	t.Run("Returns 200 on success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, `{"linked_id":"linked_id","suspect":true}`+"\n", string(body))

			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPatch)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		res, err := client.UpdateEvent(context.Background(), "123", fingerprint.EventUpdate{
			LinkedID: &linkedIDValue,
			Tags:     nil,
			Suspect:  &trueValue,
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("Update with empty body", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{}\n", string(body))

			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPatch)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		res, err := client.UpdateEvent(context.Background(), "123", *fingerprint.NewEventUpdate())

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("Update with empty tag", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{}\n", string(body))

			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPatch)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		res, err := client.UpdateEvent(context.Background(), "123", fingerprint.EventUpdate{
			Tags: nil,
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("Update with empty tag struct", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, `{"tags":{}}`+"\n", string(body))

			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPatch)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		res, err := client.UpdateEvent(context.Background(), "123", fingerprint.EventUpdate{
			Tags: map[string]interface{}{},
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("Update with just suspect=false", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			body, err := io.ReadAll(r.Body)
			assert.NoError(t, err)

			assert.Equal(t, "{\"suspect\":false}\n", string(body))

			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/events/123")
			assert.Equal(t, r.Method, http.MethodPatch)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		suspect := false

		res, err := client.UpdateEvent(context.Background(), "123", fingerprint.EventUpdate{
			Suspect: &suspect,
		})

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
	})

	t.Run("ErrorResponse", func(t *testing.T) {
		testCases := []struct {
			StatusCode       int
			MockResponsePath string
		}{
			{404, "mocks/errors/404_request_not_found.json"},
			{400, "mocks/errors/400_request_body_invalid.json"},
			{403, "mocks/errors/403_secret_api_key_required.json"},
			{409, "mocks/errors/409_state_not_ready.json"},
		}

		for _, testCase := range testCases {
			t.Run(fmt.Sprintf("Returns ErrorResponse on %d", testCase.StatusCode), func(t *testing.T) {
				mockResponse := GetMockResponse[fingerprint.ErrorResponse](testCase.MockResponsePath)
				ts := httptest.NewServer(http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request,
				) {
					integrationInfo := r.URL.Query().Get("ii")
					assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
					assert.Equal(t, r.URL.Path, "/events/123")
					assert.Equal(t, r.Method, http.MethodPatch)

					assertAuthorizationHeader(t, r, "api_key")

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(testCase.StatusCode)
					err := json.NewEncoder(w).Encode(mockResponse)
					if err != nil {
						panic(err)
					}
				}))
				defer ts.Close()

				client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

				res, err := client.UpdateEvent(context.Background(), "123", fingerprint.EventUpdate{
					LinkedID: &linkedIDValue,
					Tags:     nil,
					Suspect:  &trueValue,
				})

				assertErrorResponse(t, testCase.StatusCode, mockResponse, res, err)
			})
		}
	})
}
