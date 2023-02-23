package http_v1

import (
	"chatty/chatty/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (e *ServerHandler) GetUser(eCtx echo.Context) error {
	userLogin := entity.UserLogin(eCtx.Param("login"))

	user, err := e.uc.GetUserByLogin(eCtx.Request().Context(), userLogin)
	if err != nil {
		return err
	}

	return eCtx.JSON(http.StatusOK, userToRespBody(user))
}
