package connection

import "errors"

var ErrConnectionNotFound = errors.New("connection not found")

// ErrInvalidProvider is returned when an unsupported provider is specified.
var ErrInvalidProvider = errors.New("unsupported provider")

// ErrInvalidEndpoint is returned when an invalid endpoint URL is specified.
var ErrInvalidEndpoint = errors.New("invalid endpoint URL")
