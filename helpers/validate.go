package helpers

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateUserCreate validates user name, email and password when creating user.
// Name cannot be empty and the length must be between 2 and 100.
// Email cannot be empty, the length must be between 3 and 250 
// and it must be an email in terms of address string structure.
// Password cannot be empty and the length must be between 8 and 100.
func ValidateUserCreate(n string, e string, p string) error {
	err := validation.Errors{
		"Name": validation.Validate(n, validation.Required, validation.Length(2, 100)),
		"Email": validation.Validate(e, validation.Required, validation.Length(3, 100), is.Email),
		"Password": validation.Validate(p, validation.Required, validation.Length(8, 100)),
	}.Filter()
	if err != nil {
		return err
	}

	return nil
}

// ValidateUserAuth validates user email and password when authenticating user.
// Email cannot be empty, the length must be between 3 and 250 
// and it must be an email in terms of address string structure.
// Password cannot be empty and the length must be between 8 and 100.
func ValidateUserAuth(e string, p string) error {
	err := validation.Errors{
		"Email": validation.Validate(e, validation.Required, validation.Length(3, 100), is.Email),
		"Password": validation.Validate(p, validation.Required, validation.Length(8, 100)),
	}.Filter()
	if err != nil {
		return err
	}

	return nil
}

// // ValidateUserEmail validates user email.
func ValidateUserEmail(e string) error {
	err := validation.Errors{
		"Email": validation.Validate(e, validation.Required, validation.Length(3, 100), is.Email),
	}.Filter()
	if err != nil {
		return err
	}

	return nil
}