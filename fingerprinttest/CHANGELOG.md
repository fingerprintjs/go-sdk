# fingerprinttest

## 1.0.1

### Patch Changes

- Fix the `go.mod` module path. The package now installs as `github.com/fingerprintjs/go-sdk/fingerprinttest`. `v1.0.0` was never installable and should not be used.

## 1.0.0

### Major Changes

- Initial release of `fingerprinttest`, a testing package that provides `MockClient`, an implementation of `fingerprint.ClientInterface` for mocking API calls in unit tests
