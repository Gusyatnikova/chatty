package chatty

import (
	"context"

	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/usecase"
)

//Register adds new user and returns nil, if it ends successfully, ErrLoginDuplication if such user is already exists
func (e *chatUseCase) Register(ctx context.Context, user entity.User) error {
	_, err := e.userRepo.GetUserByLogin(ctx, user.Creds.Login)
	if err != nil {
		if errors.Is(usecase.ErrorHandling(err), usecase.ErrNoUser) {
			err = e.addUser(ctx, user)
		}
		return usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.Register():"))
	}

	return usecase.ErrUserDuplication
}

//addUser generate id and encrypted password for user and store user to db
func (e *chatUseCase) addUser(ctx context.Context, user entity.User) error {
	userInternal := user

	pwd, err := e.generateUserPassword(userInternal.Creds.Password)
	if err != nil {
		return usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.addUser.generateUserPassword():"))
	}

	userInternal.Creds.Password = entity.UserPassword(pwd)

	err = e.userRepo.AddUser(ctx, userInternal)
	if err != nil {
		return usecase.ErrorHandling(err)
	}

	return usecase.ErrorHandling(errors.Wrap(err, "err in chatUseCase.addUser.userRepo.AddUser():"))
}

//generateUserPassword returns encoded hash of password
func (e *chatUseCase) generateUserPassword(password entity.UserPassword) (string, error) {
	pwdHash, err := e.passwordService.Generate(string(password))
	if err != nil {
		return "", usecase.ErrorHandling(err)
	}

	return pwdHash, nil
}
