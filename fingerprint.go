package fingerprint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

type API = openapi.FingerprintAPI

var integrationInfo = fmt.Sprintf(`fingerprint-pro-server-go-sdk/%s`, Version)

type sdkIdentTransport struct {
	base http.RoundTripper
}

func (t *sdkIdentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clonedReq := req.Clone(req.Context())
	query := clonedReq.URL.Query()
	query.Add("ii", integrationInfo)
	clonedReq.URL.RawQuery = query.Encode()

	return t.base.RoundTrip(clonedReq)
}

// Client is the main client for interacting with the Fingerprint API.
type Client struct {
	api     *openapi.APIClient
	region  Region
	apiKey  string
	baseURL string
}

// New creates a new Fingerprint API client with the given configuration options.
func New(opts ...ConfigOption) *Client {
	cfg := openapi.NewConfiguration()
	httpClient := &http.Client{
		Transport: &sdkIdentTransport{base: http.DefaultTransport},
	}
	cfg.HTTPClient = httpClient
	c := &Client{
		api:    openapi.NewAPIClient(cfg),
		region: RegionUS,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.baseURL != "" {
		cfg.Servers = openapi.ServerConfigurations{
			{URL: c.baseURL},
		}
	}

	return c
}

func (c *Client) withRegion(ctx context.Context) context.Context {
	if ctx.Value(openapi.ContextServerIndex) != nil {
		return ctx
	}
	return WithRegionContext(ctx, c.api.GetConfig(), c.region)
}

func (c *Client) withAPIKey(ctx context.Context) context.Context {
	if ctx.Value(openapi.ContextAccessToken) != nil {
		return ctx
	}

	return context.WithValue(ctx, openapi.ContextAccessToken, c.apiKey)
}

// GetEventOption is a functional option for configuring [Client.GetEvent] requests.
type GetEventOption func(request *openapi.ApiGetEventRequest)

// WithRulesetID sets the ruleset ID for the [Client.GetEvent] request.
func WithRulesetID(rulesetID string) GetEventOption {
	return func(request *openapi.ApiGetEventRequest) {
		*request = request.RulesetID(rulesetID)
	}
}

// GetEvent retrieves an event by event ID. See [openapi.FingerprintAPI.GetEvent] for details.
func (c *Client) GetEvent(ctx context.Context, eventID string, opts ...GetEventOption) (*openapi.Event, *http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)

	req := c.api.FingerprintAPI.GetEvent(eventID)

	for _, opt := range opts {
		opt(&req)
	}

	return req.Execute(ctx)
}

// NewSearchEventsRequest creates a search event request. See [openapi.FingerprintAPI.SearchEvents] for details.
func NewSearchEventsRequest() openapi.ApiSearchEventsRequest {
	return openapi.ApiSearchEventsRequest{ApiService: nil}
}

// SearchEvents sends a search event request. See [openapi.FingerprintAPI.SearchEvents] for details.
func (c *Client) SearchEvents(ctx context.Context, req openapi.ApiSearchEventsRequest) (*openapi.EventSearch, *http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.SearchEventsExecute(ctx, req)
}

// UpdateEvent updates an event. See [openapi.FingerprintAPI.UpdateEvent] for details.
func (c *Client) UpdateEvent(ctx context.Context, eventId string, eventUpdateReq openapi.EventUpdate) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.UpdateEvent(eventId).EventUpdate(eventUpdateReq).Execute(ctx)
}

// DeleteVisitorData deletes data by visitor ID. See [openapi.FingerprintAPI.DeleteVisitorData] for details.
func (c *Client) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.DeleteVisitorData(visitorId).Execute(ctx)
}
