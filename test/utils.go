package test

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

func assertAuthorizationHeader(t *testing.T, r *http.Request, expectedKey string) {
	apiKey := r.Header.Get("Authorization")
	assert.Equal(t, apiKey, "Bearer api_key")
}

func assertErrorResponse(t *testing.T, expectedStatusCode int, expectedErrorResponse fingerprint.ErrorResponse, res *http.Response, err error) {
	assert.Error(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expectedStatusCode, res.StatusCode)

	errorModel := err.(*fingerprint.GenericOpenAPIError).Model().(fingerprint.ErrorResponse)

	assert.IsType(t, errorModel, fingerprint.ErrorResponse{})
	assert.Equal(t, expectedErrorResponse, errorModel)
}

func readFromFileAndUnmarshal(path string, i interface{}) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &i)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMockResponse[T any](path string) T {
	var mockResponse T
	readFromFileAndUnmarshal(path, &mockResponse)
	return mockResponse
}
