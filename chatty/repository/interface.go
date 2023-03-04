package repository

import (
	"context"

	"chatty/chatty/entity"
)

type ChattyUserRepo interface {
	AddUser(context.Context, entity.User) error
	GetUserByLogin(context.Context, entity.UserLogin) (entity.User, error)
	GetUserByID(context.Context, entity.UserID) (entity.User, error)

	HealthCheck(ctx context.Context) error
}
