package usecase

import (
	"context"
	"fmt"

	"chatty/chatty/entity"
)

type chatUseCase struct {
	userRepo  ChatUserRepo
	tokenRepo ChatTokenRepo
}

func NewChatUseCase(userRepo ChatUserRepo) ChatUseCase {
	return &chatUseCase{
		userRepo:  userRepo,
		tokenRepo: nil,
	}
}

func (e *chatUseCase) Register(ctx context.Context, user entity.UserCred) error {
	taken, err := e.userRepo.CheckUsername(ctx, user.Username)

	//todo: add register flow
	fmt.Println(taken, err)

	return nil
}

func (e *chatUseCase) Login(ctx context.Context, user entity.UserCred) error { return nil }
