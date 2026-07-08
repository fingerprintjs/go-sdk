# fingerprinttest

`fingerprinttest` provides a pre-built [testify/mock](https://pkg.go.dev/github.com/stretchr/testify/mock)-based mock for [`fingerprint.ClientInterface`](https://pkg.go.dev/github.com/fingerprintjs/go-sdk/v8#ClientInterface). Use it to test code that depends on the Fingerprint SDK without making real API calls. No mock generation or boilerplate required.

## Installation

```shell
go get "github.com/fingerprintjs/go-sdk/fingerprinttest"
```

## Usage

```go
import (
    "context"
    "errors"
    "testing"

    fingerprint "github.com/fingerprintjs/go-sdk/v8"
    "github.com/fingerprintjs/go-sdk/fingerprinttest"
    "github.com/stretchr/testify/assert"
)

func TestCheckFraud_APIError(t *testing.T) {
    mock := fingerprinttest.NewMockClient(t)
    mock.SetGetEventResponse(nil, nil, errors.New("API unavailable"))

    client := fingerprint.New(fingerprint.WithClientInterface(mock))

    _, err := CheckFraud(context.Background(), client, "evt_123")
    assert.Error(t, err)
}
```

`NewMockClient` registers automatic assertion of expected calls at test teardown, so no manual `AssertExpectations` call is needed.

## API

### Constructor

```go
func NewMockClient(t TestingT) *MockClient
```

Creates a `MockClient` and registers cleanup to assert all expected calls were made when the test ends.

### Response setters

Each method on `ClientInterface` has a corresponding setter that configures the mock's return values. All setters return `*mock.Call` for optional chaining (`.Once()`, `.Times(n)`, `.Maybe()`, etc.).

| Setter | Configures |
|--------|------------|
| `SetGetEventResponse(event, resp, err)` | `GetEvent` |
| `SetSearchEventsResponse(result, resp, err)` | `SearchEvents` |
| `SetUpdateEventResponse(resp, err)` | `UpdateEvent` |
| `SetDeleteVisitorDataResponse(resp, err)` | `DeleteVisitorData` |

### Example: assert a method is called exactly once

```go
mock.SetGetEventResponse(event, nil, nil).Once()
```
