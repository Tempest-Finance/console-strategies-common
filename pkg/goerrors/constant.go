package goerrors

const (
	ErrCodeRequired = "REQUIRED"
	ErrMsgRequired  = "Missing required fields"

	ErrCodeNotAcceptedValue = "NOT_ACCEPTED_VALUE"
	ErrMsgNotAcceptedValue  = "Data is not in the accepted values"

	ErrCodeOutOfRange = "OUT_OF_RANGE"
	ErrMsgOutOfRange  = "Data is out of range"

	ErrCodeInvalidFormat = "INVALID_FORMAT"
	ErrMsgInvalidFormat  = "Data has an invalid format"

	ErrCodeInvalid = "INVALID"
	ErrMsgInvalid  = "Data is invalid"

	ErrCodeUnauthenticated = "UNAUTHENTICATED"
	ErrMsgUnauthenticated  = "Unauthenticated"

	ErrCodeUnauthorized = "UNAUTHORIZED"
	ErrMsgUnauthorized  = "Unauthorized"

	ErrCodeNotFound = "NOT_FOUND"
	ErrMsgNotFound  = "Not found"

	ErrCodeDuplicate = "DUPLICATE"
	ErrMsgDuplicate  = "Duplicate data"

	ErrCodeAlreadyExists = "ALREADY_EXISTS"
	ErrMsgAlreadyExists  = "Data already exists"

	ErrCodeTooManyRequests = "TOO_MANY_REQUESTS"
	ErrMsgTooManyRequests  = "Too many requests"

	ErrCodeUnknown = "UNKNOWN"
	ErrMsgUnknown  = "Unknown"
)

const (
	//HTTP 200 - OK
	ClientErrCodeOK = 0
	ClientErrMsgOK  = "Successfully"

	//HTTP 400 - Bad Request
	ClientErrCodeRequired = 4000
	ClientErrMsgRequired  = "Missing required fields"

	ClientErrCodeNotAcceptedValue = 4001
	ClientErrMsgNotAcceptedValue  = "Input is not in the accepted values"

	ClientErrCodeOutOfRange = 4002
	ClientErrMsgOutOfRange  = "Input is out of range"

	ClientErrCodeInvalidFormat = 4003
	ClientErrMsgInvalidFormat  = "Input has an invalid format"

	ClientErrCodeInvalid = 4004
	ClientErrMsgInvalid  = "Input is invalid"

	//HTTP 401 - Unauthenticated
	ClientErrCodeUnauthenticated = 4010
	ClientErrMsgUnauthenticated  = "Unauthenticated"

	//HTTP 403 - Unauthorized
	ClientErrCodeUnauthorized = 4030
	ClientErrMsgUnauthorized  = "Unauthorized"

	//HTTP 404 - Not found
	ClientErrCodeNotFound = 4040
	ClientErrMsgNotFound  = "Not found"

	//HTTP 409 - Duplicate
	ClientErrCodeDuplicate = 4090
	ClientErrMsgDuplicate  = "Duplicate data"

	ClientErrCodeAlreadyExists = 4091
	ClientErrMsgAlreadyExists  = "Data already exists"

	//HTTP 429 - Too Many Requests
	ClientErrCodeTooManyRequests = 4290
	ClientErrMsgTooManyRequests  = "Too many requests"

	//HTTP 500 - Internal Server Error
	ClientErrCodeInternal = 5000
	ClientErrMsgInternal  = "Internal Server Error"
)
