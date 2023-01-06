package http_v1

import (
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/delivery"
	"chatty/chatty/entity"
	"chatty/pkg/http/jwt"
)

type loginReqBody struct {
	UserCreds
} //@name LoginRequestBody

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

	jwtToken, err := jwt.GenerateToken(e.jwtCfg, userCreds.Login)
	if err != nil {
		return err
	}

	writeAuthBearerHeader(eCtx, jwtToken)

	return e.uc.Login(eCtx.Request().Context(), userCreds)
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
