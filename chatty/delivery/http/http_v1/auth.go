package http_v1

import (
	"chatty/chatty/delivery"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"

	"chatty/chatty/entity"
)

type UserCreds struct {
	Login    entity.UserLogin    `json:"login"  example:"testUser123"`
	Password entity.UserPassword `json:"password"  example:"q123!@#Q"`
} //@name UserCredentials

type UserContacts struct {
	Email       string `json:"email" example:"example@gmail.com"`
	PhoneNumber string `json:"phone_number" example:"+71234567890"`
}

type registerReqBody struct {
	UserCreds
	UserContacts
} //@name RegisterRequestBody

type loginReqBody struct {
	UserCreds
} //@name LoginRequestBody

type registerRespBody struct {
	ID        string           `json:"id" example:"018496f4-77d7-0ef1-c2d2-f2b09e7b3fb1"`
	UserLogin entity.UserLogin `json:"login"  example:"testUser123"`
	UserContacts
} //@name RegisterResponseBody

// Login godoc
// @Summary Login user in system by checking the specified password
// @Tags    User operations
// @Param   request body loginReqBody true "Login and password for user"
// @Success 200
// @Failure 400     "Request body is incorrect or data validation have failed"
// @Failure 404     "User with the specified login is not exists"
// @Failure 415     "Content-Type application/json is missing"
// @Router  /login [post]
func (e *ServerHandler) Login(eCtx echo.Context) error {
	loginBody, err := parseLoginReqBody(eCtx)
	if err != nil {
		return err
	}

	userCreds := loginReqBodyToUserCreds(loginBody)
	if err := userCreds.Validate(); err != nil {
		return err
	}

	return e.uc.Login(eCtx.Request().Context(), userCreds)
}

// Register godoc
// @Summary Register new user
// @Tags    User operations
// @Produce json
// @Param   request body registerReqBody true "Login, Password, Email, Phone number for user"
// @Success 201     {object} registerRespBody
// @Failure 400     "Request body is incorrect or data validation have failed"
// @Failure 409     "User with the specified login | email | phone number is already exists"
// @Failure 415     "Content-Type application/json is missing"
// @Router  /register [post]
func (e *ServerHandler) Register(eCtx echo.Context) error {
	registerBody, err := parseRegisterReqBody(eCtx)
	if err != nil {
		return err
	}

	user := registerReqBodyToUser(registerBody)
	if err := user.Validate(); err != nil {
		return err
	}

	if err := e.uc.Register(eCtx.Request().Context(), user); err != nil {
		return err
	}

	newUser, err := e.uc.GetUserByLogin(eCtx.Request().Context(), user.Creds.Login)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusCreated, userToRegisterRespBody(newUser))
}

func parseRegisterReqBody(eCtx echo.Context) (registerReqBody, error) {
	body := registerReqBody{}

	if !isRequestBodyIsJSON(eCtx) {
		return body, delivery.ErrBadContentType
	}

	err := eCtx.Bind(&body)
	if err != nil {
		return body, delivery.ErrBadRequestBody
	}

	body.Login = entity.UserLogin(strings.ToLower(string(body.Login)))

	return body, nil
}

func parseLoginReqBody(eCtx echo.Context) (loginReqBody, error) {
	body := loginReqBody{}

	if !isRequestBodyIsJSON(eCtx) {
		return body, delivery.ErrBadContentType
	}

	err := eCtx.Bind(&body)
	if err != nil {
		return body, delivery.ErrBadRequestBody
	}

	body.Login = entity.UserLogin(strings.ToLower(string(body.Login)))

	return body, nil
}

func registerReqBodyToUser(body registerReqBody) entity.User {
	return entity.User{
		ID:       entity.UserID{},
		Creds:    entity.UserCreds(body.UserCreds),
		Contacts: entity.UserContacts(body.UserContacts),
	}
}

func loginReqBodyToUserCreds(body loginReqBody) entity.UserCreds {
	return entity.UserCreds{
		Login:    body.Login,
		Password: body.Password,
	}
}

func userToRegisterRespBody(user entity.User) registerRespBody {
	return registerRespBody{
		UserLogin: user.Creds.Login,
		UserContacts: UserContacts{
			Email:       user.Contacts.Email,
			PhoneNumber: user.Contacts.PhoneNumber,
		},
		ID: user.GetID(),
	}
}
