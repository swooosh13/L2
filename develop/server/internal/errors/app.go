package errors

import "errors"

var (
	MethodNotAllowed = errors.New("method not allowed")
	JsonDecodingError = errors.New("json decoding error")
	EventNotFound = errors.New("event not found")
)
