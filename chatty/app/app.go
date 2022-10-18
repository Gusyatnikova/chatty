package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"chatty/chatty/app/config"
	"chatty/chatty/repository/postgres"
	"chatty/chatty/usecase"
	"chatty/pkg/http"
	"chatty/pkg/logger"
	dbConn "chatty/pkg/repository/postgres"
)

//todo: delete this layer and move echo server here

type chatService struct {
	httpServer http.ChatServer
}

//todo: rename to application
func NewChatService(ctx context.Context) usecase.ChatService {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("err in config.NewConfig(): %s", err)
	}

	usrRepo := postgres.NewPgUserRepo(dbConn.Connection(cfg.PG))

	chatUC := usecase.NewChatUseCase(usrRepo)

	httpServerCfg := http.ServerConfig{
		Address:     fmt.Sprint(cfg.HTTP.Host, ":", cfg.HTTP.Port),
		ChatUsecase: chatUC,
		//todo: use zero-log, delete logger
		Logger: logger.NewLogger(),
	}
	httpServer := http.NewChatHttpServer(ctx, httpServerCfg)

	return &chatService{httpServer: httpServer}
}

func (e *chatService) Run() {
	e.httpServer.Run()

}

func (e *chatService) Shutdown() {
	e.httpServer.Shutdown()
}

func (e *chatService) ListenForShutdown() {
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
