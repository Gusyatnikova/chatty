package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"chatty/chatty/entity"
)

func (e *JWTManager) GenerateAccessToken(user entity.User) (string, time.Time, error) {
	expAt := time.Now().Add(time.Second * time.Duration(e.cfg.TTL))

	token := jwt.NewWithClaims(SigningMethod, jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expAt},
		Subject:   string(user.GetID()),
	})

	ss, err := token.SignedString([]byte(e.cfg.Sign))

	return ss, expAt, errors.Wrap(err, "err in JWTManager.GenerateAccessToken():")
}
