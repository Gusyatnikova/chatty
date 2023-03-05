package jwt

import (
	"time"

	"chatty/chatty/app/config"
)

type TokenManager interface {
	GenerateAccessToken(sub string) (string, time.Time, error)
	ExtractSub(token string) (string, error)
	GetConfig() config.JWT
}
