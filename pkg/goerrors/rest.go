package goerrors

import (
	"fmt"
	"net/http"
	"strings"
)

type RestAPIError struct {
	HttpStatus    int           `json:"-"`
	Code          int           `json:"code"`
	Message       string        `json:"message"`
	ErrorEntities []string      `json:"errorEntities"`
	Details       []interface{} `json:"details"`
	RootCause     error         `json:"-"`
}

func NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error) *RestAPIError {
	return &RestAPIError{
		HttpStatus:    httpStatus,
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *RestAPIError) Error() string {
	return fmt.Sprintf("API ERROR: {Code: %d, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func AppendEntitiesToErrMsg(message string, entities []string) string {
	if len(entities) > 0 {
		message += ": "
		message += strings.Join(entities, ",")
	}
	return message
}

func NewRestAPIErrRequired(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgRequired, entities)
	return NewRestAPIError(http.StatusBadRequest, ClientErrCodeRequired, message, entities, rootCause)
}

func NewRestAPIErrInvalidFormat(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgInvalidFormat, entities)
	return NewRestAPIError(http.StatusBadRequest, ClientErrCodeInvalidFormat, message, entities, rootCause)
}

func NewRestAPIErrInvalid(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgInvalid, entities)
	return NewRestAPIError(http.StatusBadRequest, ClientErrCodeInvalid, message, entities, rootCause)
}

func NewRestAPIErrNotAcceptedValue(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgNotAcceptedValue, entities)
	return NewRestAPIError(http.StatusBadRequest, ClientErrCodeNotAcceptedValue, message, entities, rootCause)
}

func NewRestAPIErrOutOfRange(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgOutOfRange, entities)
	return NewRestAPIError(http.StatusBadRequest, ClientErrCodeOutOfRange, message, entities, rootCause)
}

func NewRestAPIErrUnauthenticated(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgUnauthenticated, entities)
	return NewRestAPIError(http.StatusUnauthorized, ClientErrCodeUnauthenticated, message, entities, rootCause)
}

func NewRestAPIErrUnauthorized(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgUnauthorized, entities)
	return NewRestAPIError(http.StatusForbidden, ClientErrCodeUnauthorized, message, entities, rootCause)
}

func NewRestAPIErrNotFound(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgNotFound, entities)
	return NewRestAPIError(http.StatusNotFound, ClientErrCodeNotFound, message, entities, rootCause)
}

func NewRestAPIErrDuplicate(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgDuplicate, entities)
	return NewRestAPIError(http.StatusConflict, ClientErrCodeDuplicate, message, entities, rootCause)
}

func NewRestAPIErrAlreadyExits(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgAlreadyExists, entities)
	return NewRestAPIError(http.StatusConflict, ClientErrCodeAlreadyExists, message, entities, rootCause)
}

func NewRestAPIErrTooManyRequests(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(ClientErrMsgTooManyRequests, entities)
	return NewRestAPIError(http.StatusTooManyRequests, ClientErrCodeTooManyRequests, message, entities, rootCause)
}

func NewRestAPIErrInternal(rootCause error, entities ...string) *RestAPIError {
	return NewRestAPIError(http.StatusInternalServerError, ClientErrCodeInternal, ClientErrMsgInternal, entities, rootCause)
}
