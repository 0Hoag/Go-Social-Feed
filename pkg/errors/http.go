package errors

import "net/http"

type HTTPError struct {
	Code       int
	Message    string
	StatusCode int
}

// NewHTTPError returns a new HTTPError with the given code and message.
func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

// NewHTTPError returns a new HTTPError with the given code and message.
func NewUnauthorizedHTTPError() *HTTPError {
	return &HTTPError{
		Code:       401,
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func NewForbiddenHTTPError() *HTTPError {
	return &HTTPError{
		Code:       403,
		Message:    "Forbidden",
		StatusCode: http.StatusForbidden,
	}
}

// Error returns the error message.
func (e HTTPError) Error() string {
	return e.Message
}
