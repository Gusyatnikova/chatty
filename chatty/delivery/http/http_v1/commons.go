package http_v1

import (
	"strings"

	"github.com/labstack/echo/v4"

	"chatty/chatty/entity"
)

type (
	UserCreds struct {
		Login    entity.UserLogin    `json:"login"  example:"testUser123"`
		Password entity.UserPassword `json:"password"  example:"q123!@#Q"`
	} //@name UserCredentials

	UserContacts struct {
		Email       string `json:"email" example:"example@gmail.com"`
		PhoneNumber string `json:"phone_number" example:"+71234567890"`
	}
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

//writeAuthBearerHeader set Authorization header for response in eCtx
//to value "Bearer {token}"
func writeAuthBearerHeader(eCtx echo.Context, token string) {
	eCtx.Response().Header().Set("Authorization", "Bearer "+token)
}
