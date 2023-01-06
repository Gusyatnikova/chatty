package jwt

import "github.com/pkg/errors"

var (
	ErrUnableGenerateToken = errors.New("Error while generating JWT")
)
