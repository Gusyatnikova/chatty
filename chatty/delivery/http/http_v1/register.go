package http_v1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/delivery"
	"chatty/chatty/entity"
	"chatty/pkg/http/jwt"
)

type (
	registerReqBody struct {
		UserCreds
		UserContacts
	} //@name RegisterRequestBody

	registerRespBody struct {
		ID        string           `json:"id" example:"018496f4-77d7-0ef1-c2d2-f2b09e7b3fb1"`
		UserLogin entity.UserLogin `json:"login"  example:"testUser123"`
		UserContacts
	} //@name RegisterResponseBody
)

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

	jwtToken, err := jwt.GenerateToken(e.jwtCfg, newUser.Creds.Login)
	if err != nil {
		return err
	}

	writeAuthBearerHeader(eCtx, jwtToken)

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

func registerReqBodyToUser(body registerReqBody) entity.User {
	return entity.User{
		ID:       entity.UserID{},
		Creds:    entity.UserCreds(body.UserCreds),
		Contacts: entity.UserContacts(body.UserContacts),
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
