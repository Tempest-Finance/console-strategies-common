package goerrors

import (
	"fmt"
)

type Error struct {
	Code          string   `json:"code"`
	Message       string   `json:"message"`
	ErrorEntities []string `json:"errorEntities"`
	RootCause     error    `json:"-"`
}

func NewError(code string, message string, entities []string, rootCause error) *Error {
	return &Error{
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("ERROR: {Code: %s, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func NewErrRequired(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeRequired, ErrMsgRequired, entities, rootCause)
}

func NewErrInvalidFormat(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeInvalidFormat, ErrMsgInvalidFormat, entities, rootCause)
}

func NewErrInvalid(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeInvalid, ErrMsgInvalid, entities, rootCause)
}

func NewErrNotAcceptedValue(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeNotAcceptedValue, ErrMsgNotAcceptedValue, entities, rootCause)
}

func NewErrOutOfRange(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeOutOfRange, ErrMsgOutOfRange, entities, rootCause)
}

func NewErrUnauthenticated(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeUnauthenticated, ErrMsgUnauthenticated, entities, rootCause)
}

func NewErrUnauthorized(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeUnauthorized, ErrMsgUnauthorized, entities, rootCause)
}

func NewErrNotFound(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeNotFound, ErrMsgNotFound, entities, rootCause)
}

func NewErrDuplicate(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeDuplicate, ErrMsgDuplicate, entities, rootCause)
}

func NewErrAlreadyExits(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeAlreadyExists, ErrMsgAlreadyExists, entities, rootCause)
}

func NewErrTooManyRequests(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeTooManyRequests, ErrMsgTooManyRequests, entities, rootCause)
}

func NewErrUnknown(rootCause error, entities ...string) *Error {
	return NewError(ErrCodeUnknown, ErrMsgUnknown, entities, rootCause)
}
