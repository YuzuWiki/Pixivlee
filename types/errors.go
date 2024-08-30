package types

import "fmt"

type IError interface {
	Code() int
	Message() string

	Error() string
}

type TError struct {
	errcode int
	errmsg  string
}

func (e *TError) Code() int {
	return e.errcode
}

func (e *TError) Message() string {
	return e.errmsg
}

func (e *TError) Error() string {
	return fmt.Sprintf("(%d) %s", e.errcode, e.errmsg)
}

func NewError(code int, msg string) IError {
	return &TError{
		errcode: code,
		errmsg:  msg,
	}
}
