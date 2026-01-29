package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestDeleteVisitorData(t *testing.T) {
	t.Run("Returns 200 on success", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			integrationInfo := r.URL.Query().Get("ii")
			assert.Equal(t, integrationInfo, fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version))
			assert.Equal(t, r.URL.Path, "/visitors/123")
			assert.Equal(t, r.Method, http.MethodDelete)

			assertAuthorizationHeader(t, r, "api_key")

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
		}))
		defer ts.Close()

		client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

		res, err := client.DeleteVisitorData(context.Background(), "123")

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.StatusCode, 200)
	})

	t.Run("ErrorResponse", func(t *testing.T) {
		testCases := []struct {
			StatusCode       int
			MockResponsePath string
		}{
			{404, "mocks/errors/404_visitor_not_found.json"},
			{429, "mocks/errors/429_too_many_requests.json"},
			{400, "mocks/errors/400_visitor_id_invalid.json"},
			{403, "mocks/errors/403_feature_not_enabled.json"},
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
					assert.Equal(t, r.URL.Path, "/visitors/123")
					assert.Equal(t, r.Method, http.MethodDelete)

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

				res, err := client.DeleteVisitorData(context.Background(), "123")

				assertErrorResponse(t, testCase.StatusCode, mockResponse, res, err)
			})
		}
	})
}
