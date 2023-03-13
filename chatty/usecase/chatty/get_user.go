package chatty

import (
	"context"

	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/usecase"
)

func (e *chatUseCase) GetUserByLogin(ctx context.Context, login entity.UserLogin) (entity.User, error) {
	user, err := e.userRepo.GetUserByLogin(ctx, login)

	if err != nil {
		return user, usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.GetUserByLogin.userRepo.GetUserByLogin():"))
	}

	return user, nil
}

func (e *chatUseCase) GetUserByID(ctx context.Context, id entity.UserID) (entity.User, error) {
	user, err := e.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return user, usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.GetUserByLogin.userRepo.GetUserByID():"))
	}

	return user, nil
}
