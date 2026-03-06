package fingerprint_test

import (
	"context"
	"net/http"
	"time"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

func ExampleWithAPIKey() {
	_ = fingerprint.New(fingerprint.WithAPIKey("SECRET_API_KEY"))
}

func ExampleWithHTTPClient() {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	_ = fingerprint.New(fingerprint.WithHTTPClient(httpClient))
}

func ExampleWithRegion() {
	// For EU
	_ = fingerprint.New(fingerprint.WithRegion(fingerprint.RegionEU))

	// For Asia
	_ = fingerprint.New(fingerprint.WithRegion(fingerprint.RegionAsia))
}

func ExampleWithDebug() {
	_ = fingerprint.New(fingerprint.WithDebug(true))
}

func ExampleWithBaseURL() {
	_ = fingerprint.New(fingerprint.WithBaseURL("https://example.com"))
}

type FingerprintApiImpl struct{}

func (f FingerprintApiImpl) DeleteVisitorData(visitorID string) openapi.ApiDeleteVisitorDataRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) DeleteVisitorDataExecute(ctx context.Context, r openapi.ApiDeleteVisitorDataRequest) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) GetEvent(eventID string) openapi.ApiGetEventRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) GetEventExecute(ctx context.Context, r openapi.ApiGetEventRequest) (*openapi.Event, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) SearchEvents() openapi.ApiSearchEventsRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) SearchEventsExecute(ctx context.Context, r openapi.ApiSearchEventsRequest) (*openapi.EventSearch, *http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) UpdateEvent(eventID string) openapi.ApiUpdateEventRequest {
	//TODO implement me
	panic("implement me")
}

func (f FingerprintApiImpl) UpdateEventExecute(ctx context.Context, r openapi.ApiUpdateEventRequest) (*http.Response, error) {
	//TODO implement me
	panic("implement me")
}

func ExampleWithFingerprintAPI() {
	_ = fingerprint.New(fingerprint.WithFingerprintAPI(FingerprintApiImpl{}))
}
