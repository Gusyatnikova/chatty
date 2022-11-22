package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (e User) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Creds),
		validation.Field(&e.Contacts),
	)
}

func (e UserCreds) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Login, validation.Required, validation.Length(1, 256), is.Alphanumeric),
		validation.Field(&e.Password, validation.Required, validation.Length(1, 128), is.PrintableASCII),
	)
}

func (e UserContacts) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.PhoneNumber, validation.Required, is.E164),
		validation.Field(&e.Email, validation.Required, is.Email),
	)
}
