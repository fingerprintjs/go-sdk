package functional_test

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestApiFunctional(t *testing.T) {
	// Load environment variables
	godotenv.Load()

	apiKey := os.Getenv("FINGERPRINT_API_KEY")
	require.NotEmpty(t, apiKey)

	client := fingerprint.New(fingerprint.WithAPIKey(apiKey))

	end := time.Now().UnixMilli()
	start := time.Now().AddDate(0, 0, -90).UnixMilli()
	opts := fingerprint.NewSearchEventsRequest().
		Start(start).
		End(end)
	events, _, err := client.SearchEvents(context.Background(), opts)
	assert.NoError(t, err)
	assert.NotNil(t, events.Events)
	testEvent := events.Events[0]
	eventId := testEvent.EventID
	visitorId := testEvent.Identification.VisitorID

	t.Run("GetEvent", func(t *testing.T) {
		t.Run("with valid event", func(t *testing.T) {
			event, _, err := client.GetEvent(context.Background(), eventId)

			assert.NoError(t, err)
			assert.Equal(t, eventId, event.EventID)
			assert.Equal(t, visitorId, event.Identification.VisitorID)
		})

		t.Run("error 404", func(t *testing.T) {
			event, _, err := client.GetEvent(context.Background(), "1662542583652.pLBzes")

			assert.Error(t, err)
			assert.Nil(t, event)

			errorResponse, ok := fingerprint.AsErrorResponse(err)
			require.True(t, ok)

			assert.Equal(t, fingerprint.ErrorCodeEvent_not_found, errorResponse.Error.Code)
		})
	})

	t.Run("SearchEvents", func(t *testing.T) {
		t.Run("simple search", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			opts := fingerprint.NewSearchEventsRequest().
				Start(start).
				End(end).
				Limit(2)
			events, _, err := client.SearchEvents(context.Background(), opts)
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.Len(t, events.Events, 2)
		})

		t.Run("with pagination", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			events, _, err := client.SearchEvents(
				context.Background(),
				fingerprint.NewSearchEventsRequest().
					Limit(2).
					Start(start).
					End(end))
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.NotEmpty(t, events.PaginationKey)
			assert.Len(t, events.Events, 2)

			nextEvents, _, err := client.SearchEvents(context.Background(), fingerprint.NewSearchEventsRequest().
				Limit(2).
				Start(start).
				End(end).
				PaginationKey(*events.PaginationKey),
			)
			assert.NoError(t, err)
			assert.NotNil(t, nextEvents)
			assert.Len(t, nextEvents.Events, 2)
			assert.NotEqual(t, events.Events[0].EventID, nextEvents.Events[0].EventID)
			assert.NotEqual(t, events.Events[1].EventID, nextEvents.Events[1].EventID)
		})

		t.Run("with old events", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			events, _, err := client.SearchEvents(context.Background(), fingerprint.NewSearchEventsRequest().
				Start(start).
				End(end).
				Limit(2).
				Reverse(true),
			)

			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.Len(t, events.Events, 2)

			oldEvent := events.Events[0]

			// Try to get old events to check if they still could be deserialized
			oldGetEvent, _, err := client.GetEvent(context.Background(), oldEvent.EventID)
			assert.NoError(t, err)
			assert.NotNil(t, oldGetEvent)
			assert.Equal(t, oldEvent.EventID, oldGetEvent.EventID)
		})
	})
}

// Verify that Fingerprint API interactions can be mocked by SDK users in their own unit tests.
func TestMockFingerprintApi(t *testing.T) {
	mockAPI := MockFingerprintAPI{}

	client := fingerprint.New(fingerprint.WithAPIKey("test-key"), fingerprint.WithFingerprintAPI(&mockAPI))

	mockResp := fingerprint.Event{
		EventID: "test-event-id",
	}
	mockAPI.On("GetEvent", mock.Anything).Return()
	mockAPI.On("GetEventExecute", mock.Anything, mock.Anything).Return(&mockResp, &http.Response{StatusCode: 200}, nil)

	event, httpResp, err := client.GetEvent(context.Background(), "test-event-id")
	require.Nil(t, err)
	require.Equal(t, &mockResp, event)
	require.Equal(t, 200, httpResp.StatusCode)
	mockAPI.AssertCalled(t, "GetEventExecute", mock.Anything, mock.Anything)
}

type MockFingerprintAPI struct {
	mock.Mock
}

func (m *MockFingerprintAPI) DeleteVisitorData(visitorId string) fingerprint.APIDeleteVisitorDataRequest {
	args := m.Called(visitorId)
	if req, ok := args.Get(0).(fingerprint.APIDeleteVisitorDataRequest); ok {
		return req
	}
	return fingerprint.APIDeleteVisitorDataRequest{
		ApiService: m,
	}
}

func (m *MockFingerprintAPI) DeleteVisitorDataExecute(ctx context.Context, r fingerprint.APIDeleteVisitorDataRequest) (*http.Response, error) {
	args := m.Called(ctx, r)
	resp, _ := args.Get(0).(*http.Response)
	err, _ := args.Get(1).(error)
	return resp, err
}

func (m *MockFingerprintAPI) GetEvent(eventId string) fingerprint.APIGetEventRequest {
	args := m.Called(eventId)
	if len(args) > 0 {
		if req, ok := args.Get(0).(fingerprint.APIGetEventRequest); ok {
			return req
		}
	}

	return fingerprint.APIGetEventRequest{
		ApiService: m,
	}
}

func (m *MockFingerprintAPI) GetEventExecute(ctx context.Context, r fingerprint.APIGetEventRequest) (*fingerprint.Event, *http.Response, error) {
	args := m.Called(ctx, r)
	event, _ := args.Get(0).(*fingerprint.Event)
	httpResp, _ := args.Get(1).(*http.Response)
	err, _ := args.Get(2).(error)
	return event, httpResp, err
}

func (m *MockFingerprintAPI) SearchEvents() fingerprint.APISearchEventsRequest {
	args := m.Called()
	if req, ok := args.Get(0).(fingerprint.APISearchEventsRequest); ok {
		return req
	}

	return fingerprint.APISearchEventsRequest{
		ApiService: m,
	}
}

func (m *MockFingerprintAPI) SearchEventsExecute(ctx context.Context, r fingerprint.APISearchEventsRequest) (*fingerprint.EventSearch, *http.Response, error) {
	args := m.Called(ctx, r)
	events, _ := args.Get(0).(*fingerprint.EventSearch)
	httpResp, _ := args.Get(1).(*http.Response)
	err, _ := args.Get(2).(error)
	return events, httpResp, err
}

func (m *MockFingerprintAPI) UpdateEvent(eventId string) fingerprint.APIUpdateEventRequest {
	args := m.Called(eventId)
	if len(args) > 0 {
		if req, ok := args.Get(0).(fingerprint.APIUpdateEventRequest); ok {
			return req
		}
	}

	return fingerprint.APIUpdateEventRequest{
		ApiService: m,
	}
}

func (m *MockFingerprintAPI) UpdateEventExecute(ctx context.Context, r fingerprint.APIUpdateEventRequest) (*http.Response, error) {
	args := m.Called(ctx, r)
	httpResp, _ := args.Get(0).(*http.Response)
	err, _ := args.Get(1).(error)
	return httpResp, err
}
