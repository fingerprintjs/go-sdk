package functional_test

import (
	"context"
	"os"
	"testing"
	"time"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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
	opts := client.NewSearchEventsRequest(context.Background()).
		Start(start).
		End(end)
	events, _, err := client.SearchEvents(opts)
	assert.NoError(t, err)
	assert.NotNil(t, events.Events)
	testEvent := events.Events[0]
	eventId := testEvent.EventId
	visitorId := testEvent.Identification.VisitorId

	t.Run("GetEvent", func(t *testing.T) {
		t.Run("with valid event", func(t *testing.T) {
			event, _, err := client.GetEvent(context.Background(), eventId)

			assert.NoError(t, err)
			assert.Equal(t, eventId, event.EventId)
			assert.Equal(t, visitorId, event.Identification.VisitorId)
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
			opts := client.NewSearchEventsRequest(context.Background()).
				Start(start).
				End(end).
				Limit(2)
			events, _, err := client.SearchEvents(opts)
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.Len(t, events.Events, 2)
		})

		t.Run("with pagination", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			events, _, err := client.SearchEvents(
				client.NewSearchEventsRequest(context.Background()).
					Limit(2).
					Start(start).
					End(end))
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.NotEmpty(t, events.PaginationKey)
			assert.Len(t, events.Events, 2)

			nextEvents, _, err := client.SearchEvents(client.NewSearchEventsRequest(context.Background()).
				Limit(2).
				Start(start).
				End(end).
				PaginationKey(*events.PaginationKey),
			)
			assert.NoError(t, err)
			assert.NotNil(t, nextEvents)
			assert.Len(t, nextEvents.Events, 2)
			assert.NotEqual(t, events.Events[0].EventId, nextEvents.Events[0].EventId)
			assert.NotEqual(t, events.Events[1].EventId, nextEvents.Events[1].EventId)
		})

		t.Run("with old events", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			events, _, err := client.SearchEvents(client.NewSearchEventsRequest(context.Background()).
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
			oldGetEvent, _, err := client.GetEvent(context.Background(), oldEvent.EventId)
			assert.NoError(t, err)
			assert.NotNil(t, oldGetEvent)
			assert.Equal(t, oldEvent.EventId, oldGetEvent.EventId)
		})
	})
}
