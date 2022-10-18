package http_v1

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/entity"
)

func (e *ChatServerHandler) parseUserCredBody(eCtx echo.Context) (*entity.UserCred, error) {
	contentTypes := eCtx.Request().Header.Get(echo.HeaderContentType)

	if contentTypes != "" {
		for _, ct := range strings.Split(contentTypes, ";") {
			if strings.TrimSpace(ct) == echo.MIMEApplicationJSON {
				user := new(entity.UserCred)

				err := eCtx.Bind(user)
				if err != nil {
					return nil, err
				}

				return user, nil
			}
		}
	}

	return nil, errors.New("Content-Type header is missed")
}

func (e *ChatServerHandler) noContentErrResponse(eCtx echo.Context, statusCode int, errMsg string) error {
	e.logger.Error(errMsg)

	return eCtx.NoContent(statusCode)
}
