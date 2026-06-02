package openapi

import "fmt"

func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
