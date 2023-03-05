package jwt

import (
	"time"

	"chatty/chatty/app/config"
	"chatty/chatty/entity"
)

type TokenManager interface {
	GenerateAccessToken(user entity.User) (string, time.Time, error)
	ExtractSub(token string) (string, error)
	GetConfig() config.JWT
}
