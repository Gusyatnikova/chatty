package http_v1

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "chatty/chatty/delivery/docs"
	"chatty/chatty/usecase"
)

// @title         Chat Server API
// @name          Chatty
// @contact.email gusiatnikovanatalia@gmail.com
// @version       1.0
// @license.name  free-to-use-license

type ServerHandler struct {
	uc usecase.ChatUseCase
}

func NewServerHandler(e *echo.Echo, uc usecase.ChatUseCase) {
	h := &ServerHandler{uc: uc}

	e.POST("login", h.Login)
	e.POST("register", h.Register)

	e.GET("health", h.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
