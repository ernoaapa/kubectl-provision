package cmd

import (
	"fmt"

	"github.com/pkg/errors"
)

// Definitions of common error types used throughout runtime implementation.
// All errors returned by the interface will map into one of these errors classes.
var (
	ErrNotFound     = errors.New("not found")
	ErrTooManyFound = errors.New("too many found")
)

// IsNotFound returns true if the error is due to a missing resource
func IsNotFound(err error) bool {
	return errors.Cause(err) == ErrNotFound
}

// IsTooManyFound returns true if the error is due to too many resources found
func IsTooManyFound(err error) bool {
	return errors.Cause(err) == ErrTooManyFound
}
func ErrWithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessage(err, fmt.Sprintf(format, args...))
}
