package test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestUsesCorrectEndpointForRegion(t *testing.T) {
	regionsMap := make(map[fingerprint.Region]string)
	regionsMap[fingerprint.RegionUS] = "api.fpjs.io"
	regionsMap[fingerprint.RegionAsia] = "ap.api.fpjs.io"
	regionsMap[fingerprint.RegionEU] = "eu.api.fpjs.io"

	for region, endpoint := range regionsMap {
		var capturedURL *url.URL

		client := fingerprint.New(
			fingerprint.WithRegion(region),
			fingerprint.WithHTTPClient(&http.Client{
				Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					capturedURL = req.URL

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{}`)),
						Header:     make(http.Header),
						Request:    req}, nil
				}),
			}))
		_, _, _ = client.GetEvent(context.Background(), "123")
		assert.NotNil(t, capturedURL)
		if capturedURL != nil {
			assert.Equal(t, endpoint, capturedURL.Host)
			assert.Equal(t, "/v4/events/123", capturedURL.Path)
		}
	}

}
