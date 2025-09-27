package domain

import "errors"

var (
	ErrConflicted        = errors.New("conflict error")
	ErrDuplicated        = errors.New("duplicated error")
	ErrInvalidFormat     = errors.New("invalid format error")
	ErrInvalidParameters = errors.New("invalid parameters error")
	ErrNotFound          = errors.New("not found error")
	ErrNotPermitted      = errors.New("not permitted error")
	ErrTimeout           = errors.New("timeout error")
	ErrUnknownError      = errors.New("unknown error")
)
