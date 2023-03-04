package entity

import (
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type (
	UserLogin    string
	UserPassword string
)

type UserCreds struct {
	Login    UserLogin
	Password UserPassword
}

type UserContacts struct {
	Email       string
	PhoneNumber string
}

type UserID string

type User struct {
	ID       ulid.ULID
	Creds    UserCreds
	Contacts UserContacts
}

func NewUser(creds UserCreds, contacts UserContacts) *User {
	return &User{
		ID:       ulid.Make(),
		Creds:    creds,
		Contacts: contacts,
	}
}

func (e *User) GetID() UserID {
	return UserID(e.ID.String())
}

func (e *User) SetID(idStr UserID) (err error) {
	if id, err := ulid.Parse(string(idStr)); err == nil {
		e.ID = id
	}

	return errors.Wrap(err, "err in user.SetID(): invalid id")
}
