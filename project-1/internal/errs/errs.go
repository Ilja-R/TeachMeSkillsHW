package errs

import "errors"

// Generic errors
var (
	ErrNotfound           = errors.New("not found")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
)

// User specific error
var (
	ErrUserNotfound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user id")
)
