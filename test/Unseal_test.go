package test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	fingerprint "github.com/fingerprintjs/go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func base64Decode(input string) []byte {
	output, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return output
}

func readSealedResultFromFile(path string) []byte {
	encoded, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return base64Decode(string(encoded))
}

func TestUnsealEventsResponse(t *testing.T) {
	key := base64Decode("p2PA7MGy5tx56cnyJaFZMr96BCFwZeHjZV2EqMvTq53=")
	invalidKey := base64Decode("p2PA7MGy5tx56cnyJaFZMr96BCFwZeHjZV2EqMvTq54=")  // 53 => 54
	invalidKey2 := base64Decode("p2PA7MGy5tx56cnyJaFZMr96BCFwZeHjZV2EqMzTq53=") // vTq => zTq

	sealedResult := readSealedResultFromFile("mocks/sealed/get_event_200.txt")

	strResponse, err := os.ReadFile("mocks/get_event_200.json")
	require.Nil(t, err)

	var expectedResponse fingerprint.Event
	err = json.Unmarshal([]byte(strResponse), &expectedResponse)
	require.Nil(t, err)

	t.Run("with correct result and key", func(t *testing.T) {
		keys := []fingerprint.DecryptionKey{
			{
				Key:       invalidKey,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		unsealed, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.Nil(t, err)

		isEqual := reflect.DeepEqual(*unsealed, expectedResponse)

		assert.True(t, isEqual)
	})

	t.Run("with empty result", func(t *testing.T) {
		_, err := fingerprint.UnsealEventsResponse([]byte(""), []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, "invalid sealed data header", err.Error())
	})

	t.Run("with empty key", func(t *testing.T) {
		_, err := fingerprint.UnsealEventsResponse(sealedResult, []fingerprint.DecryptionKey{
			{
				Key:       base64Decode(""),
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, "unseal failed: new cipher\ncrypto/aes: invalid key size 0.", err.Error())
	})

	t.Run("with invalid sealed result", func(t *testing.T) {
		invalidSealedResult := base64Decode("noXc7VOpBstjjcavDKSKr4HTavt4mdq8h6NC32T0hUtw9S0jXT8lPjZiWL8SyHxmrF3uTGqO+g==")

		keys := []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(invalidSealedResult, keys)

		assert.Error(t, err, fingerprint.ErrInvalidEventResponse)
	})

	t.Run("with sealed result as broken json", func(t *testing.T) {
		invalidJsonSealedResult := readSealedResultFromFile("mocks/sealed/invalid_json.txt")

		keys := []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err = fingerprint.UnsealEventsResponse(invalidJsonSealedResult, keys)
		require.NotNil(t, err)

		assert.Equal(t, "unexpected end of JSON input", err.Error())
	})

	t.Run("with not compressed result", func(t *testing.T) {
		sealedResult := readSealedResultFromFile("mocks/sealed/invalid_compression.txt")

		keys := []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.Error(t, err)
		assert.IsType(t, err, &fingerprint.AggregatedUnsealError{})

		var aggregateError *fingerprint.AggregatedUnsealError
		errors.As(err, &aggregateError)
		expectedErrorMessage := "unseal failed: decompress\nflate: corrupt input before offset 13\ninflated payload read all bytes."
		assert.Equal(t, expectedErrorMessage, aggregateError.Error())
	})

	t.Run("with invalid keys", func(t *testing.T) {
		keys := []fingerprint.DecryptionKey{
			{
				Key:       invalidKey,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
			{
				Key:       base64Decode("aW52YWxpZF9rZXk="),
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
			{
				Key:       []byte("invalid key"),
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
			{
				Key:       invalidKey2,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		require.NotNil(t, err)
		assert.IsType(t, &fingerprint.AggregatedUnsealError{}, err)
		expectedErrorMessage := "unseal failed: aesgcm open\ncipher: message authentication failed, new cipher\ncrypto/aes: invalid key size 11, new cipher\ncrypto/aes: invalid key size 11, aesgcm open\ncipher: message authentication failed."
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("with invalid algorithm", func(t *testing.T) {
		keys := []fingerprint.DecryptionKey{
			{
				// Invalid key
				Key:       invalidKey,
				Algorithm: "INVALID",
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.ErrorIs(t, err, fingerprint.ErrInvalidAlgorithm)
	})

	t.Run("with invalid header", func(t *testing.T) {
		sealedResult := readSealedResultFromFile("mocks/sealed/invalid_header.txt")

		keys := []fingerprint.DecryptionKey{
			{
				Key:       invalidKey,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.ErrorIs(t, err, fingerprint.ErrInvalidHeader)
	})

	t.Run("with empty data", func(t *testing.T) {
		sealedResult := []byte("")

		keys := []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.Error(t, err, fingerprint.ErrInvalidHeader)
	})

	t.Run("with bad nonce", func(t *testing.T) {
		sealedResult := []byte{0x9E, 0x85, 0xDC, 0xED, 0xAA, 0xBB, 0xCC}

		keys := []fingerprint.DecryptionKey{
			{
				Key:       key,
				Algorithm: fingerprint.AlgorithmAES256GCM,
			},
		}
		_, err := fingerprint.UnsealEventsResponse(sealedResult, keys)

		assert.Error(t, err)
		assert.IsType(t, err, &fingerprint.AggregatedUnsealError{})

		var aggregateError *fingerprint.AggregatedUnsealError
		errors.As(err, &aggregateError)

		expectedErrorMessage := "unseal failed: nonce."
		assert.Equal(t, expectedErrorMessage, aggregateError.Error())
	})
}
