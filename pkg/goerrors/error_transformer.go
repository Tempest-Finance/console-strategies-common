package goerrors

import (
	"fmt"
)

type IErrorTransformer interface {
	RestAPIErrToErr(restAPIErr *RestAPIError) *Error
}

type errTransformFunc func(rootCause error, entities ...string) *Error

type errTransformer struct {
	mapping map[int]errTransformFunc
}

var errTransformerInstance *errTransformer

func initErrTransformerInstance() {
	if errTransformerInstance == nil {
		errTransformerInstance = &errTransformer{
			mapping: make(map[int]errTransformFunc),
		}

		errTransformerInstance.RegisterTransformFunc(ClientErrCodeRequired, NewErrRequired)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeNotAcceptedValue, NewErrNotAcceptedValue)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeOutOfRange, NewErrOutOfRange)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeInvalidFormat, NewErrInvalidFormat)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeInvalid, NewErrInvalid)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeUnauthenticated, NewErrUnauthenticated)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeUnauthorized, NewErrUnauthorized)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeNotFound, NewErrNotFound)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeDuplicate, NewErrDuplicate)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeAlreadyExists, NewErrAlreadyExits)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeTooManyRequests, NewErrTooManyRequests)
		errTransformerInstance.RegisterTransformFunc(ClientErrCodeInternal, NewErrUnknown)
	}
}

func ErrTransformerInstance() IErrorTransformer {
	return errTransformerInstance
}

// RestAPIErrToErr transforms RestAPIError to Error
func (t *errTransformer) RestAPIErrToErr(restAPIErr *RestAPIError) *Error {
	f := t.mapping[restAPIErr.Code]
	if f == nil {
		return NewErrUnknown(fmt.Errorf("can not transform rest API error, error: %v", restAPIErr))
	}
	return f(restAPIErr, restAPIErr.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform RestAPIError to Error
// if the restAPIErrorCode is already registered, the old transform function will be overridden
func (t *errTransformer) RegisterTransformFunc(restAPIErrCode int, transformFunc errTransformFunc) {
	t.mapping[restAPIErrCode] = transformFunc
}
