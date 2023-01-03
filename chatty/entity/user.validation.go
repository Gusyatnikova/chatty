package entity

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var (
	LoginValidationRegexp    = regexp.MustCompile("^[A-Za-z_]+[0-9]+$")
	PasswordValidationRegexp = regexp.MustCompile("^[A-Za-z0-9!@#$%^&*_+\\-=\\()]+$")
)

func (e User) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Creds),
		validation.Field(&e.Contacts),
	)
}

func (e UserCreds) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Login, validation.Required, validation.Length(1, 256), validation.Match(LoginValidationRegexp)),
		validation.Field(&e.Password, validation.Required, validation.Length(1, 128), validation.Match(PasswordValidationRegexp)),
	)
}

func (e UserContacts) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.PhoneNumber, validation.Required, is.E164),
		validation.Field(&e.Email, validation.Required, is.Email),
	)
}
