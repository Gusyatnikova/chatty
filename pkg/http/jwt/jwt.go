package jwt

import (
	"time"

	"chatty/chatty/app/config"
	"github.com/golang-jwt/jwt"

	"chatty/chatty/entity"
)

var SigningMethod = jwt.SigningMethodHS256

func GenerateJwtToken(cfg config.JWT, login entity.UserLogin) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, jwt.StandardClaims{
		Subject:   string(login),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(cfg.TTL)).Unix(),
	})

	signedString, err := token.SignedString([]byte(cfg.Sign))
	if err != nil {
		return "", ErrUnableGenerateToken
	}

	return signedString, nil
}
