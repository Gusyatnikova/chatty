package http_v1

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/delivery"
	"chatty/chatty/entity"
)

type loginReqBody struct {
	UserCreds
} //@name LoginRequestBody

// Login godoc
// @Summary Login user in system by checking the specified password
// @Tags    User operations
// @Param   request body loginReqBody true "Login and password for user"
// @Produce json
// @Success 200		{object} userRespBody
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
	if err := (&userCreds).Validate(); err != nil {
		return err
	}

	if err := e.uc.Login(eCtx.Request().Context(), userCreds); err != nil {
		return err
	}

	user, err := e.uc.GetUserByLogin(eCtx.Request().Context(), userCreds.Login)
	if err != nil {
		return nil
	}

	jwtToken, expTime, err := e.jwtManager.GenerateAccessToken(user)
	if err != nil {
		return err
	}

	e.setAccessToken(eCtx.Response(), jwtToken, expTime)

	return eCtx.JSON(http.StatusOK, userToRespBody(user))
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

func loginReqBodyToUserCreds(body loginReqBody) entity.UserCreds {
	return entity.UserCreds{
		Login:    body.Login,
		Password: body.Password,
	}
}
