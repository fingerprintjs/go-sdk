package test

import (
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk/sdk"
	"github.com/stretchr/testify/assert"
)

func TestApiFingerprint(t *testing.T) {
	t.Run("Create with empty config", func(t *testing.T) {
		client := fingerprint.New()

		assert.NotNil(t, client)
	})
}
