package jwt

import (
	"chatty/chatty/app/config"
	"github.com/golang-jwt/jwt"
)

var SigningMethod = jwt.SigningMethodHS256

type JWTManager struct {
	cfg config.JWT
}

func NewJWTManager(cfg config.JWT) *JWTManager {
	return &JWTManager{cfg: cfg}
}
