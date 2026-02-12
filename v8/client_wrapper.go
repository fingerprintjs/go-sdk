package fingerprint

import (
	"context"
	"fmt"
	"net/http"
)

const Version = "8.0.0-test.3"

var IntegrationInfo = fmt.Sprintf(`fingerprint-pro-server-go-sdk/%s`, Version)

type sdkIdentTransport struct {
	base http.RoundTripper
}

func (t *sdkIdentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clonedReq := req.Clone(req.Context())
	query := clonedReq.URL.Query()
	query.Add("ii", IntegrationInfo)
	clonedReq.URL.RawQuery = query.Encode()

	return t.base.RoundTrip(clonedReq)
}

type Client struct {
	api     *APIClient
	region  Region
	apiKey  string
	baseURL string
}

func New(opts ...ConfigOption) *Client {
	cfg := NewConfiguration()
	httpClient := &http.Client{
		Transport: &sdkIdentTransport{base: http.DefaultTransport},
	}
	cfg.HTTPClient = httpClient
	c := &Client{
		api:    NewAPIClient(cfg),
		region: RegionUS,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.baseURL != "" {
		cfg.Servers = ServerConfigurations{
			{URL: c.baseURL},
		}
	}

	return c
}

func (c *Client) withRegion(ctx context.Context) context.Context {
	if ctx.Value(ContextServerIndex) != nil {
		return ctx
	}
	return WithRegionContext(ctx, c.api.GetConfig(), c.region)
}

func (c *Client) withAPIKey(ctx context.Context) context.Context {
	if ctx.Value(ContextAccessToken) != nil {
		return ctx
	}

	return context.WithValue(ctx, ContextAccessToken, c.apiKey)
}

type GetEventOption func(request *ApiGetEventRequest)

func WithRulesetID(rulesetID string) GetEventOption {
	return func(request *ApiGetEventRequest) {
		*request = request.RulesetId(rulesetID)
	}
}

/*
GetEvent Get an event by event ID. See FingerprintAPIService.GetEvent for details.
*/
func (c *Client) GetEvent(ctx context.Context, eventID string, opts ...GetEventOption) (*Event, *http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)

	req := c.api.FingerprintAPI.GetEvent(ctx, eventID)

	for _, opt := range opts {
		opt(&req)
	}

	return req.Execute()
}

/*
NewSearchEventsRequest Create a search event request. See FingerprintAPIService.SearchEvents for details.
*/
func (c *Client) NewSearchEventsRequest(ctx context.Context) ApiSearchEventsRequest {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.SearchEvents(ctx)
}

/*
SearchEvents Send a search event request. See FingerprintAPIService.SearchEvents for details.
*/
func (c *Client) SearchEvents(req ApiSearchEventsRequest) (*EventSearch, *http.Response, error) {
	return req.Execute()
}

/*
UpdateEvent Update an event. See FingerprintAPIService.UpdateEvent for details.
*/
func (c *Client) UpdateEvent(ctx context.Context, eventId string, eventUpdateReq EventUpdate) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.UpdateEvent(ctx, eventId).EventUpdate(eventUpdateReq).Execute()
}

/*
DeleteVisitorData Delete data by visitor ID. See FingerprintAPIService.DeleteVisitorData for details.
*/
func (c *Client) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.DeleteVisitorData(ctx, visitorId).Execute()
}
