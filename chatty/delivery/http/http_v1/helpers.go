package http_v1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	mw "chatty/chatty/delivery/http/http_v1/middleware"
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

//setAccessToken set Authorization header to "Bearer {token}" and accessToken cookie value to {token}
func (e *ServerHandler) setAccessToken(rw http.ResponseWriter, token string, expAt time.Time) {
	rw.Header().Set(mw.AccessTokenHeaderName, fmt.Sprintf("%s %s", mw.AuthScheme, token))

	c := &http.Cookie{
		Name:     mw.AccessTokenCookieName,
		Value:    token,
		Expires:  expAt,
		HttpOnly: true,
	}

	http.SetCookie(rw, c)
}
