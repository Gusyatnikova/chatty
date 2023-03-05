package jwt

import (
	"chatty/chatty/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func (e *JWTManager) Generate(user entity.User) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(e.cfg.TTL)).Unix(),
		Subject:   string(user.GetID()),
	})

	return token.SignedString([]byte(e.cfg.Sign))
}
