package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	fingerprint "github.com/TheUnderScorer/go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type customTransport struct {
	base http.RoundTripper
}

const customIIParam = "custom_ii_param"
const customHTTPHeaderKey = "custom_http_header_key"
const customHTTPHeaderValue = "custom_http_header_value"

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clonedReq := req.Clone(req.Context())
	query := clonedReq.URL.Query()
	query.Add("ii", customIIParam)
	clonedReq.URL.RawQuery = query.Encode()

	clonedReq.Header.Set(customHTTPHeaderKey, customHTTPHeaderValue)

	return t.base.RoundTrip(clonedReq)
}

func TestWithHTTPClient(t *testing.T) {
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

		_, _, err := fingerprint.New(fingerprint.WithBaseURL(ts.URL), fingerprint.WithHTTPClient(&http.Client{Transport: &customTransport{base: http.DefaultTransport}})).
			GetEvent(context.Background(), "123")
		assert.NoError(t, err)
	})
}
