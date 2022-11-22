package entity

import "github.com/oklog/ulid/v2"

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

type UserID ulid.ULID

type User struct {
	ID       UserID
	Creds    UserCreds
	Contacts UserContacts
}

func (e User) GetID() string {
	return ulid.ULID(e.ID).String()
}
