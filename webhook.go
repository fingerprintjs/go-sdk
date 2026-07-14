package fingerprint

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func isValidHmacSignature(signature string, data []byte, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)
	computedSignature := hex.EncodeToString(h.Sum(nil))
	return signature == computedSignature
}

// IsValidWebhookSignature Verifies the HMAC signature extracted from the "fpjs-event-signature" header of the incoming request.
// Fingerprint signs every webhook request with the signing key configured for that webhook; see https://docs.fingerprint.com/docs/webhooks#protecting-your-webhooks to learn more.
func IsValidWebhookSignature(header string, data []byte, secret string) bool {
	signatures := strings.Split(header, ",")

	for _, signature := range signatures {
		parts := strings.Split(signature, "=")
		if len(parts) == 2 && parts[0] == "v1" {
			hash := parts[1]
			if isValidHmacSignature(hash, data, secret) {
				return true
			}
		}
	}

	return false
}
