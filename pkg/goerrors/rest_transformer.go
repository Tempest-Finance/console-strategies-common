package goerrors

import (
	"encoding/json"
	errs "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type IRestTransformer interface {
	ErrToRestAPIErr(err *Error) *RestAPIError
	ValidationErrToRestAPIErr(err error) *RestAPIError
	RegisterTransformFunc(errCode string, transformFunc restTransformFunc)
	RegisterValidationTag(tag string, function restTransformFunc)
}

type restTransformFunc func(rootCause error, entities ...string) *RestAPIError

type restTransformer struct {
	mapping       map[string]restTransformFunc
	validationErr map[string]restTransformFunc
}

var restTransformerInstance *restTransformer

func initRestTransformerInstance() {
	if restTransformerInstance == nil {
		restTransformerInstance = &restTransformer{
			mapping:       make(map[string]restTransformFunc),
			validationErr: make(map[string]restTransformFunc),
		}

		restTransformerInstance.RegisterTransformFunc(ErrCodeRequired, NewRestAPIErrRequired)
		restTransformerInstance.RegisterTransformFunc(ErrCodeNotAcceptedValue, NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterTransformFunc(ErrCodeOutOfRange, NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterTransformFunc(ErrCodeInvalidFormat, NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterTransformFunc(ErrCodeInvalid, NewRestAPIErrInvalid)
		restTransformerInstance.RegisterTransformFunc(ErrCodeUnauthenticated, NewRestAPIErrUnauthenticated)
		restTransformerInstance.RegisterTransformFunc(ErrCodeUnauthorized, NewRestAPIErrUnauthorized)
		restTransformerInstance.RegisterTransformFunc(ErrCodeNotFound, NewRestAPIErrNotFound)
		restTransformerInstance.RegisterTransformFunc(ErrCodeDuplicate, NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterTransformFunc(ErrCodeAlreadyExists, NewRestAPIErrAlreadyExits)
		restTransformerInstance.RegisterTransformFunc(ErrCodeTooManyRequests, NewRestAPIErrTooManyRequests)
		restTransformerInstance.RegisterTransformFunc(ErrCodeUnknown, NewRestAPIErrInternal)

		restTransformerInstance.RegisterValidationTag("required", NewRestAPIErrRequired)
		restTransformerInstance.RegisterValidationTag("oneof", NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterValidationTag("min", NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("max", NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("numeric", NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("unique", NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterValidationTag("hexadecimal", NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("email", NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("url", NewRestAPIErrInvalidFormat)
	}
}

func RestTransformerInstance() IRestTransformer {
	return restTransformerInstance
}

// ValidationErrToRestAPIErr transforms ValidationError to RestAPIError
// this function will be used when bind JSON request to DTO in gin framework
func (t *restTransformer) ValidationErrToRestAPIErr(err error) *RestAPIError {
	var sliceValidationErr binding.SliceValidationError
	var validationErrs validator.ValidationErrors
	var unmarshalTypeErr *json.UnmarshalTypeError
	var jsonSynTaxErr *json.SyntaxError
	var numErr *strconv.NumError
	if errs.As(err, &sliceValidationErr) {
		if len(sliceValidationErr) > 0 {
			err := sliceValidationErr[0]
			if errs.As(err, &validationErrs) {
				if len(validationErrs) > 0 {
					validationErr := validationErrs[0]
					return t.apiErrForTag(validationErr.Tag(), err, ToLowerFirstLetter(validationErr.Field()))
				}
			}
		}
	}
	if errs.As(err, &validationErrs) {
		if len(validationErrs) > 0 {
			validationErr := validationErrs[0]
			return t.apiErrForTag(validationErr.Tag(), err, ToLowerFirstLetter(validationErr.Field()))
		}
	}
	if errs.As(err, &unmarshalTypeErr) {
		field := unmarshalTypeErr.Field
		fieldArr := strings.Split(field, ".")
		return NewRestAPIErrInvalidFormat(err, ToLowerFirstLetter(fieldArr[len(fieldArr)-1]))
	}
	if errs.As(err, &jsonSynTaxErr) {
		return NewRestAPIErrInvalidFormat(err)
	}
	if errs.As(err, &numErr) {
		return NewRestAPIErrInvalidFormat(err)
	}
	return NewRestAPIErrInternal(err)
}

// ErrToRestAPIErr transforms Error to RestAPIError
func (t *restTransformer) ErrToRestAPIErr(err *Error) *RestAPIError {
	f := t.mapping[err.Code]
	if f == nil {
		return NewRestAPIErrInternal(fmt.Errorf("can not transform error, error: %v", err))
	}
	return f(err, err.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform DomainError to RestAPIError
// if the domainErrCode is already registered, the old transform function will be overridden
func (t *restTransformer) RegisterTransformFunc(domainErrCode string, function restTransformFunc) {
	t.mapping[domainErrCode] = function
}

// RegisterValidationTag is used to define new validation tag and respective API error
// if the validation tag is already registered, the old respective API error will be overridden
func (t *restTransformer) RegisterValidationTag(tag string, function restTransformFunc) {
	t.validationErr[tag] = function
}

// apiErrForTag return RestAPIError which corresponds to the validation tag
func (t *restTransformer) apiErrForTag(tag string, err error, fields ...string) *RestAPIError {
	f := t.validationErr[tag]
	if f == nil {
		return NewRestAPIErrInternal(err)
	}
	return f(err, fields...)
}
