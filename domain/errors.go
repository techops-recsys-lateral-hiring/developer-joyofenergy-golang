package domain

import "errors"

var (
	ErrMissingArgument    = errors.New("missing argument")
	ErrInvalidMessageType = errors.New("invalid message-type")
	ErrNotFound = errors.New("not found")
)
