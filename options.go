package fingerprint

import "net/http"

type ConfigOption func(*Client)

func WithRegion(region Region) ConfigOption {
	return func(c *Client) {
		c.region = region
	}
}

func WithAPIKey(apiKey string) ConfigOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithBaseURL(baseURL string) ConfigOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

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
