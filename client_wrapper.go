package fingerprint

import (
	"context"
	"net/http"
)

type Client struct {
	api    *APIClient
	region Region
}

func New(opts ...ConfigOption) *Client {
	cfg := NewConfiguration()
	c := &Client{
		api:    NewAPIClient(cfg),
		region: RegionUS,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) withRegion(ctx context.Context) context.Context {
	if ctx.Value(ContextServerIndex) != nil {
		return ctx
	}
	return WithRegionContext(ctx, c.api.cfg, c.region)
}

/*
GetEvent Get an event by event ID. See (FingerprintAPIService.GetEvent) for details.
*/
func (c *Client) GetEvent(ctx context.Context, eventId string) (*Event, *http.Response, error) {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.GetEvent(ctx, eventId).Execute()
}

/*
CreateSearchEventsRequest Create a search event request. See FingerprintAPIService.SearchEvents for details.
*/
func (c *Client) CreateSearchEventsRequest(ctx context.Context) ApiSearchEventsRequest {
	ctx = c.withRegion(ctx)
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
	return c.api.FingerprintAPI.UpdateEvent(ctx, eventId).EventUpdate(eventUpdateReq).Execute()
}

/*
DeleteVisitorData Delete data by visitor ID. See FingerprintAPIService.DeleteVisitorData for details.
*/
func (c *Client) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	return c.api.FingerprintAPI.DeleteVisitorData(ctx, visitorId).Execute()
}
