package errors

import "fmt"

type Code interface {
	// Code get the error code.
	Code() string
}

type Error[T Code] struct {
	// Code error code.
	Code T `json:"code"`

	// Msg error message.
	Msg string `json:"msg"`
}

// New create a new error.
func New[T Code](code T, msg string, args ...interface{}) *Error[T] {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	return &Error[T]{
		Code: code,
		Msg:  msg,
	}
}

// Error implement the error interface.
func (e *Error[T]) Error() string {
	return e.Code.Code() + ":" + e.Msg
}
