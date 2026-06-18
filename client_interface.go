package fingerprint

import (
	"context"
	"net/http"
)

// ClientInterface defines the contract for interacting with the Fingerprint API.
// Use WithClientInterface to inject a custom implementation.
type ClientInterface interface {
	GetEvent(ctx context.Context, eventID string, opts ...GetEventOption) (*Event, *http.Response, error)
	SearchEvents(ctx context.Context, req SearchEventRequest) (*EventSearch, *http.Response, error)
	UpdateEvent(ctx context.Context, eventId string, eventUpdateReq EventUpdate) (*http.Response, error)
	DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error)
}

var _ ClientInterface = (*Client)(nil)
