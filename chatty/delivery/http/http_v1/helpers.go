package http_v1

import (
	"strings"

	"github.com/labstack/echo/v4"
)

//isRequestBodyIsJSON returns true if header Contain-type with value application/json are in the request
func isRequestBodyIsJSON(eCtx echo.Context) bool {
	contentTypes := eCtx.Request().Header.Get(echo.HeaderContentType)

	if contentTypes != "" {
		for _, ct := range strings.Split(contentTypes, ";") {
			if strings.TrimSpace(ct) == echo.MIMEApplicationJSON {
				return true
			}
		}
	}

	return false
}
