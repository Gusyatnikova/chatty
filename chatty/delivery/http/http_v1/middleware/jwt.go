package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	"chatty/chatty/app/config"
	"chatty/chatty/delivery"
	chattyjwt "chatty/pkg/http/jwt"
)

var whiteListPaths = []string{
	"/login",
	"/register",
	"/health",
	"/swagger/*",
}

func JWTHandlerMiddleware(cfg config.JWT) echo.MiddlewareFunc {
	mw.ErrJWTMissing.Message = delivery.ErrUnauthorizied.Error()
	mw.ErrJWTMissing.Code = http.StatusUnauthorized

	return mw.JWTWithConfig(mw.JWTConfig{
		SigningKey:    []byte(cfg.Sign),
		SigningMethod: chattyjwt.SigningMethod.Name,
		Skipper:       skipAuth,
	})
}

func skipAuth(e echo.Context) bool {
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}

	return false
}
