package model

import "fmt"

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
func NewDBError(m string) error {
	return &err{
		code:    DBError,
		message: fmt.Sprintf("[DB ERROR]: %s", m),
	}
}

// NewInvalidParamError returns error message related param
func NewInvalidParamError(m string) error {
	return &err{
		code:    BadRequestError,
		message: fmt.Sprintf("[INVALID PARAMETER]: %s", m),
	}
}

// NewValidationError returns error message related validation
func NewValidationError(m string) error {
	return &err{
		code:    ValidationError,
		message: fmt.Sprintf("[VALIDATION ERROR]: %s", m),
	}
}

// NewInternalServerError returns error message related syntax or other issues
func NewInternalServerError(m string) error {
	return &err{
		code:    InternalServerError,
		message: fmt.Sprintf("[INTERNAL SERVER ERROR]: %s", m),
	}
}

type err struct {
	code    string
	message string
}

func (e *err) Error() string { return e.message }
func (e *err) Code() string  { return e.code }
