package chatty

import (
	"context"

	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/usecase"
)

//Login performs authorization. Returns nil if auth ends successfully,
//ErrNoUser if user with creds.Login was not found, ErrUserUnauthorized if auth process ends with error.
func (e *chatUseCase) Login(ctx context.Context, creds entity.UserCreds) error {
	isAuth, err := e.authUser(ctx, creds)
	if err != nil {
		return err
	}
	if !isAuth {
		return usecase.ErrUserUnauthorized
	}

	return nil
}

//authUser checks if user exists and then perform password checking. It returns
//(true, nil) if auth was successful
//(false, ErrNoUser) if user with creds.Login was not found
//(false, error) if other errors occur
func (e *chatUseCase) authUser(ctx context.Context, creds entity.UserCreds) (bool, error) {
	storedUser, err := e.userRepo.GetUserByLogin(ctx, creds.Login)
	if err != nil {
		return false, usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.userRepo.GetUserByLogin():"))
	}

	match, err := e.passwordService.Compare(string(creds.Password), string(storedUser.Creds.Password))

	return match, usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.authUser.passwordService.Compare():"))
}
