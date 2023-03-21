package usecase

import (
	"context"

	"chatty/chatty/entity"
)

//go:generate mockery --name ChatUseCase
type ChatUseCase interface {
	Register(context.Context, entity.User) error
	Login(context.Context, entity.UserCreds) error
	GetUserByLogin(context.Context, entity.UserLogin) (entity.User, error)
	GetUserByID(context.Context, entity.UserID) (entity.User, error)

	HealthCheck(ctx context.Context) error
}
