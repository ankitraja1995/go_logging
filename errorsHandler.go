package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type ErrorType uint

const (
	NoType ErrorType = iota
	FluentdPostError
	InvalidMethodNameError
)

type customError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
}

type errorContext struct {
	Field   string
	Message string
}

func (error customError) Error() string {

	return error.originalError.Error()
}

func (errorType ErrorType) Newf(msg string, args ...interface{}) error {

	return customError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

func Wrapf(err error, msg string, args ...interface{}) error {
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
