package core_errors

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrConflict        = errors.New("conflict")
)

var (
	ErrRedisUnavailable       = errors.New("redis unavailable")
	ErrRedisConnectionRefused = errors.New("connection refused")
)
