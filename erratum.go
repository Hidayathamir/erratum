package erratum

import (
	"fmt"
	"strings"
)

// Wrap return wrapped inner error with outer error, the format is
// outer: inner.
func Wrap(inner error, outer error) error {
	return fmt.Errorf("%w: %w", outer, inner)
}

// AddCause return wrapped outer error with inner error, the format is
// err: cause.
func AddCause(err error, cause error) error {
	return fmt.Errorf("%w: %w", err, cause)
}

// Cause return string from the deepest error.
func Cause(err error) string {
	errTrace := strings.Split(err.Error(), ":")
	return strings.TrimSpace(errTrace[len(errTrace)-1])
}
