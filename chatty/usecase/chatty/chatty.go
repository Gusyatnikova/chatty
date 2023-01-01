package chatty

import (
	"chatty/pkg/password"
	"context"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"

	"chatty/chatty/entity"
	"chatty/chatty/repository"
	"chatty/chatty/usecase"
)

type chatUseCase struct {
	userRepo        repository.ChattyUserRepo
	passwordService password.Passworder
}

func NewChattyUseCase(userRepo repository.ChattyUserRepo, pwdService password.Passworder) usecase.ChatUseCase {
	return &chatUseCase{
		userRepo:        userRepo,
		passwordService: pwdService,
	}
}

//Register adds new user and returns nil, if it ends successfully, ErrLoginDuplication if such user is already exists
func (e *chatUseCase) Register(ctx context.Context, user entity.User) error {
	_, err := e.userRepo.GetUserByLogin(ctx, user.Creds.Login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			err = e.addUser(ctx, user)
		}
		return errors.Wrap(err, "err in chatUseCase.Register():")
	}

	return usecase.ErrUserDuplication
}

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

func (e *chatUseCase) GetUserByLogin(ctx context.Context, login entity.UserLogin) (entity.User, error) {
	user := entity.User{}
	user, err := e.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return user, usecase.ErrNoUser
		}
		return user, errors.Wrap(err, "err in chatUseCase.GetUserByLogin.userRepo.GetUserByLogin():")
	}

	return user, nil
}

//addUser generate id and encrypted password for user and store user to db
func (e *chatUseCase) addUser(ctx context.Context, user entity.User) error {
	userInternal := user
	userInternal.ID = generateID()

	pwd, err := e.generateUserPassword(userInternal.Creds.Password)
	if err != nil {
		return errors.Wrap(err, "err in chatUseCase.addUser.generateUserPassword():")
	}

	userInternal.Creds.Password = entity.UserPassword(pwd)

	err = e.userRepo.AddUser(ctx, userInternal)
	if errors.Is(err, repository.ErrDuplicate) {
		return errors.Wrap(usecase.ErrDataDuplication, err.Error())
	}

	return errors.Wrap(err, "err in chatUseCase.addUser.userRepo.AddUser():")
}

//generateUserPassword returns encoded hash of password
func (e *chatUseCase) generateUserPassword(password entity.UserPassword) (string, error) {
	pwdHash, err := e.passwordService.Generate(string(password))
	if err != nil {
		return "", err
	}

	return pwdHash, nil
}

//generateID returns unique UserID
func generateID() entity.UserID {
	return entity.UserID(ulid.Make())
}

//authUser checks if user exists and then perform password checking. It returns
//(true, nil) if auth was successful
//(false, ErrNoUser) if user with creds.Login was not found
//(false, error) if other errors occur
func (e *chatUseCase) authUser(ctx context.Context, creds entity.UserCreds) (bool, error) {
	storedUser, err := e.userRepo.GetUserByLogin(ctx, creds.Login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, usecase.ErrNoUser
		}
		return false, errors.Wrap(err, "err in chatUseCase.userRepo.GetUserByLogin():")
	}

	match, err := e.passwordService.Compare(string(creds.Password), string(storedUser.Creds.Password))
	return match, errors.Wrap(err, "err in chatUseCase.authUser.passwordService.Compare():")
}
