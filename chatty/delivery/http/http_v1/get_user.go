package http_v1

import (
	"chatty/chatty/delivery"
	"chatty/chatty/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

// GetUserByLogin godoc
// @Summary Return information about user based on login param
// @Tags    User operations
// @Param   login path string true "User's login"
// @Produce json
// @Success 200		{object} userRespBody
// @Failure 400     "Data validation have failed"
// @Failure 404     "User with the specified login is not exists"
// @Router  /user/{login} [get]
func (e *ServerHandler) GetUserByLogin(eCtx echo.Context) error {
	userLogin := entity.UserLogin(eCtx.Param("login"))

	if err := (&userLogin).Validate(); err != nil {
		return err
	}

	user, err := e.uc.GetUserByLogin(eCtx.Request().Context(), userLogin)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, userToRespBody(user))
}

// GetUserByToken godoc
// @Summary Return information about user based on jwt token
// @Tags    User operations
// @Produce json
// @Success 200		{object} userRespBody
// @Failure 404     "User is not found"
// @Router  /whoami [get]
func (e *ServerHandler) GetUserByToken(eCtx echo.Context) error {
	errMsg := "err in ServerHandler.GetUserByToken"

	token := e.getToken(eCtx.Request())

	idStr, err := e.jwtManager.ExtractSub(token)
	if err != nil {
		return errors.Wrap(delivery.ErrBadRequest, fmt.Sprintf("%s.jwtManager.ExtractSub() with token %s", errMsg, token))
	}

	user := entity.User{}

	user, err = e.uc.GetUserByID(eCtx.Request().Context(), entity.UserID(idStr))
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, userToRespBody(user))
}

func (e *ServerHandler) getToken(r *http.Request) string {
	jwtConfig := e.jwtManager.GetConfig()

	if atCookie, err := r.Cookie(jwtConfig.AccessTokenCookieName); err == nil {
		return atCookie.Value
	}

	authScheme := jwtConfig.AuthScheme
	authHeader := r.Header.Get(jwtConfig.AccessTokenHeaderName)

	headerParts := strings.Split(authHeader, authScheme)
	if len(headerParts) < 2 {
		return ""
	}

	return strings.TrimSpace(headerParts[1])
}
