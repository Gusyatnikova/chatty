package http_v1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/delivery"
	"chatty/chatty/entity"
)

type registerReqBody struct {
	UserCreds
	UserContacts
} //@name RegisterRequestBody

// Register godoc
// @Summary Register new user
// @Tags    User operations
// @Produce json
// @Param   request body registerReqBody true "Login, Password, Email, Phone number for user"
// @Success 201     {object} userRespBody
// @Failure 400     "Request body is incorrect or data validation have failed"
// @Failure 409     "User with the specified login | email | phone number is already exists"
// @Failure 415     "Content-Type application/json is missing"
// @Router  /register [post]
func (e *ServerHandler) Register(eCtx echo.Context) error {
	registerBody, err := parseRegisterReqBody(eCtx)
	if err != nil {
		return err
	}

	newUser := registerReqBodyToUser(registerBody)
	if err := newUser.Validate(); err != nil {
		return err
	}

	if err := e.uc.Register(eCtx.Request().Context(), *newUser); err != nil {
		return err
	}

	user, err := e.uc.GetUserByLogin(eCtx.Request().Context(), newUser.Creds.Login)
	if err != nil {
		return err
	}

	jwtToken, err := e.jwtManager.GenerateJwtToken(user)
	if err != nil {
		return err
	}

	writeAuthBearerHeader(eCtx, jwtToken)

	return eCtx.JSON(http.StatusCreated, userToRespBody(user))
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

func registerReqBodyToUser(body registerReqBody) *entity.User {
	return entity.NewUser(
		entity.UserCreds(body.UserCreds),
		entity.UserContacts(body.UserContacts),
	)
}
