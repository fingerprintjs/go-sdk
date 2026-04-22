package test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestFingerprint(t *testing.T) {
	t.Run("Create with empty config", func(t *testing.T) {
		client := fingerprint.New()

		assert.NotNil(t, client)
	})

	t.Run("WithBaseURL", func(t *testing.T) {
		var capturedURL *url.URL

		client := fingerprint.New(
			fingerprint.WithBaseURL("https://example.org"),
			fingerprint.WithRegion(fingerprint.RegionEU),
			fingerprint.WithHTTPClient(&http.Client{
				Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					capturedURL = req.URL

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{}`)),
						Header:     make(http.Header),
						Request:    req,
					}, nil
				}),
			}),
		)

		_, _, _ = client.GetEvent(context.Background(), "123")
		assert.NotNil(t, capturedURL)
		if capturedURL != nil {
			assert.Equal(t, "example.org", capturedURL.Host)
			assert.Equal(t, "/events/123", capturedURL.Path)
		}
	})

	t.Run("WithRegion", func(t *testing.T) {
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
							Request:    req,
						}, nil
					}),
				}))
			_, _, _ = client.GetEvent(context.Background(), "123")
			assert.NotNil(t, capturedURL)
			if capturedURL != nil {
				assert.Equal(t, endpoint, capturedURL.Host)
				assert.Equal(t, "/v4/events/123", capturedURL.Path)
			}
		}
	})

	t.Run("WithHTTPClient", func(t *testing.T) {
		const customIIParam = "custom_ii_param"
		const customHTTPHeaderKey = "custom_http_header_key"
		const customHTTPHeaderValue = "custom_http_header_value"

		httpClient := http.Client{
			Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
				clonedReq := req.Clone(req.Context())
				query := clonedReq.URL.Query()
				query.Add("ii", customIIParam)
				clonedReq.URL.RawQuery = query.Encode()

				clonedReq.Header.Set(customHTTPHeaderKey, customHTTPHeaderValue)

				return http.DefaultTransport.RoundTrip(clonedReq)
			}),
		}

		t.Run("Preserves customer transport behavior while adding SDK ii", func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				iiParams, ok := r.URL.Query()["ii"]
				require.True(t, ok)

				var integrationInfo = fmt.Sprintf("fingerprint-pro-server-go-sdk/%s", fingerprint.Version)
				var integrationInfoParamExists = false
				var customParamExists = false
				for _, param := range iiParams {
					if param == integrationInfo {
						integrationInfoParamExists = true
					}
					if param == customIIParam {
						customParamExists = true
					}
				}

				assert.True(t, integrationInfoParamExists)
				assert.True(t, customParamExists)

				customHTTPHeaderV := r.Header.Get(customHTTPHeaderKey)
				assert.Equal(t, customHTTPHeaderV, customHTTPHeaderValue)

				w.WriteHeader(http.StatusOK)
			}))

			defer ts.Close()

			_, _, err := fingerprint.New(fingerprint.WithBaseURL(ts.URL), fingerprint.WithHTTPClient(&httpClient)).
				GetEvent(context.Background(), "123")
			assert.NoError(t, err)
		})
	})
}
