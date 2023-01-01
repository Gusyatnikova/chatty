package app

import (
	chatty2 "chatty/chatty/usecase/chatty"
	"chatty/pkg/password"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"chatty/chatty/app/config"
	"chatty/chatty/delivery"
	"chatty/chatty/repository/postgres"
	"chatty/chatty/usecase"
	"chatty/pkg/http"
	dbConn "chatty/pkg/repository/postgres"
)

type chatty struct {
	httpServer http.Server
	usecase    usecase.ChatUseCase
}

func NewChatty(ctx context.Context) delivery.ChattyServer {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Panic().Msgf("err in NewChatty.NewConfig(): %s", err.Error())
	}

	initLogger()

	pgConn, err := dbConn.Connection(ctx, cfg.Pg)
	if err != nil {
		log.Panic().Msgf("err in NewChatty.dbConn.Connection(): %s", err.Error())
	}

	chattyRepo := postgres.NewPgChattyRepo(pgConn)
	pwdService := password.NewService(cfg.Password)
	uc := chatty2.NewChattyUseCase(chattyRepo, pwdService)

	httpServerCfg := http.ServerConfig{
		Address: fmt.Sprint(cfg.Http.Host, ":", cfg.Http.Port),
		JwtCfg:  cfg.Jwt,
	}
	httpServer := http.NewServer(ctx, httpServerCfg, uc)

	return &chatty{
		httpServer: httpServer,
		usecase:    uc,
	}
}

func (e *chatty) Run() {
	go e.httpServer.Run()
}

func (e *chatty) Shutdown() {
	e.httpServer.Shutdown()
}

func (e *chatty) ListenForShutdown() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	for v := range ch {
		if v == os.Interrupt {
			e.Shutdown()

			break
		}
	}
	close(ch)
}

func initLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
