package Pixivlee

import (
	_errors "errors"
)

// kernel exception
var (
	ErrKernel = _errors.New("kernel error")

	ErrAuth = _errors.New("auth error")
)

// api exception
var (
	ErrApi = _errors.New("an error occurred in the api")
)
