package fingerprint

import (
	"net/http"
)

// ConfigOption is a functional option for configuring the Client.
type ConfigOption func(*Client)

// WithRegion sets the region for API requests. See [Region] for available regions.
func WithRegion(region Region) ConfigOption {
	return func(c *Client) {
		c.region = region
	}
}

// WithAPIKey sets the API key for authenticating requests.
func WithAPIKey(apiKey string) ConfigOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithBaseURL sets a custom base URL for API requests.
func WithBaseURL(baseURL string) ConfigOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client for making requests.
// The client's transport will be wrapped to add SDK identification.
func WithHTTPClient(httpClient *http.Client) ConfigOption {
	return func(c *Client) {
		if httpClient == nil {
			return
		}

		// Perform a shallow copy to modify the transport without modifying the original
		cloned := *httpClient

		base := httpClient.Transport
		if base == nil {
			base = http.DefaultTransport
		}

		// Wrap the transport with our own
		cloned.Transport = &sdkIdentTransport{base: base}

		c.api.GetConfig().HTTPClient = &cloned
	}
}

// WithFingerprintAPI sets a custom Fingerprint API implementation.
func WithFingerprintAPI(fingerprintAPI API) ConfigOption {
	return func(c *Client) {
		c.api.FingerprintAPI = fingerprintAPI
	}
}

// WithDebug enables or disables debug mode for the API client.
func WithDebug(debug bool) ConfigOption {
	return func(c *Client) {
		c.api.GetConfig().Debug = debug
	}
}
