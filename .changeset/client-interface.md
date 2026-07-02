---
'go-sdk': minor
---

**client-interface**: Add `ClientInterface` and `WithClientInterface` option, allowing custom implementations of the Fingerprint client without needing to implement the low-level API internals required by `WithFingerprintAPI`. `WithFingerprintAPI` is now deprecated in favor of `WithClientInterface`.
