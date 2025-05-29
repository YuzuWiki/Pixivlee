package session

import (
	_errors "errors"
)

// api exception
var (
	ErrApi  = _errors.New("an error occurred in the api")
	ErrAuth = _errors.New("auth error")
)
