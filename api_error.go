package fingerprint

import (
	"errors"

	"github.com/fingerprintjs/go-sdk/v8/internal"
)

func AsErrorResponse(err error) (*internal.ErrorResponse, bool) {
	if err == nil {
		return nil, false
	}

	var apiErr *internal.GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return nil, false
	}

	v, ok := apiErr.Model().(internal.ErrorResponse)
	if !ok {
		return nil, false
	}

	return &v, true
}
