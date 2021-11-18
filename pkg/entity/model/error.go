package model

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// DBError is error code of database
	DBError = "00001"
	// NotFoundError is error code of not found
	NotFoundError = "000002"
	// ValidationError is error code of validation
	ValidationError = "000100"
	// BadRequestError is error code of request
	BadRequestError = "000400"
	// InternalServerError is error code of server error
	InternalServerError = "009999"
)

// NewDBError returns error message related database
func NewDBError(e error) error {
	return newError(
		DBError,
		fmt.Sprintf("[DB ERROR]: %s", e.Error()),
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
	ok := hasStackTrace(e)
	if ok {
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
	_, ok := e.(interface{ StackTrace() errors.StackTrace })
	return ok
}
