package http_v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//todo: add swagger

func (e *ChatServerHandler) Signup(eCtx echo.Context) error {
	usrCred, err := e.parseUserCredBody(eCtx)
	if err != nil {
		return e.noContentErrResponse(eCtx, http.StatusBadRequest,
			fmt.Sprintf("err in ChatServerHandler.Signup.parseUserCredBody(): %v", err))
	}

	return e.uc.Register(eCtx.Request().Context(), *usrCred)
}

func (e *ChatServerHandler) Signin(eCtx echo.Context) error {
	return nil
}
