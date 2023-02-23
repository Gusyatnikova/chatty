package jwt

import (
	"chatty/chatty/app/config"
	"github.com/golang-jwt/jwt"
	"time"
)

var SigningMethod = jwt.SigningMethodHS256

func GenerateJwtToken(cfg config.JWT, sub string) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(cfg.TTL)).Unix(),
		Subject:   sub,
	})

	return token.SignedString([]byte(cfg.Sign))
}
