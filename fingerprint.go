package fingerprint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

type API = openapi.FingerprintAPI
type APIDeleteVisitorDataRequest = openapi.ApiDeleteVisitorDataRequest
type APIGetEventRequest = openapi.ApiGetEventRequest
type APISearchEventsRequest = openapi.ApiSearchEventsRequest
type APIUpdateEventRequest = openapi.ApiUpdateEventRequest

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
	return withRegionContext(ctx, c.api.GetConfig(), c.region)
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
// Parameters:
//   - ctx: context for cancellation, deadlines, and authentication.
//   - eventID: unique [event identifier]
//   - opts: Optional functional options that modify the query parameters, it can be used to pass optional ruleset ID to evaluate. Defaults to none.
//
// Using Rulesets:
// Provide a rulesetID query parameter to evaluate the event against a ruleset. To learn more, refer to example located in [example/getEventWithRulesetEvaluation.go].
//
// [event identifier]: https://dev.fingerprint.com/reference/get-function#event_id
// [example/getEventWithRulesetEvaluation.go]: https://github.com/fingerprintjs/go-sdk/blob/main/example/getEventWithRulesetEvaluation.go
func (c *Client) GetEvent(ctx context.Context, eventID string, opts ...GetEventOption) (*Event, *http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)

	req := c.api.FingerprintAPI.GetEvent(eventID)

	for _, opt := range opts {
		opt(&req)
	}

	return req.Execute(ctx)
}

type SearchEventRequest = openapi.ApiSearchEventsRequest

// NewSearchEventsRequest creates a search event request. See [openapi.FingerprintAPI.SearchEvents] for details.
//
// The request can be configured using its builder methods to filter events by various criteria
// such as visitor ID, request ID, linked ID, IP address, and time range.
func NewSearchEventsRequest() SearchEventRequest {
	return SearchEventRequest{ApiService: nil}
}

// SearchEvents sends a search event request to retrieve a list of events matching the specified criteria.
// To build the request, use [NewSearchEventsRequest].
// See [openapi.FingerprintAPI.SearchEvents] for details.
//
// Parameters:
//   - ctx: context for cancellation, deadlines, and authentication.
//   - req: configured search request with filters (created using [NewSearchEventsRequest]).
//
// Returns a paginated list of events matching the search criteria.
func (c *Client) SearchEvents(ctx context.Context, req SearchEventRequest) (*EventSearch, *http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.SearchEventsExecute(ctx, req)
}

// UpdateEvent updates an event by event ID. See [openapi.FingerprintAPI.UpdateEvent] for details.
//
// Parameters:
//   - ctx: context for cancellation, deadlines, and authentication.
//   - eventId: unique [event identifier] to update.
//   - eventUpdateReq: event update payload containing the fields to modify.
//
// This method allows you to modify event properties such as suspect status and other metadata.
//
// [event identifier]: https://dev.fingerprint.com/reference/get-function#event_id
func (c *Client) UpdateEvent(ctx context.Context, eventId string, eventUpdateReq EventUpdate) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.UpdateEvent(eventId).EventUpdate(eventUpdateReq).Execute(ctx)
}

// DeleteVisitorData deletes all data associated with a visitor ID. See [openapi.FingerprintAPI.DeleteVisitorData] for details.
//
// Parameters:
//   - ctx: context for cancellation, deadlines, and authentication.
//   - visitorId: unique [visitor identifier] whose data should be deleted.
//
// This method permanently removes all events and data associated with the specified visitor.
// Use with caution as this operation cannot be undone.
//
// [visitor identifier]: https://dev.fingerprint.com/reference/get-function#visitor_id
func (c *Client) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	ctx = c.withRegion(ctx)
	ctx = c.withAPIKey(ctx)
	return c.api.FingerprintAPI.DeleteVisitorData(visitorId).Execute(ctx)
}
