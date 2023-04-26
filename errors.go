package dea

import "errors"

var (
	ErrDisposableEmail          = errors.New("email is a disposable email address")
	ErrInvalidEmail             = errors.New("not valid email format")
	ErrCouldNotLookupSPFRecords = errors.New("could not lookup SPF records")
)
