package fingerprint

import (
	"context"
	"net/http"

	sdk "github.com/fingerprintjs/go-sdk/sdk"
)

var ContextAccessToken = sdk.ContextAccessToken

type EventUpdate = sdk.EventUpdate

// todo read from codegen or package.json
const IntegrationInfo = "go-sdk/1"

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
	api    *sdk.APIClient
	region Region
}

func New(opts ...ConfigOption) *Client {
	cfg := sdk.NewConfiguration()
	httpClient := &http.Client{
		Transport: &sdkIdentTransport{base: http.DefaultTransport},
	}
	cfg.HTTPClient = httpClient
	c := &Client{
		api:    sdk.NewAPIClient(cfg),
		region: RegionUS,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) withRegion(ctx context.Context) context.Context {
	if ctx.Value(sdk.ContextServerIndex) != nil {
		return ctx
	}
	return WithRegionContext(ctx, c.api.GetConfig(), c.region)
}

/*
GetEvent Get an event by event ID. See FingerprintAPIService.GetEvent for details.
*/
func (c *Client) GetEvent(ctx context.Context, eventId string) (*sdk.Event, *http.Response, error) {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.GetEvent(ctx, eventId).Execute()
}

/*
NewSearchEventsRequest Create a search event request. See FingerprintAPIService.SearchEvents for details.
*/
func (c *Client) NewSearchEventsRequest(ctx context.Context) sdk.ApiSearchEventsRequest {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.SearchEvents(ctx)
}

/*
SearchEvents Send a search event request. See FingerprintAPIService.SearchEvents for details.
*/
func (c *Client) SearchEvents(req sdk.ApiSearchEventsRequest) (*sdk.EventSearch, *http.Response, error) {
	return req.Execute()
}

/*
UpdateEvent Update an event. See FingerprintAPIService.UpdateEvent for details.
*/
func (c *Client) UpdateEvent(ctx context.Context, eventId string, eventUpdateReq sdk.EventUpdate) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.UpdateEvent(ctx, eventId).EventUpdate(eventUpdateReq).Execute()
}

/*
DeleteVisitorData Delete data by visitor ID. See FingerprintAPIService.DeleteVisitorData for details.
*/
func (c *Client) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.DeleteVisitorData(ctx, visitorId).Execute()
}
