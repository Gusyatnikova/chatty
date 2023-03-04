package http_v1

import (
	"chatty/chatty/entity"
)

type UserContacts struct {
	Email       string `json:"email" example:"example@gmail.com"`
	PhoneNumber string `json:"phone_number" example:"+71234567890"`
} //@name UserContacts

type UserCreds struct {
	Login    entity.UserLogin    `json:"login"  example:"testUser123"`
	Password entity.UserPassword `json:"password"  example:"q123!@#Q"`
} //@name UserCredentials

type userRespBody struct {
	ID        string           `json:"id" example:"018496f4-77d7-0ef1-c2d2-f2b09e7b3fb1"`
	UserLogin entity.UserLogin `json:"login"  example:"testUser123"`
	UserContacts
} //@name UserResponseBody

func userToRespBody(user entity.User) userRespBody {
	return userRespBody{
		UserLogin: user.Creds.Login,
		UserContacts: UserContacts{
			Email:       user.Contacts.Email,
			PhoneNumber: user.Contacts.PhoneNumber,
		},
		ID: string(user.GetID()),
	}
}
