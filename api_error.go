package fingerprint

import (
	"errors"

	"github.com/fingerprintjs/go-sdk/v8/internal/openapi"
)

func AsErrorResponse(err error) (*openapi.ErrorResponse, bool) {
	if err == nil {
		return nil, false
	}

	var apiErr *openapi.GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return nil, false
	}

	v, ok := apiErr.Model().(openapi.ErrorResponse)
	if !ok {
		return nil, false
	}

	return &v, true
}
