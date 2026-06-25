package fingerprinttest

import (
	"context"
	"net/http"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/mock"
)

// MockClient is a testify/mock-based implementation of fingerprint.ClientInterface.
// Use NewMockClient to create an instance. It registers automatic assertion of
// expected calls at test teardown.
type MockClient struct {
	mock.Mock
}

// NewMockClient creates a MockClient and registers cleanup to assert all expected
// calls were made when the test ends.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	m := &MockClient{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

// compile-time check that MockClient satisfies fingerprint.ClientInterface
var _ fingerprint.ClientInterface = (*MockClient)(nil)

// GetEvent implements fingerprint.ClientInterface.
func (m *MockClient) GetEvent(ctx context.Context, eventID string, opts ...fingerprint.GetEventOption) (*fingerprint.Event, *http.Response, error) {
	args := m.Called(ctx, eventID, opts)
	event := args.Get(0).(*fingerprint.Event)
	resp := args.Get(1).(*http.Response)
	return event, resp, args.Error(2)
}

// SearchEvents implements fingerprint.ClientInterface.
func (m *MockClient) SearchEvents(ctx context.Context, req fingerprint.SearchEventRequest) (*fingerprint.EventSearch, *http.Response, error) {
	args := m.Called(ctx, req)
	result := args.Get(0).(*fingerprint.EventSearch)
	resp := args.Get(1).(*http.Response)
	return result, resp, args.Error(2)
}

// UpdateEvent implements fingerprint.ClientInterface.
func (m *MockClient) UpdateEvent(ctx context.Context, eventId string, eventUpdateReq fingerprint.EventUpdate) (*http.Response, error) {
	args := m.Called(ctx, eventId, eventUpdateReq)
	resp := args.Get(0).(*http.Response)
	return resp, args.Error(1)
}

// DeleteVisitorData implements fingerprint.ClientInterface.
func (m *MockClient) DeleteVisitorData(ctx context.Context, visitorId string) (*http.Response, error) {
	args := m.Called(ctx, visitorId)
	resp := args.Get(0).(*http.Response)
	return resp, args.Error(1)
}

// SetGetEventResponse configures the mock to return the given values for any GetEvent call.
// Returns the underlying *mock.Call for optional chaining (.Once(), .Times(n), etc.).
func (m *MockClient) SetGetEventResponse(event *fingerprint.Event, resp *http.Response, err error) *mock.Call {
	return m.On("GetEvent", mock.Anything, mock.Anything, mock.Anything).Return(event, resp, err)
}

// SetSearchEventsResponse configures the mock to return the given values for any SearchEvents call.
// Returns the underlying *mock.Call for optional chaining (.Once(), .Times(n), etc.).
func (m *MockClient) SetSearchEventsResponse(result *fingerprint.EventSearch, resp *http.Response, err error) *mock.Call {
	return m.On("SearchEvents", mock.Anything, mock.Anything).Return(result, resp, err)
}

// SetUpdateEventResponse configures the mock to return the given values for any UpdateEvent call.
// Returns the underlying *mock.Call for optional chaining (.Once(), .Times(n), etc.).
func (m *MockClient) SetUpdateEventResponse(resp *http.Response, err error) *mock.Call {
	return m.On("UpdateEvent", mock.Anything, mock.Anything, mock.Anything).Return(resp, err)
}

// SetDeleteVisitorDataResponse configures the mock to return the given values for any DeleteVisitorData call.
// Returns the underlying *mock.Call for optional chaining (.Once(), .Times(n), etc.).
func (m *MockClient) SetDeleteVisitorDataResponse(resp *http.Response, err error) *mock.Call {
	return m.On("DeleteVisitorData", mock.Anything, mock.Anything).Return(resp, err)
}
