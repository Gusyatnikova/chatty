package jwt

import (
	gojwt "github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (e *JWTManager) ExtractSub(tokenString string) (string, error) {
	token, err := gojwt.Parse(tokenString, func(token *gojwt.Token) (interface{}, error) {
		return []byte(e.cfg.Sign), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "err in JWTManager.ExtractSub.Parse():")
	}

	if claims, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
		return claims["sub"].(string), nil
	}

	return "", errors.Wrap(err, "err in JWTManager.ExtractSub():")
}
