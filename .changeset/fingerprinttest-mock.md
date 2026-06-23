---
'go-sdk': minor
---

**fingerprinttest**: Add `fingerprinttest` package with a pre-built `MockClient` for testing code that depends on the Fingerprint SDK. Import `github.com/fingerprintjs/go-sdk/v8/fingerprinttest` and use `fingerprinttest.NewMockClient(t)` with `Set*Response` convenience methods instead of writing or generating your own mock.
