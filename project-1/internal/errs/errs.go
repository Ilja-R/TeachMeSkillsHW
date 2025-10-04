package errs

import "errors"

// Generic errors
var (
	ErrNotfound           = errors.New("not found")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
)

// Employee specific error
var (
	ErrEmployeeNotfound  = errors.New("employee not found")
	ErrInvalidEmployeeID = errors.New("invalid employee id")
)
