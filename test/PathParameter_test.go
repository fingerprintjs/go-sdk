package test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

type pathParameterEndpoint struct {
	Name   string
	Param  string
	Prefix string
	// Response is encoded as the stub response body, so a call that reaches
	// the server can be asserted to succeed.
	Response any
	Call     func(client *fingerprint.Client, id string) error
}

// pathParameterEndpoints returns every operation that takes an ID in the path.
func pathParameterEndpoints() []pathParameterEndpoint {
	return []pathParameterEndpoint{
		{
			Name:     "GetEvent",
			Param:    "eventID",
			Prefix:   "/events/",
			Response: GetMockResponse[fingerprint.Event]("mocks/events/get_event_200.json"),
			Call: func(client *fingerprint.Client, id string) error {
				_, _, err := client.GetEvent(context.Background(), id)

				return err
			},
		},
		{
			Name:   "UpdateEvent",
			Param:  "eventID",
			Prefix: "/events/",
			Call: func(client *fingerprint.Client, id string) error {
				_, err := client.UpdateEvent(context.Background(), id, fingerprint.EventUpdate{})

				return err
			},
		},
		{
			Name:   "DeleteVisitorData",
			Param:  "visitorID",
			Prefix: "/visitors/",
			Call: func(client *fingerprint.Client, id string) error {
				_, err := client.DeleteVisitorData(context.Background(), id)

				return err
			},
		},
	}
}

// TestPathParameterEncoding asserts that an ID travels as a single, opaque path
// segment, so a value containing slashes cannot inject extra segments.
//
// Assertions are made against the literal request target the server received,
// since net/url keeps the path it was handed and only the bytes on the wire show
// what the API would actually be asked for.
func TestPathParameterEncoding(t *testing.T) {
	ids := []struct {
		Name    string
		ID      string
		Encoded string
	}{
		{"Path traversal", "../events", "..%2Fevents"},
		{"Nested path traversal", "../../events", "..%2F..%2Fevents"},
		{"Leading slash", "/events/123", "%2Fevents%2F123"},
		{"Absolute URL", "https://domain.tld/evil", "https:%2F%2Fdomain.tld%2Fevil"},
		{"Query injection", "123?limit=1", "123%3Flimit=1"},
		{"Fragment injection", "123#fragment", "123%23fragment"},
		{"Whitespace", "hello world", "hello%20world"},
		{"Empty", "", ""},
		// Only a segment made up solely of dots is a dot-segment, so anything
		// else containing a dot keeps passing through untouched.
		{"Regular ID with a dot", "1708102555327.NLOjmg", "1708102555327.NLOjmg"},
		{"Three dots", "...", "..."},
	}

	for _, endpoint := range pathParameterEndpoints() {
		for _, id := range ids {
			t.Run(fmt.Sprintf("%s: %s", endpoint.Name, id.Name), func(t *testing.T) {
				var requestTarget, decodedPath string

				ts := httptest.NewServer(http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request,
				) {
					requestTarget, _, _ = strings.Cut(r.RequestURI, "?")
					decodedPath = r.URL.Path

					w.Header().Set("Content-Type", "application/json")

					if endpoint.Response != nil {
						if err := json.NewEncoder(w).Encode(endpoint.Response); err != nil {
							log.Fatal(err)
						}
					}
				}))
				defer ts.Close()

				client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

				err := endpoint.Call(client, id.ID)

				assert.Nil(t, err)
				// The ID stays inside a single, opaque segment on the wire...
				assert.Equal(t, endpoint.Prefix+id.Encoded, requestTarget)
				// ...and decodes back to exactly the ID the caller asked for.
				assert.Equal(t, endpoint.Prefix+id.ID, decodedPath)
			})
		}
	}
}

// TestPathParameterDotSegmentIsRejected asserts that an ID of exactly "." or
// ".." is refused before any request is made.
//
// Percent-encoding cannot help here: both stay literal dots once escaped, so a
// URL normalizer between the SDK and the API would collapse the segment and
// address a different endpoint than the caller asked for. No resource can be
// identified by such a value in the first place.
func TestPathParameterDotSegmentIsRejected(t *testing.T) {
	ids := []struct {
		Name string
		ID   string
	}{
		{"Dot segment", "."},
		{"Parent dot segment", ".."},
	}

	for _, endpoint := range pathParameterEndpoints() {
		for _, id := range ids {
			t.Run(fmt.Sprintf("%s: %s", endpoint.Name, id.Name), func(t *testing.T) {
				requested := false

				ts := httptest.NewServer(http.HandlerFunc(func(
					w http.ResponseWriter,
					r *http.Request,
				) {
					requested = true

					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{}`))
				}))
				defer ts.Close()

				client := fingerprint.New(fingerprint.WithAPIKey("api_key"), fingerprint.WithBaseURL(ts.URL))

				err := endpoint.Call(client, id.ID)

				// The message stays generic on purpose: it must not tell a caller
				// which values are rejected, since it can surface to end users.
				assert.EqualError(t, err, fmt.Sprintf("invalid value for path parameter %s", endpoint.Param))
				assert.False(t, requested, "no request should have been sent")
			})
		}
	}
}
