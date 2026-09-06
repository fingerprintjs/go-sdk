package eventhydrate

import (
	"encoding/json"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/v8"
	"github.com/stretchr/testify/assert"
)

func TestEventHydrateMissingSource(t *testing.T) {
	t.Run("omitted source unmarshals as EventDevice", func(t *testing.T) {
		var event fingerprint.Event
		err := json.Unmarshal([]byte(`{"event_id":"1708102555327.NLOjmg","timestamp":1708102555327}`), &event)

		assert.NoError(t, err)
		assert.NotNil(t, event.EventDevice)
		assert.Nil(t, event.EventEdge)
		assert.Equal(t, fingerprint.EventSourceDevice, event.EventDevice.Source)
		assert.Equal(t, "1708102555327.NLOjmg", event.EventDevice.EventID)
	})

	t.Run("source edge stays EventEdge", func(t *testing.T) {
		var event fingerprint.Event
		err := json.Unmarshal([]byte(`{"event_id":"e1","timestamp":1,"source":"edge","ip_info":{}}`), &event)

		assert.NoError(t, err)
		assert.NotNil(t, event.EventEdge)
		assert.Nil(t, event.EventDevice)
		assert.Equal(t, fingerprint.EventSourceEdge, event.EventEdge.Source)
	})

	t.Run("source device stays EventDevice", func(t *testing.T) {
		var event fingerprint.Event
		err := json.Unmarshal([]byte(`{"event_id":"d1","timestamp":1,"source":"device"}`), &event)

		assert.NoError(t, err)
		assert.NotNil(t, event.EventDevice)
		assert.Nil(t, event.EventEdge)
		assert.Equal(t, fingerprint.EventSourceDevice, event.EventDevice.Source)
	})
}
