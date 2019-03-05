package schema

import (
	"errors"
)

// Commone schema validation errors.
var (
	ErrReadOnly      = errors.New("read-only")
	ErrCreateOnly    = errors.New("create-only")
	ErrNotString     = errors.New("not a string")
	ErrInvalidFormat = errors.New("invalid format")

	ErrNotGoTime = errors.New("not a time.Time")
)
