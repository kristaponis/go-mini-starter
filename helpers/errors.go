package helpers

import (
	"errors"
	"strings"
)

var (
	ErrGeneric       = errors.New("something went wrong, please try again")
	ErrUserNotFound  = errors.New("user not found")
	ErrPasswordMatch = errors.New("incorrect password")
	ErrEmailDupKey   = errors.New("this email is already taken")
)

// UserError contains processed error message.
type UserError struct {
	Message string
}

// NewUserError catches error, capitalizes first letter and
// returns that error as UserError.Message string.
// This is used in user handlers to pass errors to the templates.
func NewUserError(err error) *UserError {
	errMsg := err.Error()
	split := strings.Split(errMsg, " ")
	split[0] = strings.Title(split[0])
	message := strings.Join(split, " ")
	return &UserError{
		Message: message,
	}
}
