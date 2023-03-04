package chatty

import (
	"chatty/chatty/repository"
	"chatty/chatty/usecase"
	"chatty/pkg/password"
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
