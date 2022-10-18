package http

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	"chatty/chatty/delivery/http/http_v1"
	"chatty/chatty/usecase"
)

type ChatServer interface {
	Run()
	Shutdown()
}

type ServerConfig struct {
	Address     string
	ChatUsecase usecase.ChatUseCase
	Logger      usecase.ChatLogger
}

type chatHttpServer struct {
	ctx        context.Context
	cfg        ServerConfig
	httpServer *echo.Echo
}

func NewChatHttpServer(ctx context.Context, cfg ServerConfig) ChatServer {
	e := echo.New()

	e.Server.Addr = cfg.Address
	//todo: change to logging with logLevel,
	//use mw to work wityh tokens
	e.Use(
		mw.LoggerWithConfig(mw.LoggerConfig{
			Format: `{"time":"${time_rfc3339_nano}","remote_ip":"${remote_ip}","host":"${host}","method":"${method}",` +
				`"uri":"${uri}","query":"${query}","user_agent":"${user_agent}","status":${status},"error":"${error}"}` + "\n",
		}),
		mw.Recover())

	//ToDo: add tokens handler to midleware chain

	http_v1.NewChatServerHandler(e, cfg.ChatUsecase, cfg.Logger)

	return &chatHttpServer{
		ctx:        ctx,
		cfg:        cfg,
		httpServer: e,
	}
}

func (e *chatHttpServer) Run() {
	e.cfg.Logger.Info(fmt.Sprintf("HTTP server listening at %v", e.httpServer.Server.Addr))
	if err := e.httpServer.Server.ListenAndServe(); err != nil {
		e.cfg.Logger.Panic(fmt.Sprintf("err in chatHttpServer.Run(): %s", err.Error()))
	}
}

func (e *chatHttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(e.ctx, 2*time.Second)
	defer cancel()

	e.cfg.Logger.Warn("Shutting down the chatHttpServer")

	if err := e.httpServer.Server.Shutdown(ctx); err != nil {
		e.cfg.Logger.Error(fmt.Sprintf("err in chatHttpServer.Shutdown(): %s", err.Error()))
	}
}
