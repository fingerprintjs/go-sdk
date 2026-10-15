package test

import (
	"encoding/json"
	"os"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEventSourceDeserialization(t *testing.T) {
	t.Run("Event without source unmarshals with nil Source", func(t *testing.T) {
		rawJSON, err := os.ReadFile("mocks/events/get_event_200.json")
		require.NoError(t, err)

		var body map[string]interface{}
		require.NoError(t, json.Unmarshal(rawJSON, &body))
		delete(body, "source")
		rawJSON, err = json.Marshal(body)
		require.NoError(t, err)

		var event fingerprint.Event
		require.NoError(t, json.Unmarshal(rawJSON, &event))

		assert.Nil(t, event.Source)
		assert.Equal(t, "Ibk1527CUFmcnjLwIs4A9", event.Identification.VisitorID)
	})

	t.Run("Edge event unmarshals with edge Source", func(t *testing.T) {
		rawJSON, err := os.ReadFile("mocks/events/get_event_with_edge_200.json")
		require.NoError(t, err)

		var event fingerprint.Event
		require.NoError(t, json.Unmarshal(rawJSON, &event))

		require.NotNil(t, event.Source)
		assert.Equal(t, fingerprint.EventSourceEdge, *event.Source)
		assert.Nil(t, event.Identification)
		assert.NotNil(t, event.BotInfo)
		assert.NotNil(t, event.IPInfo)
	})
}
