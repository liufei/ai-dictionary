package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	MissingMandatoryParamErrMsg = "Missing mandatory parameter(s) : %v"
	internalServerErrMsg        = "Unable to process the request due to an internal error. Please contact the system administrator"
)

// RestErr represents a REST API error
type RestErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
	Error      string `json:"error"`
}

// New returns a error.
func New(msg string) error {
	return errors.New(msg)
}

// Newf returns a formatted error.
func Newf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}

// BadRequestError returns a bad request RestErr.
func BadRequestError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      http.StatusText(http.StatusBadRequest),
	}
}

// BadRequestErrorf returns a formatted bad request RestErr.
func BadRequestErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusBadRequest,
		Error:      http.StatusText(http.StatusBadRequest),
	}
}

// UnauthorizedError returns a unauthorized RestErr.
func UnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Error:      http.StatusText(http.StatusUnauthorized),
	}
}

// UnauthorizedErrorf returns a formatted unauthorized RestErr.
func UnauthorizedErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusUnauthorized,
		Error:      http.StatusText(http.StatusUnauthorized),
	}
}

// ForbiddenError returns a forbidden RestErr.
func ForbiddenError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusForbidden,
		Error:      http.StatusText(http.StatusForbidden),
	}
}

// ForbiddenErrorf returns a formatted forbidden RestErr.
func ForbiddenErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusForbidden,
		Error:      http.StatusText(http.StatusForbidden),
	}
}

// NotFoundError returns a not found RestErr.
func NotFoundError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      http.StatusText(http.StatusNotFound),
	}
}

// NotFoundErrorf returns a fromatted not found RestErr.
func NotFoundErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusNotFound,
		Error:      http.StatusText(http.StatusNotFound),
	}
}

// ConflictError returns a conflict RestErr.
func ConflictError(message string) *RestErr {
	return &RestErr{
		Message:    message,
		StatusCode: http.StatusConflict,
		Error:      http.StatusText(http.StatusConflict),
	}
}

// ConflictErrorf returns a formatted conflict RestErr.
func ConflictErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    fmt.Sprintf(format, a...),
		StatusCode: http.StatusConflict,
		Error:      http.StatusText(http.StatusConflict),
	}
}

// InternalServerError returns a internal server RestErr.
func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message:    internalServerErrMsg,
		StatusCode: http.StatusInternalServerError,
		Error:      message,
	}
}

// InternalServerErrorf returns a formatted internal server RestErr.
func InternalServerErrorf(format string, a ...interface{}) *RestErr {
	return &RestErr{
		Message:    internalServerErrMsg,
		StatusCode: http.StatusInternalServerError,
		Error:      fmt.Sprintf(format, a...),
	}
}
