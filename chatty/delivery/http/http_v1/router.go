package http_v1

import (
	"github.com/labstack/echo/v4"

	"chatty/chatty/usecase"
)

type ChatServerHandler struct {
	uc     usecase.ChatUseCase
	logger usecase.ChatLogger
}

//todo: replace echo with interface?
func NewChatServerHandler(e *echo.Echo, uc usecase.ChatUseCase, logger usecase.ChatLogger) {
	h := &ChatServerHandler{uc: uc, logger: logger}

	e.POST("/signup/", h.Signup)
	e.POST("/signin/", h.Signin)
}
