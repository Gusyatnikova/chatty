package jwt

import (
	"chatty/chatty/app/config"
	"github.com/golang-jwt/jwt/v4"
)

var SigningMethod = jwt.SigningMethodHS256

type JWTManager struct {
	cfg config.JWT
}

func NewTokenManager(cfg config.JWT) TokenManager {
	return &JWTManager{cfg: cfg}
}
