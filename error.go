package version

import "fmt"

// Version error.
var (
	ErrInvalidVersion = Error{Message: "invalid version %q"}
)

// Error represents a version error.
type Error struct {
	Message string
	Version string
}

// Error retrieves the message of a REST API error.
func (e Error) Error() string {
	return fmt.Sprintf(e.Message, e.Version)
}

// Is determines whether the Error is an instance of the target.
// https://pkg.go.dev/errors#Is
//
// This implementation does not compare versions.
func (e Error) Is(target error) bool {
	if t, ok := target.(Error); ok {
		return t.Message == e.Message
	}
	return false
}

func invalid(version string) Error {
	return Error{
		Message: ErrInvalidVersion.Message,
		Version: version,
	}
}
