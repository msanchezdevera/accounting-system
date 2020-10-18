package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

// ErrorType is the type of an error
type ErrorType uint

const (
	// NoType is an error without a specific type
	NoType ErrorType = iota
	// UserError indicates the user data is incorrect
	UserError
	// StatusUnsupportedMediaType indicates that the server refuses to accept the request because the payload format is in an unsupported format
	StatusUnsupportedMediaType
	// Forbidden indicates that the server understood the request but refuses to authorize it
	Forbidden
	NotFound
)

// Error represents a custom error used throughout the project
type Error interface {
	Type() ErrorType
	Error() string
	Context() map[string]string
	AddSingleContext(string, string) Error
}

type customError struct {
	errorType     ErrorType
	originalError error
	context       map[string]string
}

// New creates a new customError
func (errorType ErrorType) New(msg string) Error {
	return customError{errorType: errorType, originalError: errors.New(msg)}
}

// Newf creates a new customError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) Error {
	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrapf creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) Error {
	return customError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

// New creates a no type error
func New(msg string) Error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) Error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) Error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return customError{errorType: NoType, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func (err customError) AddSingleContext(field, message string) Error {
	if err.context == nil {
		err.context = map[string]string{}
	}

	err.context[field] = message

	return err
}

// GetErrorContext returns the error context
func (err customError) Context() map[string]string {
	return err.context
}

// Error returns the mssage of a customError
func (err customError) Error() string {
	return err.originalError.Error()
}

// GetType returns the error type
func (err customError) Type() ErrorType {
	return err.errorType
}
