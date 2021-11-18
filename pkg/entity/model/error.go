package model

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// DBError is error code of database
	DBError = "DB_ERROR"
	// NotFoundError is error code of not found
	NotFoundError = "NOT_FOUND_ERROR"
	// ValidationError is error code of validation
	ValidationError = "VALIDATION_ERROR"
	// BadRequestError is error code of request
	BadRequestError = "BAD_REQUEST_ERROR"
	// InternalServerError is error code of server error
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

// StackTrace is in an interface to check to see if the error has already been wrapped by errors.WithStack
type StackTrace interface {
	StackTrace() errors.StackTrace
}

// NewDBError returns error message related database
func NewDBError(e error) error {
	return newError(
		DBError,
		fmt.Sprintf("[DB ERROR]: %s", e.Error()),
		e,
	)
}

// NewNotFoundError returns error message related not found
func NewNotFoundError(e error) error {
	return newError(
		NotFoundError,
		fmt.Sprintf("[NOT FOUND ERROR]: %s", e.Error()),
		e,
	)
}

// NewInvalidParamError returns error message related param
func NewInvalidParamError(e error) error {
	return newError(
		BadRequestError,
		fmt.Sprintf("[INVALID PARAMETER]: %s", e.Error()),
		e,
	)
}

// NewValidationError returns error message related validation
func NewValidationError(e error) error {
	return newError(
		ValidationError,
		fmt.Sprintf("[VALIDATION ERROR]: %s", e.Error()),
		e,
	)
}

// NewInternalServerError returns error message related syntax or other issues
func NewInternalServerError(e error) error {
	return newError(
		InternalServerError,
		fmt.Sprintf("[INTERNAL SERVER ERROR]: %s", e.Error()),
		e,
	)
}

type err struct {
	code    string
	message string
}

func (e *err) Error() string { return e.message }
func (e *err) Code() string  { return e.code }

func newError(code string, message string, e error) error {
	newErr := &err{
		code:    code,
		message: message,
	}
	if hasStackTrace(e) {
		return newErr
	}

	return withStackTrace(newErr)
}

func withStackTrace(e error) error {
	ews := errors.WithStack(e)

	fmt.Printf("%+v", ews)

	return ews
}

func hasStackTrace(e error) bool {
	_, ok := e.(StackTrace)
	return ok
}
