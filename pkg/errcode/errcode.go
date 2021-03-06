package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]struct{}{}

// NewError warp error
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("The error code %d exists", code))
	}
	codes[code] = struct{}{}
	return &Error{code: code, msg: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// StatusCode trans err code to http status code
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ErrInternalServer.Code():
		return http.StatusInternalServerError
	case ErrInvalidParam.Code():
		return http.StatusBadRequest
	case ErrToken.Code():
		fallthrough
	case ErrInvalidToken.Code():
		fallthrough
	case ErrTokenTimeout.Code():
		return http.StatusUnauthorized
	case ErrTooManyRequests.Code():
		return http.StatusTooManyRequests
	case ErrServiceUnavailable.Code():
		return http.StatusServiceUnavailable
	}

	return http.StatusInternalServerError
}

// Err represents an error
type Err struct {
	Code int
	Msg  string
	Err  error
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Msg, err.Err)
}

// DecodeErr decode the error and return the error code and error message
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.code, Success.msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Error:
		return typed.code, typed.msg
	default:
	}

	return ErrInternalServer.Code(), err.Error()
}
