package usecase

import (
	"context"

	"chatty/chatty/entity"
)

type ChatService interface {
	Run()
	Shutdown()
	ListenForShutdown()
}

type ChatUseCase interface {
	Register(ctx context.Context, user entity.UserCred) error
	Login(ctx context.Context, user entity.UserCred) error
}

type ChatLogger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
}

type ChatUserRepo interface {
	AddUser(ctx context.Context, user entity.User) error
	CheckUsername(ctx context.Context, username string) (bool, error)
	GetPassword(ctx context.Context, username string) (string, error)
}

type ChatTokenRepo interface {
}
